# Concurrent

see [concurrency](https://golang.org/doc/effective_go.html#concurrency) for details.

1. 通过交流来共享，而不是通过共享来交流

    go的并行借鉴的是`Unix pipline`

2. 一句`go xxxx` 等价于在linux中后台执行一个命令(使用 `command &` )！

    这句话真是一句惊醒梦中人啊！的确就是啊！