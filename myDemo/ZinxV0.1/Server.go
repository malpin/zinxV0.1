package main

import "zinxDemo/zinx/znet"

func main() {
	//创建一个server句柄 使用zinx的api
	s := znet.NewServer("[zinx v1.0]")
	//启动server
	s.Serve()
}
