package event_service

import (
	"fmt"

	"github.com/duanhf2012/origin/event"
	"github.com/duanhf2012/origin/node"
	"github.com/duanhf2012/origin/service"
)

func init() {
	node.Setup(&Event_Service2{})
	node.Setup(&Event_Service1{})
}

type Event_Service1 struct {
	service.Service
}

type Event_Module struct {
	service.Module
}

func (slf *Event_Module) OnInit() error {

	//在当前node中查找TestService4
	pService := node.GetService("Event_Service2")

	//在TestModule中，往TestService4中注册EVENT1类型事件监听
	pService.(*Event_Service2).GetEventProcessor().RegEventReciverFunc(EVENT1, slf.GetEventHandler(), slf.OnModuleEvent)
	return nil
}

func (slf *Event_Module) OnModuleEvent(ev *event.Event) {
	fmt.Printf("OnModuleEvent type :%d data:%+v\n", ev.Type, ev.Data)
}

//服务初始化函数，在安装服务时，服务将自动调用OnInit函数
func (slf *Event_Service1) OnInit() error {
	fmt.Println("event service1 begin init")
	//通过服务名获取服务对象
	pService := node.GetService("Event_Service2")

	////在TestModule中，往TestService4中注册EVENT1类型事件监听
	pService.(*Event_Service2).GetEventProcessor().RegEventReciverFunc(EVENT1, slf.GetEventHandler(), slf.OnServiceEvent)
	slf.AddModule(&Event_Module{})
	return nil
}

func (slf *Event_Service1) OnServiceEvent(ev *event.Event) {
	fmt.Printf("OnServiceEvent type :%d data:%+v\n", ev.Type, ev.Data)
}
