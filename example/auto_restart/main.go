//本示例, 将启动一个后台运行的守护进程. 然后由守护进程启动和维护最终子进程

package main

import (
	"github.com/zh-five/xdaemon"
	"flag"
	"log"
	"os"
	"time"
)

func main() {
	d := flag.Bool("d", false, "是否后台守护进程方式运行")
	flag.Parse()

	//启动守护进程
	if *d {
		//创建一个Daemon对象
		logFile := "daemon.log"
		d := xdaemon.NewDaemon(logFile)
		//调整一些运行参数(可选)
		d.MaxCount = 2 //最大重启次数

		//执行守护进程模式
		d.Run()
	}

	//当 *d = true 时以下代码只有最终子进程会执行, 主进程和守护进程都不会执行
	log.Println(os.Getpid(), "start...")
	time.Sleep(time.Second * 10)
	log.Println(os.Getpid(), "end")

}
