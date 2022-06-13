package ziface

/*
	IResquest接口:
	实际上是把客户端请求的链接信息和请求的数据,包装到一个Request中
*/

type IResquest interface {
	// 得到当前链接
	GetConnection() IConnection
	//得到当前的消息数据
	GetData() []byte
}
