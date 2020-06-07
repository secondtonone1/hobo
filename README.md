# hobo
a game server based on origin
## 感谢origin库
感谢duanhf2012提供的开源网络框架，连接地址如下
https://github.com/duanhf2012/origin
本项目基于origin做应用和扩充，感谢原作者origin支持。 
## 文件结构
main.go为入口程序，也是启动程序
simple_service文件夹下实现了两个简单的service，simple_service2开启了定时器
mulgo_service文件夹下mulgo_service.go设置了多协程处理
profile_service文件夹下profile_service.go设置了性能检测
module_service文件夹下module_service.go 向一个service中加入多个module

## 如何启动
go build hobo/main.go 生成main可执行文件
main start nodeid=1 即可启动节点1