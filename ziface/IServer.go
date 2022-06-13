package ziface

type IServer interface {
	//stats server
	Start()
	//stop the server
	Stop()
	// run the server
	Server()

	// 路由模块
	AddRouter(router IRouter)
}
