package main

import (
	"fmt"
	"midarch-go-v11/src/executionenvironment/executionenvironment"
)

func main() {

	// start configuration
	executionenvironment.ExecutionEnvironment{}.Deploy("SenderReceiver.conf")

	fmt.Scanln()
}

