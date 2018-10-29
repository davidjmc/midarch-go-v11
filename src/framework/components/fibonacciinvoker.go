package components

import (
	"framework/messages"
	"fmt"
	"apps/fibonacci/fibonacci"
)

type FibonacciInvoker struct{}

func (FibonacciInvoker) I_PosInvP(msg *messages.SAMessage) {
	op := msg.Payload.(messages.MIOP).Body.RequestHeader.Operation

	switch op {
	case "Fibo":
		// process request
		_args := msg.Payload.(messages.MIOP).Body.RequestBody.Args
		_argsX := _args.([]interface{})
		_p1 := int(_argsX[0].(float64))
		_r := fibonacci.Fibonacci{}.Fibo(_p1) // dispatch

		// send reply
		_replyHeader := messages.ReplyHeader{Status: 1} // 1 - Success
		_replyBody := messages.ReplyBody{Reply: _r}
		_miopHeader := messages.MIOPHeader{Magic: "MIOP"}
		_miopBody := messages.MIOPBody{ReplyHeader: _replyHeader, ReplyBody: _replyBody}
		_miop := messages.MIOP{Header: _miopHeader, Body: _miopBody}
		*msg = messages.SAMessage{_miop}
	default:
		fmt.Println("FIBONACCIINVOKER:: Operation " + op + " not supported")
	}
}
