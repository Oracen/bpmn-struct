package process

type Monitoring struct {
}

func CreateMonitoring() Monitoring {
	return Monitoring{}
}

func (m Monitoring) Validate(name string) []error {
	return []error{}
}
