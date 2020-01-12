package task

// ID is a data type to identify the task.
type ID string

type DataAccess interface {
	Get(id ID) (Task, error)
	Put(id ID, t Task) error
	Post(t Task) (ID, error)
	Delete(id ID) error
}
