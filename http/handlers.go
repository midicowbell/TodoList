package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"study/todo"
	"time"

	"github.com/gorilla/mux"
)

// также много дублирования с проверкой ошибок, нужно фиксить
type HTTPHandlers struct {
	todoList *todo.List
}

func NewHTTPHandlers(todolist *todo.List) *HTTPHandlers {
	return &HTTPHandlers{
		todoList: todolist,
	}
}

func CreateErrroDTO(err error) ErrorDTO {
	return ErrorDTO{
		Message: err.Error(),
		Time:    time.Now(),
	}
}

func (h *HTTPHandlers) HandleAddTask(w http.ResponseWriter, r *http.Request) {
	var task TaskDTO
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		errDTO := CreateErrroDTO(err)

		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}
	if err := task.ValidateForCreate(); err != nil {
		errDTO := CreateErrroDTO(err)

		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}

	todoTask := todo.NewTask(task.Title, task.Description)

	if err := h.todoList.AddTask(todoTask); err != nil {
		errDTO := CreateErrroDTO(err)

		if errors.Is(err, todo.ErrTaskAlreadyExist) {
			http.Error(w, errDTO.ToString(), http.StatusConflict)
		} else {
			http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		}
		return
	}
	b, err := json.MarshalIndent(todoTask, "", "    ")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to response HTTP: ", err)
		return
	}

}

func (h *HTTPHandlers) HandleGetTask(w http.ResponseWriter, r *http.Request) {
	title := mux.Vars(r)["title"]

	task, err := h.todoList.GetTask(title)
	if err != nil {
		if errors.Is(err, todo.ErrTaskNotFound) {
			errorDTO := CreateErrroDTO(todo.ErrTaskNotFound)
			http.Error(w, errorDTO.ToString(), http.StatusNotFound)
		} else {
			errorDTO := CreateErrroDTO(todo.ErrTaskNotFound)
			http.Error(w, errorDTO.ToString(), http.StatusNotFound)
		}
		return
	}
	b, err := json.MarshalIndent(task, "", "   ")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to response HTTP: ", err)
		return
	}
}
func (h *HTTPHandlers) HandleGetAllTask(w http.ResponseWriter, r *http.Request) {
	tasks := h.todoList.ListTasks()
	b, err := json.MarshalIndent(tasks, "", "    ")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to HTTP response")
		return
	}
}

func (h *HTTPHandlers) HandleGetAllUncompletedTask(w http.ResponseWriter, r *http.Request) {
	uncompletedTasks := h.todoList.ListUncompletedTask()
	b, err := json.MarshalIndent(uncompletedTasks, "", "    ")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to HTTP response")
		return
	}
}

func (h *HTTPHandlers) HandleCompleteTask(w http.ResponseWriter, r *http.Request) {
	var completeTaskDTO CompleteTaskDTO
	if err := json.NewDecoder(r.Body).Decode(&completeTaskDTO); err != nil {
		errDTO := CreateErrroDTO(err)
		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}
	title := mux.Vars(r)["title"]
	var (
		task todo.Task
		err  error
	)
	if completeTaskDTO.Complete {
		task, err = h.todoList.CompleteTask(title)
	} else {
		task, err = h.todoList.UncompleteTask(title)
	}
	if err != nil {
		if errors.Is(err, todo.ErrTaskNotFound) {
			errDTO := CreateErrroDTO(err)
			http.Error(w, errDTO.ToString(), http.StatusNotFound)
		} else {
			errDTO := CreateErrroDTO(err)
			http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		}
		return
	}
	b, err := json.MarshalIndent(task, "", "    ")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to HTTP response: ", err)
		return
	}
}

func (h *HTTPHandlers) HandleDeleteTask(w http.ResponseWriter, r *http.Request) {
	title := mux.Vars(r)["title"]
	if err := h.todoList.DeleteTask(title); err != nil {
		if errors.Is(err, todo.ErrTaskNotFound) {
			errDTO := CreateErrroDTO(err)
			http.Error(w, errDTO.ToString(), http.StatusNotFound)
		} else {
			errDTO := CreateErrroDTO(err)
			http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		}
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
