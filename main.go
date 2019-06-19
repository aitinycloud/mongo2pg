// Project  : mongo2pg
// package :  main
// file    :  main.go
// Copyright (c) 2018-2019
package main

import (
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"./handle"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	handle.Setup()
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, os.Kill, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-sc
}
