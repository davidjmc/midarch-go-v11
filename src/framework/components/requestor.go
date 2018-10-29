package components

import (
	"framework/messages"
)

type Requestor struct{}

func (Requestor) I_PosInvP(msg *messages.SAMessage) {
	requestHeader := messages.RequestHeader{Operation: msg.Payload.(messages.Invocation).Op}
	requestBody := messages.RequestBody{Args: msg.Payload.(messages.Invocation).Args}
	miopHeader := messages.MIOPHeader{Magic: "MIOP"}
	miopBody := messages.MIOPBody{RequestHeader: requestHeader, RequestBody: requestBody}
	miop := messages.MIOP{Header: miopHeader, Body: miopBody}
	toCRH := messages.ToCRH{Host: msg.Payload.(messages.Invocation).Host, Port: msg.Payload.(messages.Invocation).Port, MIOP: miop}

	*msg = messages.SAMessage{toCRH}
}

func (Requestor) I_PosTerR(msg *messages.SAMessage) {
	payload := msg.Payload.(map[string]interface{})

	miopBody := payload["Body"]
	miopReplyBody := miopBody.(map[string]interface{})
	reply := miopReplyBody["ReplyBody"]
	msgTemp := messages.SAMessage{reply}
	*msg = msgTemp
}
