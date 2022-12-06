package linting

import "fmt"

const (
	LEVEL_ERROR   = "error"
	LEVEL_WARNING = "warning"
	LEVEL_INFO    = "info"
)

type Issue struct {
	Level   string
	Message string
	Rule    string
}

func (i *Issue) Error() string {
	return fmt.Sprintf("[%s] %s", i.Level, i.Message)
}

func (i *Issue) String() string {
	return i.Error()
}

func (i *Issue) IsError() bool {
	return i.Level == LEVEL_ERROR
}
