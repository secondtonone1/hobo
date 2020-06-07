package mulgo_service

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/duanhf2012/origin/node"
	"github.com/duanhf2012/origin/service"
)

func init() {
	node.Setup(&MulGo_Service{rwlock: sync.RWMutex{}})
}

type MulGo_Service struct {
	service.Service
	rwlock sync.RWMutex
}

//多协程模式，并发, 要注意该MulGo_Service的所有方法是协程危险的，所以要加锁
func (ss *MulGo_Service) OnInit() error {
	fmt.Println("mulgo service oninit success")
	logs.Debug("beego log: mulgo service oninit success")
	ss.AfterFunc(time.Second*1, ss.OnSecondTick)
	ss.SetGoRouterNum(10)
	return nil
}

func (ss *MulGo_Service) goID() uint64 {
	ss.rwlock.Lock()
	defer ss.rwlock.Unlock()
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	fmt.Println("current goroutine id is %d", n)
	return n
}

//设置定时器,不同的协程都会有自己的定时器，打印自己的协程id
func (ss *MulGo_Service) OnSecondTick() {
	ss.goID()
	ss.AfterFunc(time.Second*1, ss.OnSecondTick)
}
