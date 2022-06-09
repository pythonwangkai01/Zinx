package ziface

type IServer interface {
	//stats server
	Start()
	//stop the server
	Stop()
	// run the server
	Server()
}
