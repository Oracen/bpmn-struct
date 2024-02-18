package event

type Event struct{}

func CreateEvent(id string) Event {
	return Event{}
}

func (e Event) Validate(name string) []error {
	return []error{}
}
