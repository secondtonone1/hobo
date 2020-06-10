package rpc_service

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/logs"

	"github.com/duanhf2012/origin/node"
	"github.com/duanhf2012/origin/service"
)

type RPC_Call struct {
	service.Service
}

func init() {
	node.Setup(&Simple_RPC{})
	node.Setup(&RPC_Call{})
}

func (rc *RPC_Call) OnInit() error {
	rc.AfterFunc(time.Second*2, rc.AsynCall)
	//rc.AfterFunc(time.Second*2, rc.SyncCall)
	//rc.AfterFunc(time.Second*2, rc.GoCall)
	return nil
}

//同步调用
func (rc *RPC_Call) SyncCall() {
	var input InPut
	input.Str1 = "Hello "
	input.Str2 = "World !"
	input.N1 = 100
	input.N2 = 200
	var output OutPut
	err := rc.Call("Simple_RPC.RPC_StrJoin", &input, &output)
	if err != nil {
		logs.Debug("sync call Simple_RPC.RPC_StrJoin failed!")
		return
	}
	fmt.Printf("sync call Simple_RPC.RPC_StrJoin success, output is %v\n", output)
	logs.Debug("sync call Simple_RPC.RPC_StrJoin success, output is %v", output)
}

//异步调用
func (rc *RPC_Call) AsynCall() {
	//异步调用支持传入一个函数，当异步调用结束时，调用这个函数，这个函数两个参数
	//第一个参数为RPC函数的第二个参数，第二个参数为RPC函数的返回值
	var input InPut
	input.Str1 = "Hello "
	input.Str2 = "World !"
	input.N1 = 100
	input.N2 = 200
	//fmt.Printf("input addr is %p\n", &input)
	//可以指定节点，表示异步调用哪个节点的服务
	/*slf.AsyncCallNode(1,"TestService6.RPC_Sum",&input,func(output *int,err error){
	})*/
	//不指定节点，就是调用同节点的这个服务
	rc.AsyncCall("Simple_RPC.RPC_StrJoin", &input, func(output *OutPut, err error) {
		if err != nil {
			fmt.Printf("AsyncCall error :%+v\n", err)
		} else {
			fmt.Printf("AsyncCall output %v\n", output)
		}
	})
}

//非阻塞派发消息
func (rc *RPC_Call) GoCall() {
	var input InPut
	input.Str1 = "Hello "
	input.Str2 = "World !"
	input.N1 = 100
	input.N2 = 200
	//非阻塞调用，类似于消息派发，只传递输入参数即可，不关心返回值和结果
	err := rc.Go("Simple_RPC.RPC_StrJoin", &input)
	if err != nil {
		logs.Debug("Go call error")
	}

}
