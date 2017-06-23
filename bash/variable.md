- [shell 变量](#vshell-变量)
  - [拼接字符串](#拼接字符串)
  - [获取字符串长度](#获取字符串长度)
  - [提取子字符串](#提取子字符串)
- [shell 数组](#shell-数组)
  - [读取数组](#读取数组)
  - [获取数组的长度](#获取数组的长度)
- [字符串截取](#字符串截取)

## shell 变量

"" 号的优点：
- 双引号里可以有变量
- 双引号里可以出现转义自符

#### 拼接字符串
```bash
your_name="qinjx"
greeting="hello, "$your_name" !"
greeting_1="hello, ${your_name} !"
echo $greeting $greeting_1
```

#### 获取字符串长度
```bash
string="abcd"
echo ${ #string }
```

#### 提取子字符串
```bash
string="runoob is a great site"
echo ${string:1:4} # 输出 unoo
```

## shell 数组
```
数组名=(值1 值2 ... 值n)
```
eg:
```bash
array_name=(value0 value1 value2 value3)
```
#### 读取数组
```bash
valuen=${array_name[n]}
```
使用@符号可以获取数组中的所有元素，例如：
```bash
echo ${array_name[@]}
```
#### 获取数组的长度
获取数组长度的方法与获取字符串长度的方法相同，例如：
```bash
# 取得数组元素的个数
length=${#array_name[@]}
# 或者
length=${#array_name[*]}
# 取得数组单个元素的长度
lengthn=${#array_name[n]}
```
## 字符串截取
假设有变量 var=http://www.aaa.com/123.htm
```bash
1. # 号截取，删除左边字符，保留右边字符。
echo ${var#*//}
其中 var 是变量名，# 号是运算符，*// 表示从左边开始删除第一个 // 号及左边的所有字符
即删除 http://
结果是 ：www.aaa.com/123.htm
```
```bash
2. ## 号截取，删除左边字符，保留右边字符。
echo ${var##*/}
##*/ 表示从左边开始删除最后（最右边）一个 / 号及左边的所有字符
即删除 http://www.aaa.com/
结果是 123.htm
```
```bash
3. %号截取，删除右边字符，保留左边字符
echo ${var%/*}
%/* 表示从右边开始，删除第一个 / 号及右边的字符
结果是：http://www.aaa.com
```
```bash
4. %% 号截取，删除右边字符，保留左边字符
echo ${var%%/*}
%%/* 表示从右边开始，删除最后（最左边）一个 / 号及右边的字符
结果是：http:
```