package message_flow

type InteractionNode struct {
}

func CreateInteractionNode(id string) InteractionNode {
	return InteractionNode{}
}

func (i InteractionNode) Validate() []error {
	checks := []error{}
	return checks
}
