package tasks

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type TaskController struct {
	service *TaskService
}

func NewTaskController(service *TaskService) (*TaskController, error) {
	return &TaskController{service}, nil
}

func (controller *TaskController) FindAll(w http.ResponseWriter, r *http.Request) {
	log.Print("Inside the FindAll API\n")
	tasks, err := controller.service.FindAll()

	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, tasks)
}

type TaskCreateRequest struct {
	Name string
}

func (controller *TaskController) CreateTask(w http.ResponseWriter, r *http.Request) {
	log.Print("Inside the CreateTask API\n")
	var request TaskCreateRequest
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		w.WriteHeader(400)
		return
	}

	newId, err := controller.service.CreateTask(request.Name)

	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{ "message": "Successfully created task", "newId": "%d" }`, newId)
}
