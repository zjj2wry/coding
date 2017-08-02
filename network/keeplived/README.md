1、通俗易懂的解释：
http://www.cnblogs.com/codebean/archive/2011/07/25/2116043.html

2、配置 demo：
http://linuxvirtualserver.org/docs/ha/keepalived.html

3、配置参数解释
http://outofmemory.cn/wiki/keepalived-configuration

4、基于keepalived 实现VIP转移，lvs，nginx的高可用
http://limian.blog.51cto.com/7542175/1301776

5、keeplived-vip（https://github.com/kubernetes/contrib/tree/master/keepalived-vip）
大致的思路：使用 configmap 保存 vip 和 service 的地址，controller 会监听service 和 endpoint 对象，然后和 configmap 中保存的 service 对比，没有问题会读取 pod，然后获取 node 的信息，得到网卡信息（因为 pod 使用的是hostnetwork），
将数据渲染到 keeplived 的配置中。对于 vip 的添加 和 keeplived 的配置启动、reload 等都是使用二进制实现的。
模版如下：
```
{{ $iface := .iface }}{{ $netmask := .netmask }}

global_defs {
  vrrp_version 3
  vrrp_iptables {{ .iptablesChain }}
}

vrrp_instance vips {
  state BACKUP
  interface {{ $iface }}
  virtual_router_id {{ .vrid }}
  priority {{ .priority }}
  nopreempt
  advert_int 1

  track_interface {
    {{ $iface }}
  }

  {{ if .useUnicast }}
  unicast_src_ip {{ .myIP }}
  unicast_peer { {{ range .nodes }}
    {{ . }}{{ end }}
  }
  {{ end }}

  virtual_ipaddress { {{ range .vips }}
    {{ . }}{{ end }}
  }
}

{{ range $i, $svc := .svcs }}
{{ if eq $svc.LVSMethod "VIP" }}
# VIP Service with no pods: {{ $svc.IP }}
{{ else }}
# Service: {{ $svc.Name }}
virtual_server {{ $svc.IP }} {{ $svc.Port }} {
  delay_loop 5
  lvs_sched wlc
  lvs_method {{ $svc.LVSMethod }}
  persistence_timeout 1800
  protocol {{ $svc.Protocol }}

  {{ range $j, $backend := $svc.Backends }}
  real_server {{ $backend.IP }} {{ $backend.Port }} {
    weight 1
    TCP_CHECK {
      connect_port {{ $backend.Port }}
      connect_timeout 3
    }
  }
{{ end }}
}
{{ end }}
{{ end }}
```

渲染后的配置：
```
global_defs {
  vrrp_version 3
  vrrp_iptables KUBE-KEEPALIVED-VIP
}

vrrp_instance vips {
  state BACKUP
  interface ens3
  virtual_router_id 50
  priority 101
  nopreempt
  advert_int 1

  track_interface {
    ens3
  }


  unicast_src_ip 192.168.16.174
  unicast_peer {
    192.168.16.173
  }


  virtual_ipaddress {
    192.168.18.229
  }
}



# Service: default/echoheaders
virtual_server 192.168.18.229 80 {
  delay_loop 5
  lvs_sched wlc
  lvs_method DR
  persistence_timeout 1800
  protocol TCP


  real_server 192.168.71.238 8080 {
    weight 1
    TCP_CHECK {
      connect_port 8080
      connect_timeout 3
    }
  }

}
```

运行成功后可以在pod或者节点中查看：
```
[root@c5v174 ~]# ipvsadm -Ln
IP Virtual Server version 1.2.1 (size=4096)
Prot LocalAddress:Port Scheduler Flags
  -> RemoteAddress:Port           Forward Weight ActiveConn InActConn
TCP  192.168.18.229:80 wlc persistent 1800
  -> 192.168.71.238:8080          Route   1      0          0
```


Q&A：
1、DR 模式需要设置arp 抑制以及在 lo 设备上添加 vip 地址
```
/*在回环设备上绑定了一个虚拟IP地址，并设定其子网掩码为255.255.255.255，与Director Server上的虚拟IP保持互通*/
ifconfig lo:0 192.168.132.254 broadcast 192.168.132.254 netmask 255.255.255.255 up
route add -host 192.168.132.254 dev lo:0
//禁用本机的ARP请求echo "1" >/proc/sys/net/ipv4/conf/lo/arp_ignore
echo "1" >/proc/sys/net/ipv4/conf/all/arp_ignore
echo "2" >/proc/sys/net/ipv4/conf/lo/arp_announce
echo "2" >/proc/sys/net/ipv4/conf/all/arp_announce
```

2、ds 和 rs 在同一个节点的情况下，存在死循环的问题：
```
配置的输出如下，可以见到有一个地方不一样，就是显示有个地方是 Local。
#  ipvsadm -L -n
ip Virtual Server version 1.2.1 (size=4096)
Prot LocalAddress:Port Scheduler Flags
  -> RemoteAddress:Port           Forward Weight ActiveConn InActConn
FWM  6 rr
  -> 192.168.1.233:3306           Local   1      0          0
  -> 192.168.1.213:3306           route   1      0          0
```
二台机器上的配置都是这样(^-^上面其实有小修改)，在这种架构中，我只要停止我的备份的 keepalived ，就能正常的工作，不然一定有其中一台数据库连接到最后就中断停止在那个地方不动了。通过抓包，大量包被不断的转发来转发去。
我们来详细分析一下。
客户端发送连接的请求到  VIP 的指定的端口
当前的 director 会选择二个 realserver 来转发请求,会传送数据给 localnode 的本机网卡或者备份的机器上指定的那个 MAC 的 eth0  (同时他也做为 realserver) . 正常的 LVS ，这些数据包会被监听这个 VIP 的程序接收。
如果数据包是发送给备份的 director 服务器的 eth0 接口。它会不能正常的被监听指定端口的程序所接收，因为数据包会首先先经过 ip_vs()。
这时，有 50% 的机会，这个包被转发给 。这时会生成标准的回应数据包给客户端。因为能正常的回应客户端，这时 LVS 的功能是正常的。我们想所有的包都直接给监听程序来处理。并不想直接通过 ip_vs() 来转发。
这时还有 50% 的数据包会直接在次转给主 LVS 的 eth0/VIP。
我们不想数据包从备份的 LVS 在次转回去给主 LVS 这样会形成 。
所以我们要让 eth0 上发到给 VIP 的包， 只要不是其它 LVS 发的，才使用 ip_vs() 来处理。
简单来讲：当客户端发送数据包给 VIP 。比如我们的 Director1 (Master 主)这个接口正在工作，这时 LVS 能接收到这个包，然后根据 keepalived 的配置进行 load balance 。这时 Director1 会使用 LVS-DR 的功能给包路由给自己或者 Director2 (Backup)。
这时有个问题，在这个例子中因为我们使用了 keepalived 。这时 Director2 这台是一台 VIP 的备份服务器。这时 keepalived 默认会立即启动使用 ipvsadm 的规则来配置这台服务器怎么样做备份的处理.来使得更快的故障转移。所以这时这些规则这台备份的 Director2 主机都会存在。
这就有问题了。当从 Director1 (Master 主)，比如使用 rr 。会转发大约 50% 的包从 Director1 到  Director2 (Backup)的 3306 的端口。这时因为 Director2 因为这些 LVS-DR 的配置规则会接着给这些包，在做一次 load balance 。又发回去给 Director1，这时会产生一个死的循环。
随着时间的推移，不但不能正常的处理连接，您的服务器也会崩溃，在他们中间或后端不断的反复连接。
解决方案：给进入 eth0 的包打包 mark 的标记，当数据包是发给 VIP:80  并且 MAC 不其它 LVS 服务器的话. 才做个 mark ，这样才会对指定的 fwmark 进行 loadbalance 放入到 LVS 中处理。只要数据包是从任意其它的 MAC 地址(非 LVS 的转发)会被发往 VIP:port， 会不在进行 loadbalanced 而是直接转给后面监听的  demon 程序进行应用的处理。实际就是我们使用 iptables 来对进入的流量设置 MARK.然后配置 keepalived 只处理有 MARK 过的流量。不在使用以前绑定 VIP 和端口。
iptables 的配置如下：
同时服务于 LVS-DR，又要做为数据库的后端。所以我们要注意，只接收一个 director 的数据包。
这时我们在 Director1 中设置($MAC_Director2 是指我在  Director1 上所能见到从  Director2 发过来包的 MAC 地址) ：
iptables -t mangle -I PREROUTING -d $VIP -p tcp -m tcp --dport $VPORT -m mac  ! --mac-source $MAC_Director2 -j MARK --set-mark 0x3
并在备份的 keepalived 的服务器 Director2 中设置：
iptables -t mangle -I PREROUTING -d $VIP -p tcp -m tcp --dport $VPORT -m mac  ! --mac-source $MAC_Director1 -j MARK --set-mark 0x4
 接着在 keepalived 中分别配置这二个。
Director1: virtual_server fwmark 3 {
Director2: virtual_server fwmark 4 {
其实这个的完整配置象如下：
keepalived 来根据 MARK 来配置的方法
virtual_server fwmark 4  {
    delay_loop 10
    lb_algo rr
    lb_kind DR
    protocol TCP
    real_server 192.168.1.213 3306 {
    weight 1
    MISC_CHECK {
        misc_path "/etc/keepalived/check_slave.pl 192.168.1.213"
        misc_dynamic
    }
    }
    real_server 192.168.1.233 3306 {
    weight 1
    MISC_CHECK {
        misc_path "/etc/keepalived/check_slave.pl 192.168.1.233"
        misc_dynamic
        }
    }
}
