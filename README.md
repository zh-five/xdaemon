# damon
A library for writing system daemons in golang.
一个让go程序快速后台运行的库.
支持linux和windows

# 两种运行模式
## 1.纯后台进程模式
请参考 example/background.go
```go
//本示例, 将把进程转为后台运行, 并保留所有参数不变

package main

import (
	"github.com/zh-five/xdaemon"
	"log"
	"os"
	"time"
)

func main() {
	logFile := "daemon.log"

	//启动一个子进程后主程序退出
	xdaemon.Background(logFile, true)

	//以下代码只有子程序会执行
	log.Println(os.Getpid(), "start...")
	time.Sleep(time.Second * 10)
	log.Println(os.Getpid(), "end")
}
```
## 2.守护进程模式
- 运行主进程时, 启动一个守护进程后退出
- 守护进程启动一个最终的子进程
- 若最终子进程退出, 守护进程将尝试再次启动
- 支持最大重启次数等一些参数的设置
- 最终系统中会存在两个进程:守护进程和最终子进程

请参考 example/auto_restart.go
```go
//本示例, 将把进程转为后台运行, 并保留所有参数不变

package main

import (
	"github.com/zh-five/xdaemon"
	"log"
	"os"
	"time"
)

func main() {
	logFile := "daemon.log"

	//启动一个子进程后主程序退出
	xdaemon.Background(logFile, true)

	//以下代码只有子程序会执行
	log.Println(os.Getpid(), "start...")
	time.Sleep(time.Second * 10)
	log.Println(os.Getpid(), "end")
}

```

## 3.本次开发过程的博客记录
[https://zhuanlan.zhihu.com/p/146192035](https://zhuanlan.zhihu.com/p/146192035)