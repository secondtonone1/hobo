package main

import (
	_ "hobo/mulgo_service"
	_ "hobo/profile_service"
	_ "hobo/simple_service"
	"time"

	_ "hobo/components"

	node "github.com/duanhf2012/origin/node"
)

func main() {
	//打开性能分析报告功能，并设置10秒汇报一次
	node.OpenProfilerReport(time.Second * 10)
	node.Start()
}
