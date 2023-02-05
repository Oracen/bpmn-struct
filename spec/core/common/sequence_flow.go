package common

import (
	"github.com/Oracen/bpmn-struct/shared"
	"github.com/Oracen/bpmn-struct/validation"
)

type SequenceFlow struct {
	FlowElement
	SourceRef           FlowNode     `xml:"sourceRef" json:"sourceRef"`
	TargetRef           FlowNode     `xml:"targetRef" json:"targetRef"`
	ConditionExpression []Expression `xml:"conditionExpression" json:"conditionExpression"`
	IsImmediate         []bool       `xml:"isImmediate" json:"isImmediate"`
}

func CreateSequenceFlow(id string, source, target FlowNode) SequenceFlow {
	flowElement := CreateFlowElement(id)

	return SequenceFlow{
		FlowElement:         flowElement,
		SourceRef:           source,
		TargetRef:           target,
		ConditionExpression: []Expression{},
		IsImmediate:         []bool{false},
	}
}

func (s SequenceFlow) Validate(name string) []error {
	checks := []error{}

	name = shared.TypeNameString(name, s, s.Id)
	checks = append(checks, s.FlowElement.Validate(name)...)
	checks = append(checks, s.SourceRef.Validate(name)...)
	checks = append(checks, s.TargetRef.Validate(name)...)
	checks = append(checks, validation.ArrCheckItems(name, s.ConditionExpression)...)
	checks = append(
		checks,
		validation.ArrZeroOne(name, "ConditionExpression", s.ConditionExpression),
		validation.ArrZeroOne(name, "IsImmediate", s.IsImmediate),
	)

	return validation.FilterErrors(checks)
}
