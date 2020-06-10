package rpc_service

import (
	"fmt"

	"github.com/astaxie/beego/logs"
	"github.com/duanhf2012/origin/service"
)

func init() {

}

type Simple_RPC struct {
	service.Service
}

func (sr *Simple_RPC) OnInit() error {
	logs.Debug("Simple_RPC service init success")
	return nil
}

type InPut struct {
	Str1 string
	Str2 string
	N1   int
	N2   int
}

type OutPut struct {
	Str string
}

//函数名RPC开头，后面自己定义
//两个参数，一个传入，一个传出
func (sr *Simple_RPC) RPC_StrJoin(in *InPut, out *OutPut) error {
	//fmt.Printf("RPC_StrJoin  in addr is %p\n", in)
	fmt.Println("in.str1 is ", in.Str1)
	fmt.Println("in.str2 is ", in.Str2)
	fmt.Println("in.n1 is ", in.N1)
	fmt.Println("in.n2 is ", in.N2)
	out.Str = in.Str1 + in.Str2
	fmt.Println("Simple_RPC out str is ", out.Str)
	return nil
}
