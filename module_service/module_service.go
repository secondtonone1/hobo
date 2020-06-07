package module_service

import (
	"fmt"
	"sync"
	"time"

	"github.com/duanhf2012/origin/node"
	"github.com/duanhf2012/origin/service"
)

func init() {
	node.Setup(&Module_Service{rwlock: sync.RWMutex{}})
}

type Module_Service struct {
	service.Service
	rwlock sync.RWMutex
}

type Module1 struct {
	service.Module
}

func (slf *Module1) OnInit() error {
	fmt.Printf("Module1 OnInit.\n")
	return nil
}

func (slf *Module1) OnRelease() {
	fmt.Printf("Module1 Release.\n")
}

type Module2 struct {
	service.Module
}

func (slf *Module2) OnInit() error {
	fmt.Printf("Module2 OnInit.\n")
	return nil
}

func (slf *Module2) OnRelease() {
	fmt.Printf("Module2 Release.\n")
}

func (slf *Module_Service) OnInit() error {
	fmt.Printf("TestService1 OnInit.\n")
	//打开性能分析工具
	slf.OpenProfiler()
	//监控超过1秒的慢处理
	slf.GetProfiler().SetOverTime(time.Second * 1)
	//监控超过10秒的超慢处理，您可以用它来定位是否存在死循环
	//比如以下设置10秒，我的应用中是不会发生超过10秒的一次函数调用
	//所以设置为10秒。
	slf.GetProfiler().SetMaxOverTime(time.Second * 10)

	slf.AfterFunc(time.Second*2, slf.Loop)
	//打开多线程处理模式，10个协程并发处理
	//slf.SetGoRouterNum(10)
	return nil
}

func (slf *Profile_Service) Loop() {
	for {
		time.Sleep(time.Second * 1)
	}
}
