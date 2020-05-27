package components

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/logs"
)

var BConfig config.Configer = nil

func init() {
	BConfig, err := config.NewConfig("ini", "config/conf/server.conf")
	if err != nil {
		panic("config init error")
		return
	}

	maxlines, lerr := BConfig.Int64("log::maxlines")
	if lerr != nil {
		maxlines = 1000
	}

	logConf := make(map[string]interface{})
	logConf["filename"] = BConfig.String("log::log_path")
	level, _ := BConfig.Int("log::log_level")
	logConf["level"] = level
	logConf["maxlines"] = maxlines

	confStr, err := json.Marshal(logConf)
	if err != nil {
		fmt.Println("marshal failed,err:", err)
		return
	}
	logs.SetLogger(logs.AdapterFile, string(confStr))
	logs.SetLogFuncCall(true)
}
