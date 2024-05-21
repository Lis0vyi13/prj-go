package resources

import (
	"time"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
)

type TaskDto struct {
	Id          uint64            `json:"id"`
	UserId      uint64            `json:"userId"`
	Name        string            `json:"name"`
	Description string            `json:"description,omitempty"`
	Deadline    time.Time         `json:"deadline"`
	Status      domain.TaskStatus `json:"status"`
}

func (d TaskDto) DomainToDto(t domain.Task) TaskDto {
	return TaskDto{
		Id:          t.Id,
		UserId:      t.UserId,
		Name:        t.Name,
		Description: t.Description,
		Deadline:    t.Deadline,
		Status:      t.Status,
	}
}
