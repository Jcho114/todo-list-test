package tasks

type TaskService struct {
	repository *TaskRepository
}

func NewTaskService(repository *TaskRepository) (*TaskService, error) {
	return &TaskService{repository}, nil
}

func (service *TaskService) FindAll() (string, error) {
	return service.repository.GetAllTasks()
}

func (service *TaskService) CreateTask(name string) (int64, error) {
	return service.repository.CreateTask(name)
}
