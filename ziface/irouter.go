package ziface

/*
	IRouter抽象接口
	路由里的数据都是IRouter
*/

type IRouter interface {
	//处理Conn业务之前的钩子方法HooK
	PreHandle(request IResquest)
	//处理Conn业务的钩子方法HooK
	Handle(request IResquest)
	//处理Conn业务之后的钩子方法HooK
	PostHandle(request IResquest)
}
