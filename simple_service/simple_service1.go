package simple_service

import (
	"fmt"

	"github.com/astaxie/beego/logs"
	"github.com/duanhf2012/origin/node"
	"github.com/duanhf2012/origin/service"
)

func init() {
	node.Setup(&Simple_Service1{})
}

type Simple_Service1 struct {
	service.Service
}

func (ss *Simple_Service1) OnInit() error {
	fmt.Println("simple service1 oninit success")
	logs.Debug("beego log: simple service1 oninit success")
	return nil
}
