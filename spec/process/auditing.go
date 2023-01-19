package process

type Auditing struct {
}

func CreateAuditing() Auditing {
	return Auditing{}
}

func (m Auditing) Validate(name string) []error {
	return []error{}
}
