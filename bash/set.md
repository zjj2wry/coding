- [set 命令](#set-命令)
  - [set -o errexit](#set--o-errexit)
  - [set -o nounset](#set--o-nounset)
  - [set -o pipefail](#set--o-pipefail)

## set 命令

k8s 常用的3个 set 命令,set -o 表示打开，set +o 表示关闭：
#### set -o errexit
```解释
脚本中有命令的返回值为非0，则脚本立即退出，后续命令不再执行
```
#### set -o nounset
```解释
脚本中存在没有初始化的变量，会报错并退出。
如果没有设置这个，会默认把变量的值设置为空，脚本中变量名称写错不好排查
```
##### 例子：
```bash
#! /bin/bash
set +o nounset
((ret=ret+$?))
```
##### 输出结果：
```bash
admindeMacBook-Pro:bash admin$ ./test.sh 
./test.sh: line 5: ret: unbound variable
```

#### set -o pipefail
```bash
表示在管道连接的命令序列中，只要有任何一个命令返回非0值，则整个管道返回非0值，即使最后一个命令返回0
```
例子：
```
#!/bin/bash
# testset.sh
echo 'disable exit on non-zero return status and pipefail track'
set +e
set +o pipefail
a=$[1/0]|b=2
echo 'return status = '$?

echo 'disable exit on non-zero return status but enable pipefail track'
set +e
set -o pipefail
a=$[1/0]|b=2
echo 'return status = '$?

echo 'enable exit on non-zero return status and pipefail track'
set -e
set -o pipefail
a=$[1/0]|b=2
echo 'return status = '$?
```
输入结果：
```bash
admindeMacBook-Pro:bash admin$ ./testset.sh 
disable exit on non-zero return status and pipefail track
./testset.sh: line 6: 1/0: division by 0 (error token is "0")
return status = 0
disable exit on non-zero return status but enable pipefail track
./testset.sh: line 12: 1/0: division by 0 (error token is "0")
return status = 1
enable exit on non-zero return status and pipefail track
./testset.sh: line 18: 1/0: division by 0 (error token is "0")
```
