package znet

import (
	"fmt"
	"net"
	"wkzinx/ziface"
)

type Server struct {
	//服务器名称
	Name string
	//服务器绑定的ip版本
	IPVersion string
	//服务器监听IP
	IP string
	//服务器监听的端口
	Port int
}

//开始方法
func (s *Server) Start() {
	//1.获取一个tcp的addr
	fmt.Printf("[Start] Server listening at IP: %s,Port:%d is starting\n", s.IP, s.Port)

	go func() {
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("ResolveTCPAddr err:", err)
			return
		}
		//2.监听服务器的地址
		listen, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("ListenTCP err:", err)
			return
		}
		fmt.Println("start Zinx server run succ:", s.Name)
		//3.阻塞的等待客户端链接,处理客户端链接业务(读写)
		for {
			//客户端链接过来，阻塞会返回
			conn, err := listen.AcceptTCP()
			if err != nil {
				fmt.Println("accept tcp err", err)
				continue
			}

			//已经链接客户端，做一些业务，最基础的最大512字节回显业务
			go func() {
				for {
					buf := make([]byte, 512)
					cnt, err := conn.Read(buf)
					if err != nil {
						fmt.Println("conn.Read err", err)
						continue
					}
					//回显
					if _, err := conn.Write(buf[:cnt]); err != nil {
						fmt.Println("Write buf err", err)
						continue
					}
				}
			}()
		}
	}()

}

//停止服务
func (s *Server) Stop() {
	//TODO 将一些服务器的资源，状态或者一些已经开辟的链接信息释放

}

//定义服务器
func (s *Server) Server() {
	s.Start()

	//TODO 做一些启动服务器之后的额外业务

	//阻塞状态
	select {}

}

/*
	初始化Server
*/
func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8999,
	}
	return s
}
