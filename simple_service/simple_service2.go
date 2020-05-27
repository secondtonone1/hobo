package simple_service

import (
	"fmt"

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

	return nil
}

//设置定时器
func (ss *Simple_Service2) onSecondTick() {

}
