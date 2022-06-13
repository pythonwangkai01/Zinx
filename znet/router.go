package znet

import "Zinx/ziface"

/*
实现router时，先嵌入这个baseRouter基类，然后根据需求对这个基类的方法进行重写就好了
*/

type BaseRouter struct{}

//处理conn业务之前的钩子方法
func (br *BaseRouter) PreHandle(request ziface.IResquest) {}

//处理conn业务的主方法
func (br *BaseRouter) Handle(request ziface.IResquest) {}

//处理conn业务之后的钩子方法
func (br *BaseRouter) PostHandle(request ziface.IResquest) {}
