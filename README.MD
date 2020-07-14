
# 工具介绍
该工具用于发送报文

# tcp action
### 介绍
用于向某个服务发送tcp报文，以向百度发送tcp报文为例介绍如何使用
### 通过ping www.baidu.com获取百度的ip
```
>$ ping www.baidu.com
PING www.a.shifen.com (182.61.200.6): 56 data bytes
```

### 通过sktools连接182.61.200.6
```
./sktools -a tcp -host 182.61.200.6 -port 80
```

### 粘贴以下报文
```
GET / HTTP/1.1
Host: www.baidu.com


```
注意上面的空行
### ^]+enter换行
^] 为 control+]
### 获取到百度的响应

> 虽然telnet可以实现该功能，但是telnet只能一行一行的发送，如果想实现发送多行，比较麻烦，故开发该工具

# tcp test action
### 介绍
用于批量向某个服务发送tcp报文
### 示例
```
sktools -a tcptest  -host 127.0.0.1 -port 9000 -c 10 -content abc

```
* c 多少个连接
* content 内容，可为空