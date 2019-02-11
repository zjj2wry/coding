# coding=utf-8
import tensorflow as tf
import numpy as np

# argmax 返回一个数组最大值所在的坐标
# 第二个参数 axis 可以传 0 或者 1, 0 表示纵向，1 表示横向
test = np.array([[1, 2, 3], [2, 3, 4], [5, 4, 3], [8, 7, 2]])
print(np.argmax(test, 0))
print(np.argmax(test, 1))
eq = tf.equal(tf.argmax(test, 0),tf.argmax(test, 0))
cast = tf.cast(eq, tf.float32)
accuracy = tf.reduce_sum(cast)

with tf.Session() as sess:
    print(sess.run(tf.argmax(test, 0)))
    print(sess.run(tf.argmax(test, 1)))
    print(sess.run(eq))
    print(sess.run(cast))
    print(sess.run(accuracy))

# result:
# [3 3 1]
# [2 2 0 0]
# [ True  True  True]
# [1. 1. 1.]
# 3.0