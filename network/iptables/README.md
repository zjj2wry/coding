## iptables

### 查看命令 -L; 常用iptables -l -n -v
```
-n：以数字的方式显示ip，它会将ip直接显示出来，如果不加-n，则会将ip反向解析成主机名。
-v：显示详细信息
-vv
-vvv :越多越详细
-x：在计数器上显示精确值，不做单位换算
--line-numbers : 显示规则的行号
-t nat：显示所有的关卡的信息
```

### 规则管理命令
         -A：追加，在当前链的最后新增一个规则
         -I num : 插入，把当前规则插入为第几条。
            -I 3 :插入为第三条
         -R num：Replays替换/修改第几条规则
            格式：iptables -R 3 …………
         -D num：删除，明确指定删除第几条规则

详情看这里：http://www.zsythink.net/archives/category/%E8%BF%90%E7%BB%B4%E7%9B%B8%E5%85%B3/%E9%98%B2%E7%81%AB%E5%A2%99/