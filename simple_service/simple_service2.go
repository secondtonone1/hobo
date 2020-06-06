package simple_service

import (
	"fmt"
	"time"

	"github.com/duanhf2012/origin/node"
	"github.com/duanhf2012/origin/service"
)

func init() {
	node.Setup(&Simple_Service2{})
}

type Simple_Service2 struct {
	service.Service
}

func (ss *Simple_Service2) OnInit() error {
	fmt.Println("simple service2 oninit success")
	ss.AfterFunc(time.Second*1, ss.onSecondTick)
	return nil
}

//设置定时器
func (ss *Simple_Service2) onSecondTick() {
	fmt.Printf("tick.\n")
	ss.AfterFunc(time.Second*1, ss.onSecondTick)
}
