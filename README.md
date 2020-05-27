# hobo
a game server based on origin
## 感谢origin库
感谢duanhf2012提供的开源网络框架，连接地址如下
https://github.com/duanhf2012/origin
本项目基于origin做应用和扩充，感谢原作者origin支持。 
## 文件结构
main.go为入口程序，也是启动程序
simple_service文件夹下实现了两个简单的service
## 如何启动
go build hobo/main.go 生成main可执行文件
main start nodeid=1 即可启动节点1