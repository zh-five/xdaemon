// +build !windows,!plan9

package xdaemon

import "syscall"

func NewSysProcAttr() *syscall.SysProcAttr {
	return &syscall.SysProcAttr{
		Setsid:true,
	}
}