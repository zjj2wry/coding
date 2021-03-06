## Distributed Tensorflow 
tensorflow 集群是由一系列的 tasks 来参与 tensorflow graph 的分布式计算，每一个 task 关联一个 tensorflow server,其中 master 用来创建 session，worker 用来计算图的操作，一个集群也可以被拆分成一个或者多个 job，每一个 job 包含一个或多个 task。

为了创建一个集群，你在集群的每一个 task 上面都启动一个 tensorflow server，每一个 task 都运行在不同的机器上，你也可以运行多个 task    在相同的机器上。（例如：控制不同的 gpu 设备。

## local server
### 启动服务器端,server.join() 避免进程退出，同样会暴露一个grpc的端口。
```
$ python
>>> import tensorflow as tf
>>> c = tf.constant("Hello world!"))
>>> server = tf.train.Server.create_local_server()
>>> server.join()  
I tensorflow/core/distributed_runtime/rpc/grpc_channel.cc:206] Initialize HostPortsGrpcChannelCache for job local -> {localhost:40767}
I tensorflow/core/distributed_runtime/rpc/grpc_server_lib.cc:202] Started server with target: grpc://localhost:40767
```
### 客户端
```
$ python
>>> import tensorflow as tf
>>> c = tf.constant("Hello world!"))
>>> server_target = "grpc://localhost:40767"
>>> sess = tf.Session(server_target)
>>> print(sess.run(c))
>>> sess.close()  
'Hello world!'
```

## concept

- task: task 代表一个单独的进程，对应一个 tensorflow server。
- ps：parameter server，用来保存和更新变量。
- worker：进行图的计算,计算模型梯度的节点，得到的梯度向量会交付给ps更新模型。
- job: job 可以是 ps 或者 worker，是 task 的集合，使用 job 的类型和 task index 确定一个具体的 task，然后创建一个server。
- client：编写 tensorflow 图的计算的程序。python，c++ 等。
- cluster：由很多的 job 组成，job 由很多的 task 组成。
- master service：一个 rpc 服务提供远程连接 distribute device，实际上是一个 session target，master service 继承了 tensorflow session 的接口，主要的工作是协同 "worker services"。所有的tensorflow server 都继承了 master service。
- worker service：一个 rpc 服务用来执行 tensorflow 图的计算使用本地的 devices，一个worker service 继承 worker_service.proto. 所有的tensorflow server 都继承了 worker service。
- in-graph：整个集群由一个client来构建graph，并且由这个client来提交graph到集群中，其他worker只负责处理梯度计算的任务。
- between-graph：一个集群中多个worker可以创建多个graph，但由于worker运行的代码相同因此构建的graph也相同，并且参数都保存到相同的ps中保证训练同一个模型，这样多个worker都可以构建graph和读取训练数据，适合大数据场景。
- synchronous training：同步训练每次更新梯度需要阻塞等待所有worker的结果。
- asynchronous training：异步训练不会有阻塞，训练的效率更高，在大数据和分布式的场景下一般使用异步训练。

## example 
```python
import argparse
import sys

import tensorflow as tf

FLAGS = None

def main(_):
  ps_hosts = FLAGS.ps_hosts.split(",")
  worker_hosts = FLAGS.worker_hosts.split(",")

  # Create a cluster from the parameter server and worker hosts.
  cluster = tf.train.ClusterSpec({"ps": ps_hosts, "worker": worker_hosts})

  # Create and start a server for the local task.
  server = tf.train.Server(cluster,
                           job_name=FLAGS.job_name,
                           task_index=FLAGS.task_index)

  if FLAGS.job_name == "ps":
    server.join()
  elif FLAGS.job_name == "worker":

    # Assigns ops to the local worker by default.
    with tf.device(tf.train.replica_device_setter(
        worker_device="/job:worker/task:%d" % FLAGS.task_index,
        cluster=cluster)):

      # Build model...
      loss = ...
      global_step = tf.contrib.framework.get_or_create_global_step()

      train_op = tf.train.AdagradOptimizer(0.01).minimize(
          loss, global_step=global_step)

    # The StopAtStepHook handles stopping after running given steps.
    hooks=[tf.train.StopAtStepHook(last_step=1000000)]

    # The MonitoredTrainingSession takes care of session initialization,
    # restoring from a checkpoint, saving to a checkpoint, and closing when done
    # or an error occurs.
    with tf.train.MonitoredTrainingSession(master=server.target,
                                           is_chief=(FLAGS.task_index == 0),
                                           checkpoint_dir="/tmp/train_logs",
                                           hooks=hooks) as mon_sess:
      while not mon_sess.should_stop():
        # Run a training step asynchronously.
        # See `tf.train.SyncReplicasOptimizer` for additional details on how to
        # perform *synchronous* training.
        # mon_sess.run handles AbortedError in case of preempted PS.
        mon_sess.run(train_op)

if __name__ == "__main__":
  parser = argparse.ArgumentParser()
  parser.register("type", "bool", lambda v: v.lower() == "true")
  # Flags for defining the tf.train.ClusterSpec
  parser.add_argument(
      "--ps_hosts",
      type=str,
      default="",
      help="Comma-separated list of hostname:port pairs"
  )
  parser.add_argument(
      "--worker_hosts",
      type=str,
      default="",
      help="Comma-separated list of hostname:port pairs"
  )
  parser.add_argument(
      "--job_name",
      type=str,
      default="",
      help="One of 'ps', 'worker'"
  )
  # Flags for defining the tf.train.Server
  parser.add_argument(
      "--task_index",
      type=int,
      default=0,
      help="Index of task within the job"
  )
  FLAGS, unparsed = parser.parse_known_args()
  tf.app.run(main=main, argv=[sys.argv[0]] + unparsed)
```
运行：
```
# On ps0.example.com:
$ python trainer.py \
     --ps_hosts=ps0.example.com:2222,ps1.example.com:2222 \
     --worker_hosts=worker0.example.com:2222,worker1.example.com:2222 \
     --job_name=ps --task_index=0
# On ps1.example.com:
$ python trainer.py \
     --ps_hosts=ps0.example.com:2222,ps1.example.com:2222 \
     --worker_hosts=worker0.example.com:2222,worker1.example.com:2222 \
     --job_name=ps --task_index=1
# On worker0.example.com:
$ python trainer.py \
     --ps_hosts=ps0.example.com:2222,ps1.example.com:2222 \
     --worker_hosts=worker0.example.com:2222,worker1.example.com:2222 \
     --job_name=worker --task_index=0
# On worker1.example.com:
$ python trainer.py \
     --ps_hosts=ps0.example.com:2222,ps1.example.com:2222 \
     --worker_hosts=worker0.example.com:2222,worker1.example.com:2222 \
     --job_name=worker --task_index=1
```


