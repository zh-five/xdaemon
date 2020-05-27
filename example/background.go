//本示例, 将把进程转为后台运行, 并保留所有参数不变

package main

import (
	"github.com/zh-five/damon"
	"log"
	"os"
)

func main() {
	logFile := "/tmp/daemon.log"

	//启动一个子进程后主程序退出
	daemon.Background(logFile, true)

	//以下代码只有子程序会执行
	log.Println(os.Getpid(), ":", "执行正式代码:")
}