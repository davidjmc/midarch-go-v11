package components

import (
	"shared/parameters"
	"fmt"
	"shared/shared"
	"os"
	"framework/messages"
)

type NotificationEngine struct {}

var Queues = map[string]chan messages.MessageMOM{}
var Subscribers = map[string][]string{}   // TODO

func (NE NotificationEngine) I_PosInvP(msg *messages.SAMessage){
	inv := msg.Payload.(messages.Invocation)

	switch inv.Op {
	case "Subscribe":
		_args := inv.Args.([]interface{})
		_topic := _args[0].(string)
		_ip    := _args[1].(string)
		_r := NE.Subscribe(_topic,_ip)

		_ter := shared.QueueingTermination{_r}
		*msg = messages.SAMessage{_ter}
	case "Unsubscribe":

	case "Publish":
		_args := inv.Args.([]interface{})
		_topic := _args[0].(string)
		_msg := _args[1].(map[string]interface{})
		_msgHeader := _msg["Header"].(map[string]interface{})
		_headerDestination := _msgHeader["Destination"].(string)
		_msgPayload := _msg["PayLoad"].(string)
		_r := NE.Publish(_topic,messages.MessageMOM{Header:messages.Header{Destination:_headerDestination},PayLoad:_msgPayload})

		_ter := shared.QueueingTermination{_r}
		*msg = messages.SAMessage{_ter}
	case "Consume":
		_args := inv.Args.([]interface{})
		_topic := _args[0].(string)
		_r := NE.Consume(_topic)

		_ter := shared.QueueingTermination{_r}
		*msg = messages.SAMessage{_ter}

	default:
		fmt.Println("QueueingInvoker:: Operation " + inv.Op + " is not implemented by Queueing Server")
		os.Exit(0)
	}
}


func (NotificationEngine) Subscribe(topic string, ip string) bool {
	r := false

	if _, ok := Subscribers[topic]; !ok {
		Subscribers[topic] = []string{}
	}

	Subscribers[topic] = append(Subscribers[topic], ip)

	fmt.Println(Subscribers)

	return r
}

func (NotificationEngine) Publish(topic string, msg messages.MessageMOM) bool {
	r := false

	if _, ok := Queues[topic]; !ok {
		Queues[topic] = make(chan messages.MessageMOM, parameters.QUEUE_SIZE)
	}

	if len(Queues[topic]) < parameters.QUEUE_SIZE {
		Queues[topic] <- msg
		r = true
	} else {
		r = false
	}
	return r
}

func (NotificationEngine) Consume(topic string) messages.MessageMOM {
	r := messages.MessageMOM{}
	if _, ok := Queues[topic]; !ok {
		Queues[topic] = make(chan messages.MessageMOM, parameters.QUEUE_SIZE)
	}
	if len(Queues[topic]) == 0 {
		r = messages.MessageMOM{Header:messages.Header{Destination:topic},PayLoad:"QUEUE EMPTY"} // TODO
	} else {
		r = <-Queues[topic]
	}
	return r
}