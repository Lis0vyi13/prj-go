package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/app"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/requests"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/resources"
)

type TaskController struct {
	taskService app.TaskService
}

func NewTaskController(ts app.TaskService) TaskController {
	return TaskController{
		taskService: ts,
	}
}

func (c TaskController) Save() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		task, err := requests.Bind(r, requests.TaskRequest{}, domain.Task{})
		if err != nil {
			log.Printf("TaskController: %s", err)
			BadRequest(w, err)
			return
		}

		user := r.Context().Value(UserKey).(domain.User)
		task.UserId = user.Id
		task.Status = domain.New
		task, err = c.taskService.Save(task)
		if err != nil {
			log.Printf("TaskController: %s", err)
			InternalServerError(w, err)
			return
		}

		var tDto resources.TaskDto
		tDto = tDto.DomainToDto(task)
		Created(w, tDto)
	}
}

func (c TaskController) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		task, err := requests.Bind(r, requests.UpdateTaskRequest{}, domain.Task{})
		if err != nil {
			log.Printf("TaskController: %s", err)
			BadRequest(w, err)
			return
		}

		user := r.Context().Value(UserKey).(domain.User)
		task.UserId = user.Id
		task, err = c.taskService.Update(task)
		if err != nil {
			log.Printf("TaskController: %s", err)
			InternalServerError(w, err)
			return
		}

		var tDto resources.TaskDto
		tDto = tDto.DomainToDto(task)
		Created(w, tDto)
	}
}

func (c TaskController) FindById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var TaskFindById requests.TaskFindById
		if err := json.NewDecoder(r.Body).Decode(&TaskFindById); err != nil {
			log.Printf("TaskController: %s", err)
			BadRequest(w, err)
			return
		}

		id := TaskFindById.Id
		user := r.Context().Value(UserKey).(domain.User)
		task, err := c.taskService.FindById(id, user.Id)
		if err != nil {
			log.Printf("TaskController: %s", err)
			InternalServerError(w, err)
			return
		}
		var tDto resources.TaskDto
		tDto = tDto.DomainToDto(task)
		Success(w, tDto)

	}

}

func (c TaskController) FindByStatus() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var TaskFindByStatus requests.TaskFindByStatus
		
		if err := json.NewDecoder(r.Body).Decode(&TaskFindByStatus); err != nil {
			log.Printf("TaskController: %s", err)
			BadRequest(w, err)
			return
		}

		user := r.Context().Value(UserKey).(domain.User)
		tasks, err := c.taskService.FindByStatus(user.Id, TaskFindByStatus.Status)
		if err != nil {
			log.Printf("TaskController: %s", err)
			InternalServerError(w, err)
			return
		}

		var taskDTOs []resources.TaskDto 
		for _, task := range tasks {
			taskDTOs = append(taskDTOs, resources.TaskDto{}.DomainToDto(task)) 
		}

		Success(w, taskDTOs)
	}
}

func (c TaskController) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var TaskFindById requests.TaskFindById
		if err := json.NewDecoder(r.Body).Decode(&TaskFindById); err != nil {
			log.Printf("TaskController: %s", err)
			BadRequest(w, err)
			return
		}

		id := TaskFindById.Id
		user := r.Context().Value(UserKey).(domain.User)
		err := c.taskService.Delete(id, user.Id)
		if err != nil {
			log.Printf("TaskController: %s", err)
			InternalServerError(w, err)
			return
		}
		Ok(w)

	}

}
