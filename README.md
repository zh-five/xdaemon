# damon
A library for writing system daemons in golang.
一个让go程序快速后台运行的库

# 两种运行模式
## 1.纯后台进程模式
```go
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
```
## 2.守护进程模式
- 运行主进程时, 启动一个守护进程后退出
- 守护进程启动一个最终的子进程
- 若最终子进程退出, 守护进程将尝试再次启动
- 支持最大重启次数等一些参数的设置
- 最终系统中会存在两个进程:守护进程和最终子进程

```go
//本示例, 将启动一个后台运行的守护进程. 然后由守护进程启动和维护最终子进程

package main

import (
	"github.com/zh-five/damon"
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
		logFile := "/tmp/daemon.log"
		d := daemon.NewDaemon(logFile)

		//调整一些运行参数(可选)
		//d.MaxCount = 2000

		//执行
		d.Run()
	}

	//当 *d = true 时以下代码只有最终子进程会执行, 主进程和守护进程都不会执行
	log.Println(os.Getpid(), "start...")
	time.Sleep(time.Second * 10)
	log.Println(os.Getpid(), "end")

}

```