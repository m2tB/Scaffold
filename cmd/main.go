package main

import (
	"GhortLinks/internal/initialize/icommon"
	"GhortLinks/internal/startup"
	"GhortLinks/utils/iarray"
	"flag"
	"fmt"
	"os"
)

var (
	system  = flag.String("s", "", "The `system` of project: 'console' or 'service'")
	runtime = flag.String("r", "", "The `runtime` of project: 'testing' or 'develop' or 'product'")
)

func main() {
	// 解析命令行参数 - 非法参数直接退出
	flag.Usage = func() {
		_, _ = fmt.Fprintln(os.Stdout, "GhortLinks - Startup parameters command options:")
		flag.PrintDefaults()
	}
	flag.Parse()
	active := iarray.CheckStringInArray(*system, icommon.SYSTEM_MAP) && iarray.CheckStringInArray(*runtime, icommon.RUNTIME_MAP)
	if !active {
		_, _ = fmt.Fprintln(os.Stdout, "GhortLinks - Startup parameters are missing, please confirm through the command line -h")
		os.Exit(1)
	}
	// 启动对应服务
	icommon.CURRENT_SYSTEM_IS_SERVICE = *system == "service"
	icommon.CURRENT_RUNTIME = *runtime
	if icommon.CURRENT_SYSTEM_IS_SERVICE {
		startup.Service()
	}
	startup.Console()
}
