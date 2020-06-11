package main

import (
	_ "github.com/secondtonone1/hobo/event_service"
	_ "github.com/secondtonone1/hobo/module_service"
	_ "github.com/secondtonone1/hobo/mulgo_service"
	_ "github.com/secondtonone1/hobo/profile_service"
	_ "github.com/secondtonone1/hobo/rpc_service"
	_ "github.com/secondtonone1/hobo/simple_event"
	_ "github.com/secondtonone1/hobo/simple_http"
	_ "github.com/secondtonone1/hobo/simple_service"
	_ "github.com/secondtonone1/hobo/simple_tcp"
	"time"

	_ "github.com/secondtonone1/hobo/components"

	node "github.com/duanhf2012/origin/node"
)

func main() {
	//打开性能分析报告功能，并设置10秒汇报一次
	node.OpenProfilerReport(time.Second * 10)
	node.Start()
}
