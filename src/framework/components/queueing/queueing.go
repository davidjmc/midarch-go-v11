package queueing

import (
	"framework/message"
	"framework/components/naming/namingclientproxy"
	"fmt"
)

type QueueingService struct{}

//var Repo = map[string]ior.IOR{}

func (qs QueueingService) I_PosInvP(msg *message.Message) {

	// recover parameters
	//op := msg.Payload.(message.Invocation).Op
	//args := msg.Payload.(message.Invocation).Args

	//switch op {
	//case "publish":
	//	fmt.Println("Naming:: Register")

	fmt.Println("Queueing:: HERE")
}

func LocateNaming(host string, port int) namingclientproxy.NamingClientProxy {
	p := namingclientproxy.NamingClientProxy{Host: host, Port: port}
	return p
}
