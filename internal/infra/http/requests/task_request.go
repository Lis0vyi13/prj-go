package requests

import (
	"time"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
)

type TaskRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	Deadline    int64  `json:"deadline" validate:"required"`
}

type UpdateTaskRequest struct {
	Id          uint64            `json:"id" validate:"required"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Deadline    int64             `json:"deadline"`
	TaskStatus  domain.TaskStatus `json:"status"`
}

type TaskFindById struct {
	Id uint64 `json:"id" validate:"required"`
}

type TaskFindByStatus struct {
	Status string `json:"status" validate:"required"`
}

func (r TaskRequest) ToDomainModel() (interface{}, error) {
	return domain.Task{
		Name:        r.Name,
		Description: r.Description,
		Deadline:    time.Unix(r.Deadline, 0),
	}, nil
}

func (r UpdateTaskRequest) ToDomainModel() (interface{}, error) {
	return domain.Task{
		Id:          r.Id,
		Name:        r.Name,
		Description: r.Description,
		Deadline:    time.Unix(r.Deadline, 0),
		Status:      r.TaskStatus,
	}, nil
}
