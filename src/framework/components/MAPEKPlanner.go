package components

import (
	"framework/messages"
	"shared/shared"
)

type MAPEKPlanner struct {}

func (MAPEKPlanner) I_PosInvP(msg *messages.SAMessage, r *bool) {
	*msg = messages.SAMessage{Payload:shared.AdaptationPlan{Plan:"Adaptation Plan"}}
}