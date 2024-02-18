package activity

type Task struct{}

func CreateTask() Task {
	return Task{}
}

func (t Task) Validate(name string) []error {
	return []error{}
}
