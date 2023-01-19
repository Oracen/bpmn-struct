package common

type FormalExpression struct {
}

func CreateFormalExpression() FormalExpression {
	return FormalExpression{}
}

func (m FormalExpression) Validate(name string) []error {
	return []error{}
}
