package components

import (
	"encoding/json"
	"framework/messages"
	"net"
	"strings"
	"strconv"
	"fmt"
	"shared/errors"
)

type CRH struct {
	Port int
}

var portTmp int

func (c CRH) I_PosInvP(msg *messages.SAMessage) {

	host := msg.Payload.(messages.ToCRH).Host
	port := msg.Payload.(messages.ToCRH).Port
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
	err = encoder.Encode(msg.Payload.(messages.ToCRH).MIOP)
	if err != nil {
		fmt.Println(err)
		myError := errors.MyError{Source: "CRH", Message: "Unable to send data to " + host + ":" + strconv.Itoa(port)}
		myError.ERROR()
	}
}

func (c CRH) I_PreTerP(msg *messages.SAMessage) {

	decoder := json.NewDecoder(conn)
	err = decoder.Decode(&msg)

	if err != nil {
		fmt.Println(err)
		myError := errors.MyError{Source: "CRH", Message: "Problem in decoding at Port " + strconv.Itoa(portTmp)}
		myError.ERROR()
	}
	conn.Close()
}
