package common

type Message struct {
}

func CreateMessage() Message {
	return Message{}
}

func (m Message) Validate(name string) []error {
	return []error{}
}
