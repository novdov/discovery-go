package taskman

import (
	"github.com/novdov/discovery-go/task"
)

// ID is a data type to identify the task.
type ID string

type DataAccess interface {
	Get(id ID) (task.Task, error)
	Put(id ID, t task.Task) error
	Post(t task.Task) (ID, error)
	Delete(id ID) error
}
