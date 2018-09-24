package fibonacciclientproxy

import (
	"framework/message"
	"shared/parameters"
)

type FibonacciClientProxy struct {
	Host  string
	Port  int
	Proxy string
}

var chIn = make(chan message.Message, parameters.CHAN_BUFFER_SIZE)
var chOut = make(chan message.Message, parameters.CHAN_BUFFER_SIZE)

func (c FibonacciClientProxy) Fibo(p1 int) int {

	// configure parameters
	args := []int{p1}
	inv := message.Invocation{Host: c.Host, Port: c.Port, Op: "fibo", Args: args}
	reqMsg := message.Message{inv}

	// send invocation to the requestor
	chIn <- reqMsg

	// receive response from the requestor
	repMsg := <-chOut

	// configure reply to the client
	payload := repMsg.Payload.(map[string]interface{})
	reply := int(payload["Reply"].(float64))

	return reply
}

func (FibonacciClientProxy) I_PreInvR(msg *message.Message) {
	*msg = <-chIn
}

func (FibonacciClientProxy) I_PosTerR(msg *message.Message) {
	chOut <- *msg
}