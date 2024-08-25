package tasks

import (
	"database/sql"
	"net/http"

	common "github.com/Jcho114/todo-list-test/1/server/common"
)

type TaskRouter struct {
	Router     *http.ServeMux
	controller *TaskController
}

func NewTaskRouter(controller *TaskController) (*TaskRouter, error) {
	router := http.NewServeMux()
	return &TaskRouter{router, controller}, nil
}

func (router *TaskRouter) RegisterRoutes() {
	router.Router.HandleFunc("GET /", router.controller.FindAll)
	router.Router.HandleFunc("POST /create", router.controller.CreateTask)
}

func CreateTaskRouter(db *sql.DB) *TaskRouter {
	taskRepository, err := NewTaskRepository(db)
	common.FailOnError(err, "Failed to create task repository")
	err = taskRepository.InitDatabase()
	common.FailOnError(err, "Failed to initialize task repository")

	taskService, err := NewTaskService(taskRepository)
	common.FailOnError(err, "Failed to create task service")

	taskController, err := NewTaskController(taskService)
	common.FailOnError(err, "Failed to create task controller")

	taskRouter, err := NewTaskRouter(taskController)
	common.FailOnError(err, "Failed to create task router")
	taskRouter.RegisterRoutes()

	return taskRouter
}
