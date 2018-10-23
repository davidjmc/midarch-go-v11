package main

import (
	"os/exec"
	"shared/parameters"
)

func main() {

	// createDot("conf.csp", "")

	createDot("conf.csp", "requestor")

}

func createDot(cspfile string, process string) {

	parser := parameters.DIR_CSPARSER + "/" + "CSParser.jar"

	inputFile := parameters.DIR_CSPARSER + "/" + cspfile

	exec.Command("/usr/bin/java", parameters.JAR_COMMAND, parser, inputFile, process)
	//out, err := exec.Command("java").Output()
	//if err != nil {
	//	myError := errors.MyError{Source: "CSParser", Message: "Problem in creating .dot file"}
	//	myError.ERROR()
	//}

	//println(out)
}
