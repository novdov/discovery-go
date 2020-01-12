package task

import (
	"errors"
	"fmt"
)

var ErrTaskNotExist = errors.New("task does not exist")

type InMemoryAccess struct {
	tasks map[ID]Task
	nextID int64
}

func NewInMemoryAccess() *InMemoryAccess {
	return &InMemoryAccess{
		tasks: map[ID]Task{},
		nextID: int64(1),
	}
}

func (m *InMemoryAccess) Get(id ID) (Task, error) {
	t, exists := m.tasks[id]
	if !exists {
		return Task{}, ErrTaskNotExist
	}
	return t, nil
}

// Put overrides tasks
func (m *InMemoryAccess) Put(id ID, t Task) error {
	if _, exists := m.tasks[id]; !exists {
		return ErrTaskNotExist
	}
	m.tasks[id] = t
	return nil
}

func (m *InMemoryAccess) Post(t Task) (ID, error) {
	id := ID(fmt.Sprint(m.nextID))
	m.nextID++
	m.tasks[id] = t
	return id, nil
}

func (m *InMemoryAccess) Delete(id ID) error {
	if _, exists := m.tasks[id]; !exists {
		return ErrTaskNotExist
	}
	delete(m.tasks, id)
	return nil
}