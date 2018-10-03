package task

type Task struct {
	Name          string
	MessageID     string
	Description   string
	CreatedByName string
	CreatedByURI  string
	Step          []Step
}

func (t *Task) AddStep() *Task {
	return t
}
