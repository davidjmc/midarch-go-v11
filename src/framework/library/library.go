package library

import (
	"apps/calculator/calculatorclientproxy"
	"framework/components/requestor"
	"framework/connectors"
	"framework/components/srh"
	"framework/components/naming/namingclientproxy"
	"framework/components/crh"
	"apps/calculator/calculatorinvoker"
	"apps/fibonacci/fibonacciclientproxy"
	"apps/fibonacci/fibonacciinvoker"
	"framework/components/queueing/queueingclientproxy"
	"framework/components/naming/naminginvoker"
	"framework/components/queueing/queueinginvoker"
	"framework/components/clientX"
	"framework/components/server"
	"apps/senderreceiver/sender"
	"apps/senderreceiver/receiver"
)

type Record struct {
	RBD   string
	PRISM string
	CSP   string
	Go    interface{}
	PetriNet string
	LTS string   // *.dot file name
}

var Repository = map[string]Record{
	"calculatorclientproxy.CalculatorClientProxy": Record{LTS:"TODO",RBD: "TODO", PRISM: "TODO", Go: calculatorclientproxy.CalculatorClientProxy{}, CSP: "B = I_PreInvR -> InvR.e1 -> TerR.e1 -> I_PosTerR -> B"},
	"calculatorinvoker.CalculatorInvoker": Record{RBD: "TODO", PRISM: "TODO", Go: calculatorinvoker.CalculatorInvoker{}, CSP: "B = InvP.e1 -> I_PosInvP -> TerP.e1 -> B"},
	"fibonacciclientproxy.FibonacciClientProxy": Record{RBD: "TODO", PRISM: "TODO", Go: fibonacciclientproxy.FibonacciClientProxy{}, CSP: "B = I_PreInvR_fibonacciproxy -> InvR.e1 -> TerR.e1 -> I_PosTerR_fibonacciproxy -> B"},
	"fibonacciinvoker.FibonacciInvoker": Record{RBD: "TODO", PRISM: "TODO", Go: fibonacciinvoker.FibonacciInvoker{}, CSP: "B = InvP.e1 -> I_PosInvP_fibonacciinvoker -> TerP.e1 -> B"},
	"requestor.Requestor": Record{RBD: "TODO", PRISM: "TODO", Go: requestor.Requestor{}, CSP: "B = InvP.e1 -> I_PosInvP_requestor -> InvR.e2 -> TerR.e2 -> I_PosTerR_requestor -> TerP.e1 -> B"},
	"connectors.RequestReply": Record{RBD: "TODO", PRISM: "TODO", Go: connectors.RequestReply{}, CSP: "B = InvP.e1 -> InvR.e2 -> TerR.e2 -> TerP.e1 -> B"},
	"connectors.NTo1": Record{RBD: "TODO", PRISM: "TODO", Go: connectors.NTo1{}, CSP: "B = InvP.e1 -> InvR.e2 -> TerR.e2 -> TerP.e1 -> B [] InvP.e3 -> InvR.e2 -> TerR.e2 -> TerP.e3 -> B"},
	"connectors.OneWay": Record{RBD: "TODO", PRISM: "TODO", Go: connectors.OneWay{}, CSP: "B = InvP.e1 -> InvR.e2 -> B"},
	"sender.Sender":Record{RBD: "TODO", PRISM: "TODO", Go: sender.Sender{}, CSP: "B = I_PreInvR_sender -> InvR.e1 -> B"},
	"receiver.Receiver":Record{RBD: "TODO", PRISM: "TODO", Go: receiver.Receiver{}, CSP: "B = InvP.e1 -> I_PosInvP_receiver -> B"},
	"naminginvoker.NamingInvoker": Record{RBD: "TODO", PRISM: "TODO", Go: naminginvoker.NamingInvoker{}, CSP: "B = InvP.e1 -> I_PosInvP_naminginvoker -> TerP.e1 -> B"},
	"queueinginvoker.QueueingInvoker": Record{RBD: "TODO", PRISM: "TODO", Go: queueinginvoker.QueueingInvoker{}, CSP: "B = InvP.e1 -> I_PosInvP -> TerP.e1 -> B"},
	"namingclientproxy.NamingClientProxy": Record{RBD: "TODO", PRISM: "TODO", Go: namingclientproxy.NamingClientProxy{}, CSP: "B = I_PreInvR_namingproxy -> InvR.e1 -> TerR.e1 -> I_PosTerR_namingproxy -> B"},
	"queueingclientproxy.QueueingClientProxy": Record{RBD: "TODO", PRISM: "TODO", Go: queueingclientproxy.QueueingClientProxy{}, CSP: "B = I_PreInvR -> InvR.e1 -> TerR.e1 -> I_PosTerR -> B"},
	"srh.SRH": Record{RBD: "TODO", PRISM: "TODO", Go: srh.SRH{}, CSP: "B = I_PreInvR_srh -> InvR.e1 -> TerR.e1 -> I_PosTerR_srh -> B"},
	"crh.CRH": Record{RBD: "TODO", PRISM: "TODO", Go: crh.CRH{}, CSP: "B = InvP.e1 -> I_PosInvP_crh -> I_PreTerP_crh -> TerP.e1 -> B"}}