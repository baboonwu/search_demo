package main

import (
	"flag"
	"fmt"
	"github.com/ant0ine/go-json-rest/rest"
	"net/http"
	"runtime/debug"
	"tapi_go/route/middleware"
	"tapi_go/tlog"
	"tsearch/config"
	TSearchRoute "tsearch/route"
)

// 命令行参数格式
// eg: ./main -cfg <filename>
var strConfigFile *string = flag.String("cfg", "dev.cfg", "eg: -cfg dev.cfg")

func main() {

	defer func() {
		if err := recover(); err != nil {
			debug.PrintStack()
		}
	}()

	// 解析命令行参数
	flag.Parse()

	// init syslog
	if err := tlog.InitSyslog(); err != nil {
		tlog.Err("tsearch InitSyslog faild, ", err)
		return
	}
	defer tlog.CloseSyslog()

	// 读取配置文件
	if err := config.InitSearchConfig(strConfigFile); err != nil {
		tlog.Err("tsearch InitConfigFromFile faild, ", err)
		return
	}

	api := rest.NewApi()
	var devStack = []rest.Middleware{
		&middleware.AuthFilter{},
		&rest.AccessLogApacheMiddleware{},
		&rest.TimerMiddleware{},
		&rest.RecorderMiddleware{},
		&rest.RecoverMiddleware{},
		&rest.GzipMiddleware{},
	}
	api.Use(devStack...)

	// Register Route
	router, err := rest.MakeRouter(TSearchRoute.GetApiSearch()...)
	if err != nil {
		tlog.Err("MakeRouter", err)
		return
	}

	api.SetApp(router)

	fmt.Printf("tsearch server start. %s \n", config.TSearchConfig.IP)

	err = http.ListenAndServe(config.TSearchConfig.IP, api.MakeHandler())
	if err != nil {
		tlog.Err("ListenAndServe", err)
	}

}
