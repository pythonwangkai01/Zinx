package znet

import "Zinx/ziface"

type Request struct {
	//已经和客户端建立的链接
	conn ziface.IConnection
	//客户端请求的数据
	data []byte
}

func (r *Request) GetConnection() ziface.IConnection {
	return r.conn
}

//得到当前的消息数据
func (r *Request) GetData() []byte {
	return r.data
}
