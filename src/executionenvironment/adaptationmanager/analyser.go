package adaptationmanager

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"framework/configuration/configuration"
	"shared/parameters"
	"shared/shared"
	"net/http"
	"strconv"
	"strings"
	"unicode"
)

type Analyser struct{}

func (Analyser) Exec(conf configuration.Configuration, chanMACorrective chan shared.MonitoredCorrectiveData, chanMAEvolutive chan shared.MonitoredEvolutiveData, chanMAProactive chan shared.MonitoredProactiveData, chanAP chan shared.AnalysisResult) {

	// prepapre channels
	chanCorrective := make(chan shared.AnalysisResult)
	chanEvolutive := make(chan shared.AnalysisResult)
	chanProactive := make(chan shared.AnalysisResult)

	if parameters.IS_CORRECTIVE {
		go correctiveAnalysis(chanMACorrective, chanCorrective)
	}
	if parameters.IS_EVOLUTIVE {
		go evolutiveAnalysis(chanMAEvolutive, chanEvolutive)
	}
	if parameters.IS_PROACTIVE {
		go proactiveAnalysis(chanMAProactive, chanProactive)
	}

	for {
		select {
		case analysisResult := <-chanCorrective:
			chanAP <- analysisResult
		case analysisResult := <-chanEvolutive:
			chanAP <- analysisResult
		case analysisResult := <-chanProactive:
			chanAP <- analysisResult
		}
	}
}

func correctiveAnalysis(chanMa chan shared.MonitoredCorrectiveData, chanCorrective chan shared.AnalysisResult) {

	for {
		monitoredData := <-chanMa
		r := invokePROM(monitoredData)
		if r {
			chanCorrective <- shared.AnalysisResult{Analysis: parameters.NO_CHANGE} // TODO
		}
	}
}

func proactiveAnalysis(chanMa chan shared.MonitoredProactiveData, chanProactive chan shared.AnalysisResult) {
	analysisResult := shared.AnalysisResult{}

	for {
		monitoredData := <-chanMa
		r := invokePRISM(monitoredData)
		if r {
			analysisResult.Result = "TODO"
			analysisResult.Analysis = parameters.PROACTIVE_CHANGE
			chanProactive <- analysisResult
		}
	}
}

func evolutiveAnalysis(chanMa chan shared.MonitoredEvolutiveData, chanEvolutive chan shared.AnalysisResult) {
	analysisResult := shared.AnalysisResult{}

	for {
		listOfNewPlugins := <-chanMa // receive new plugins from Monitor
		analysisResult.Result = listOfNewPlugins
		analysisResult.Analysis = parameters.EVOLUTIVE_CHANGE
		chanEvolutive <- analysisResult
	}
}

func invokePROM(data shared.MonitoredCorrectiveData) bool {
	// TODO
	return false
}

func invokePRISM(data shared.MonitoredProactiveData) bool {

	// request
	req,_:= http.NewRequest("POST", "http://localhost:8080/ProactiveChecker/service/postchecker", bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	// response
	resp,_ := client.Do(req)
	body,_ := ioutil.ReadAll(resp.Body)

	words := strings.Split(string(body), ":")

	atr := strings.TrimFunc(words[0], func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})

	value := strings.TrimFunc(words[1], func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})

	prismResult,_ := strconv.ParseBool(value)

	fmt.Println(atr + " " + value)

	return prismResult
}
