package main

import (
	"runtime"

	"github.com/cilium/ebpf"
)

func main() {
	runtime.SetCPUProfileRate(1000)

	c, err := ebpf.LoadCollection("test_cls_redirect.o")
	if err != nil {
		panic(err)
	}
	c.Close()
}
