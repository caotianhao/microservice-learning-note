package main

import (
	"fmt"
	"net/rpc/jsonrpc"
	"time"
)

func main() {
	go func() {
		conn, _ := jsonrpc.Dial("tcp", "192.168.235.133:8024")
		var result int
		_ = conn.Call("myRPCSevName.Add", []int{1523, 45, 2, 3}, &result)
		fmt.Println("成功！", result)
	}()
	time.Sleep(time.Second)
}
