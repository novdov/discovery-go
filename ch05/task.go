package ch05

import "time"

type status int

const (
	UNKNOWN status = iota
	TODO
	DONE
)

type Deadline struct {
	time.Time
}

func NewDeadline(t time.Time) *Deadline {
	return &Deadline{t}
}

type Task struct {
	Title    string
	Status   status
	Deadline *Deadline
}

type MyStruct struct {
	Title    string `json:"title"`
	Internal string `json:"-"`
	Value    int64  `json:",omitempty"`
	ID       int64  `json:",string"`
}
