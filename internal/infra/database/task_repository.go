package database

import (
	"time"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/upper/db/v4"
)

const TasksTableName = "tasks"

type task struct {
	Id          uint64            `db:"id,omitempty"`
	UserId      uint64            `db:"user_id"`
	Name        string            `db:"name"`
	Description string            `db:"description"`
	Deadline    time.Time         `db:"deadline"`
	Status      domain.TaskStatus `db:"status"`
	CreatedDate time.Time         `db:"created_date"`
	UpdatedDate time.Time         `db:"updated_date"`
	DeletedDate *time.Time        `db:"deleted_date"`
}

type TaskRepository interface {
	Save(t domain.Task) (domain.Task, error)
	Edit(t domain.Task) (domain.Task, error)
	FindById(id uint64, userId uint64) (domain.Task, error)
	FindByStatus(userId uint64, status string) ([]domain.Task, error)
	Delete(id uint64, userId uint64) error
}

type taskRepository struct {
	coll db.Collection
	sess db.Session
}

func NewTaskRepository(session db.Session) TaskRepository {
	return taskRepository{
		coll: session.Collection(TasksTableName),
		sess: session,
	}
}

func (r taskRepository) Save(t domain.Task) (domain.Task, error) {
	tsk := r.mapDomainToModel(t)
	tsk.CreatedDate = time.Now()
	tsk.UpdatedDate = time.Now()
	err := r.coll.InsertReturning(&tsk)
	if err != nil {
		return domain.Task{}, err
	}
	t = r.mapModelToDomain(tsk)
	return t, nil
}

func (r taskRepository) Edit(t domain.Task) (domain.Task, error) {

	var tsk task
	result := r.coll.Find(db.Cond{"id": t.Id, "user_id": t.UserId})
	err := result.One(&tsk)
	if err != nil {
		return domain.Task{}, err
	}

	if t.Name != "" {
		tsk.Name = t.Name
	}
	if t.Description != "" {
		tsk.Description = t.Description
	}
	if !t.Deadline.IsZero() {
		tsk.Deadline = t.Deadline
	}
	if t.Status != "" && t.Status.IsStatusValid() {
		tsk.Status = t.Status
	}
	tsk.UpdatedDate = time.Now()

	err = result.Update(&tsk)
	if err != nil {
		return domain.Task{}, err
	}

	err = result.One(&tsk)
	if err != nil {
		return domain.Task{}, err
	}
	return r.mapModelToDomain(tsk), nil

}

func (r taskRepository) FindById(id uint64, userId uint64) (domain.Task, error) {
	var tsk task
	err := r.coll.Find(db.Cond{"id": id, "user_id": userId}).One(&tsk)
	if err != nil {
		if err == db.ErrNoMoreRows {
			return domain.Task{}, nil
		}
		return domain.Task{}, err
	}
	return r.mapModelToDomain(tsk), nil
}


func (r taskRepository) FindByStatus(userId uint64, status string) ([]domain.Task, error) {
    var tasks []task
    err := r.coll.Find(db.Cond{"user_id": userId, "status": status}).All(&tasks)
    if err != nil {
        return nil, err
    }
    
    domainTasks := make([]domain.Task, len(tasks))
    for i, tsk := range tasks {
        domainTasks[i] = r.mapModelToDomain(tsk)
    }
    
    return domainTasks, nil
}

func (r taskRepository) Delete(id uint64, userId uint64) error {
	err := r.coll.Find(db.Cond{"id": id, "user_id": userId}).Delete()
	if err != nil {
		return err
	}
	return nil
}

func (r taskRepository) mapDomainToModel(t domain.Task) task {
	return task{
		Id:          t.Id,
		UserId:      t.UserId,
		Name:        t.Name,
		Description: t.Description,
		Deadline:    t.Deadline,
		Status:      t.Status,
		CreatedDate: t.CreatedDate,
		UpdatedDate: t.UpdatedDate,
		DeletedDate: t.DeletedDate,
	}
}

func (r taskRepository) mapModelToDomain(t task) domain.Task {
	return domain.Task{
		Id:          t.Id,
		UserId:      t.UserId,
		Name:        t.Name,
		Description: t.Description,
		Deadline:    t.Deadline,
		Status:      t.Status,
		CreatedDate: t.CreatedDate,
		UpdatedDate: t.UpdatedDate,
		DeletedDate: t.DeletedDate,
	}
}
