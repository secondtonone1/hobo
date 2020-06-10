package simple_event

import (
	"time"

	"github.com/duanhf2012/origin/event"
	"github.com/duanhf2012/origin/node"
	"github.com/duanhf2012/origin/service"
)

const (
	//自定义事件类型，必需从event.Sys_Event_User_Define开始
	//event.Sys_Event_User_Define以内给系统预留
	EVENT1 event.EventType = event.Sys_Event_User_Define + 2
)

func init() {
	node.Setup(&TestService4{})
}

type TestService4 struct {
	service.Service
}

func (slf *TestService4) OnInit() error {
	//10秒后触发广播事件
	slf.AfterFunc(time.Second*10, slf.TriggerEvent)
	return nil
}

func (slf *TestService4) TriggerEvent() {
	//广播事件，传入event.Event对象，类型为EVENT1,Data可以自定义任何数据
	//这样，所有监听者都可以收到该事件
	slf.GetEventHandler().NotifyEvent(&event.Event{
		Type: EVENT1,
		Data: "event data.",
	})
}
