package common

type ItemDefinition struct {
}

func CreateItemDefinition() ItemDefinition {
	return ItemDefinition{}
}

func (m ItemDefinition) Validate(name string) []error {
	return []error{}
}
