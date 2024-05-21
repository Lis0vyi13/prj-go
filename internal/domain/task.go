package domain

import "time"

type Task struct {
	Id          uint64
	UserId      uint64
	Name        string
	Description string
	Deadline    time.Time
	Status      TaskStatus
	CreatedDate time.Time
	UpdatedDate time.Time
	DeletedDate *time.Time
}

type TaskStatus string

const (
	New        TaskStatus = "NEW"
	InProgress TaskStatus = "IN_PROGRESS"
	Done       TaskStatus = "DONE"
)

func (status TaskStatus) IsStatusValid() bool {
	switch status {
	case New, InProgress, Done:
		return true
	}
	return false
}
