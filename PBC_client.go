package main

import (
	"fmt"
	"net/rpc/jsonrpc"
	"time"
)

// 这个代码是和我自己本地的 Ubuntu 联动的
func main() {
	go func() {
		conn, _ := jsonrpc.Dial("tcp", "192.168.235.133:8024")
		var result int
		_ = conn.Call("myRPCSevName.Add", []int{1523, 45, 2, 3}, &result)
		fmt.Println("成功！", result)
	}()
	time.Sleep(time.Second)
}
