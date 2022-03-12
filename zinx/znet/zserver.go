package znet

import (
	"fmt"
	"net"
	"zinxDemo/zinx/ziface"
)

//定义IServer.go 的接口实现,定义一个server的服务器模块
type Server struct {
	//服务器名称
	Name string
	//服务器绑定的ip 版本号
	IPVersion string
	//服务器监听的IP
	IP string
	//服务器监听的端口
	Port int
}

// 启动服务器
func (s *Server) Start() {
	fmt.Printf("[start] IP:%s , Port:%d ,is starting ", s.IP, s.Port)

	go func() {
		//1.获取一个tcp的addr
		// ResolveTCPAddr返回TCP端点的地址
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addt error:", err)
			return
		}

		//2.监听服务器的地址
		listenner, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen ", s.IPVersion, "err :", err)
			return
		}
		fmt.Println("start Zinx server succ,", s.Name, "succ,Listenning...")

		//3.阻塞的等待客户端连接,处理客户端链接业务(读写)
		for true {
			//如果有连接会返回
			conn, err := listenner.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err", err)
				continue
			}
			//已经与客户端建立连接
			go func() {
				for true {
					buf := make([]byte, 512)
					cnt, err := conn.Read(buf) //cnt长度
					if err != nil {
						fmt.Println("recv buf err", err)
						continue
					}
					fmt.Printf("服务端得到:%s,长度:%d \n", buf, cnt)
					//回显功能
					if _, err := conn.Write(buf[0:cnt]); err != nil {
						fmt.Println("write back buf err ", err)
						continue
					}

				}
			}()

		}
	}()

}

// 停止服务器
func (s *Server) Stop() {
	//todo 将一些服务器的资源,状态或者一些已经开辟的链路信息,进行停止或者回收
}

//运行服务器
func (s *Server) Serve() {
	//启动服务器
	s.Start()

	//TODO 做额外业务

	//阻塞状态
	select {}
}

//初始化Server模块的方法
func NewServer(name string) ziface.IServer {
	server := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8999,
	}
	return server

}
