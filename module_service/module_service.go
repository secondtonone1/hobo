package module_service

import (
	"fmt"
	"sync"

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
	fmt.Printf("Module_Service OnInit.\n")
	//新建两个Module对象
	module1 := &Module1{}
	module2 := &Module2{}
	//将module1添加到服务中
	module1Id, _ := slf.AddModule(module1)
	//在module1中添加module2模块
	module1.AddModule(module2)
	fmt.Printf("module1 id is %d, module2 id is %d", module1Id, module2.GetModuleId())

	//释放模块module1
	slf.ReleaseModule(module1Id)
	fmt.Printf("xxxxxxxxxxx")
	return nil
}
