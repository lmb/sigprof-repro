SIGPROF causes infinite loop
---

This repo reproduces an interaction between the Go runtime CPU profiler and
the BPF_PROG_LOAD syscall on Linux.

Set SetCPUProfileRate(1000) causes the thread executing the syscall to be
interrupted every ~1 ms. This in turn leads to BPF_PROG_LOAD to never finish,
given a complex enough BPF program.

test_cls_redirect.o is a precompiled version of [test_cls_redirect.c](https://elixir.bootlin.com/linux/v5.19.9/source/tools/testing/selftests/bpf/progs/test_cls_redirect.c)
from the Linux kernel sources.

```
$ go run -exec sudo . # sudo is required to allow the BPF syscall
```
