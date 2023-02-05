package common

import (
	"github.com/Oracen/bpmn-struct/shared"
	"github.com/Oracen/bpmn-struct/spec/core/foundation"
	"github.com/Oracen/bpmn-struct/spec/core/infrastructure"
	"github.com/Oracen/bpmn-struct/validation"
)

type ItemKind int

var (
	itemKindString = [...]string{"Physical", "Information"}
)

const (
	ITEM_KIND_Physical ItemKind = iota
	ITEM_KIND_Information
)

func (i ItemKind) String() string {
	return gatewayDirectionString[i]
}

func (i ItemKind) EnumIndex() int {
	return int(i)
}

func (i ItemKind) ToEnum(input string) (relp GatewayDirection, err error) {
	for idx, item := range gatewayDirectionString {
		if item == input {
			return GatewayDirection(idx), nil
		}
	}
	err = validation.NewEnumValidationError("GatewayDirection", "String", gatewayDirectionString[:], input)
	return
}

type ItemDefinition struct {
	foundation.RootElement
	ItemKind     ItemKind                `xml:"itemKind" json:"itemKind"`
	StructureRef []any                   `xml:"structureRef" json:"structureRef"`
	Import       []infrastructure.Import `xml:"import" json:"import"`
	IsCollection bool                    `xml:"isCollection" json:"isCollection"`
}

func CreateItemDefinition(id string, itemKind ItemKind) ItemDefinition {
	rootElement := foundation.CreateRootElement(id)
	return ItemDefinition{
		RootElement:  rootElement,
		ItemKind:     itemKind,
		StructureRef: []any{},
		Import:       []infrastructure.Import{},
		IsCollection: false,
	}
}

func (i ItemDefinition) Validate(name string) []error {
	checks := []error{}

	name = shared.TypeNameString(name, i, i.Id)
	checks = append(checks, i.RootElement.Validate(name)...)
	checks = append(checks, validation.ArrCheckItems(name, i.Import)...)
	checks = append(
		checks,
		validation.ArrZeroOne(name, "StructureRef", i.StructureRef),
		validation.ArrZeroOne(name, "Import", i.Import),
	)

	return validation.FilterErrors(checks)
}
