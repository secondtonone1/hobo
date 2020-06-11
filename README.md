# hobo 
hobo是基于origin实现的实时socket服务器，支持websocket，tcp，http多种协议，
可用于游戏服务器和IM服务器等即时场景需求较高的业务
## 安装和下载
go get -u -v github.com/secondtonone1/hobo
## 感谢origin库
感谢duanhf2012提供的开源网络框架，连接地址如下
https://github.com/duanhf2012/origin
本项目基于origin做应用和扩充，感谢原作者origin支持。 
## 安装必要的origin库
go get -u -v github.com/duanhf2012/origin
## 文件结构
main.go为入口程序，也是启动程序
simple_service文件夹下实现了两个简单的service，simple_service2开启了定时器
mulgo_service文件夹下mulgo_service.go设置了多协程处理
profile_service文件夹下profile_service.go设置了性能检测
module_service文件夹下module_service.go 向一个service中加入多个module

## 如何启动
go build github.com/secondtonone1/hobo/main.go 生成main可执行文件
main start nodeid=1 即可启动节点1