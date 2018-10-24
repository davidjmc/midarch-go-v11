package crh

import (
	"encoding/json"
	"framework/message"
	"net"
	"strings"
	"strconv"
	"fmt"
	"shared/errors"
)

type CRH struct {
	Port int
}

var conn net.Conn
var err error

var portTmp int

func (e CRH) Loop(channels map[string] chan message.Message) {
	var msgPosInvP, msgPreTerP message.Message
	for {
		select {
		case <-channels["InvP"]:
		case msgPosInvP = <-channels["I_PosInvP"]:
			e.I_PosInvP(&msgPosInvP)
		case msgPreTerP = <-channels["I_PreTerP"]:
			e.I_PreTerP(&msgPreTerP)
		case channels["TerP"] <- msgPreTerP:
		}
	}
}

func (c CRH) I_PosInvP(msg *message.Message) {

	host := msg.Payload.(message.ToCRH).Host
	port := msg.Payload.(message.ToCRH).Port
	addr := strings.Join([]string{host, strconv.Itoa(port)}, ":")
	conn, err = net.Dial("tcp", addr)

	//defer conn.Close()

	portTmp = port
	if err != nil {
		fmt.Println(err)
		myError := errors.MyError{Source: "CRH", Message: "Unable to open connection to " + host + " : " + strconv.Itoa(port)}
		myError.ERROR()
	}

	encoder := json.NewEncoder(conn)
	err = encoder.Encode(msg.Payload.(message.ToCRH).MIOP)
	if err != nil {
		fmt.Println(err)
		myError := errors.MyError{Source: "CRH", Message: "Unable to send data to " + host + ":" + strconv.Itoa(port)}
		myError.ERROR()
	}
}

func (c CRH) I_PreTerP(msg *message.Message) {

	decoder := json.NewDecoder(conn)
	err = decoder.Decode(&msg)

	if err != nil {
		fmt.Println(err)
		myError := errors.MyError{Source: "CRH", Message: "Problem in decoding at Port " + strconv.Itoa(portTmp)}
		myError.ERROR()
	}
	conn.Close()
}
