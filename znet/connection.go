package znet

import (
	"Zinx/ziface"
	"fmt"
	"net"
)

/*
  链接模块
*/

type Connection struct {
	//当前链接的socket TCP 套接字
	Conn *net.TCPConn

	//链接的ID
	ConnID uint32

	//当前的链接状态
	isClosed bool

	//告知当前链接已经退出的/停止 channel
	ExitChan chan bool

	//该链接处理的方法Router
	Router ziface.IRouter
}

//初始化链接模块的方法
func NewConnection(conn *net.TCPConn, ConnID uint32, router ziface.IRouter) *Connection {
	c := &Connection{
		Conn:     conn,
		ConnID:   ConnID,
		Router:   router,
		isClosed: false,
		ExitChan: make(chan bool, 1),
	}
	return c
}

func (c *Connection) StartReader() {
	fmt.Println("Reader Goroutine is running...")
	defer fmt.Println("connid=", c.ConnID, "Reader is exit,remoteAddr is", c.RemoteAddr().String())
	defer c.Stop()

	for {
		//读取客户端的数据到buf中，最大512字节
		buf := make([]byte, 512)
		_, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("recv buf err", err)
			continue
		}
		/*
			//调用当前链接所绑定的HandleAPI
			if err := c.handleAPI(c.Conn, buf, cnt); err != nil {
				fmt.Println("Connid", c.ConnID, "handle is err", err)
				break
			}
		*/
		//得到当前Conn数据的request请求数据
		req := Request{
			conn: c,
			data: buf,
		}

		//执行注册的路由方法
		go func(request ziface.IResquest) {
			c.Router.PreHandle(request)
			c.Router.Handle(request)
			c.Router.PostHandle(request)
		}(&req)

	}

}

// 启动链接，让当前的链接准备开始工作
func (c *Connection) Start() {
	fmt.Println("conn start() ... connid=", c.ConnID)
	//启动从当前链接的读取数据的业务
	go c.StartReader()
}

// 停止链接 结束当前链接的工作
func (c *Connection) Stop() {
	fmt.Println("conn stop ... connid=", c.ConnID)
	//if c.isClosed == true
	if c.isClosed == true {
		return
	}
	c.isClosed = true
	//执行关闭
	c.Conn.Close()
	close(c.ExitChan)

}

// 获取当前的链接的绑定socket conn
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

// 获取当前链接模块的链接ID
func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

// 获取远程客户端的 TCP状态 IP Port
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

// 发送数据，将数据发送给远程的客户端
func (c *Connection) Send(data []byte) error {
	return nil
}
