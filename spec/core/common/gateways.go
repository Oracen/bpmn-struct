package common

import (
	"github.com/Oracen/bpmn-struct/shared"
	"github.com/Oracen/bpmn-struct/validation"
)

type GatewayDirection int

var (
	gatewayDirectionString = [...]string{"Unspecified", "Converging", "Diverging", "Mixed"}
)

const (
	GATEWAY_DIRECTION_None GatewayDirection = iota
	GATEWAY_DIRECTION_Forward
	GATEWAY_DIRECTION_Backward
	GATEWAY_DIRECTION_Both
)

func (g GatewayDirection) String() string {
	return gatewayDirectionString[g]
}

func (g GatewayDirection) EnumIndex() int {
	return int(g)
}

func (g GatewayDirection) ToEnum(input string) (relp GatewayDirection, err error) {
	for idx, item := range gatewayDirectionString {
		if item == input {
			return GatewayDirection(idx), nil
		}
	}
	err = validation.NewEnumValidationError("GatewayDirection", "String", gatewayDirectionString[:], input)
	return
}

type Gateway struct {
	FlowNode
	GatewayDirection GatewayDirection `xml:"GatewayDirection" json:"GatewayDirection"`
}

func CreateGateway(id string) FlowNode {
	flowElement := CreateFlowElement(id)
	return FlowNode{
		FlowElement: flowElement,
	}
}

func (g Gateway) Validate(name string) []error {
	checks := []error{}

	name = shared.TypeNameString(name, g, g.Id)
	checks = append(checks, g.FlowNode.Validate(name)...)
	return validation.FilterErrors(checks)
}
