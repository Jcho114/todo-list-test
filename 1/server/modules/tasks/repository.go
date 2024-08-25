package tasks

import (
	"database/sql"
	"fmt"
	"log"

	common "github.com/Jcho114/todo-list-test/1/server/common"

	_ "github.com/mattn/go-sqlite3"
)

type TaskRepository struct {
	*sql.DB
}

func NewTaskRepository(db *sql.DB) (*TaskRepository, error) {
	return &TaskRepository{db}, nil
}

func (repository *TaskRepository) InitDatabase() error {
	log.Print("Initializing task database...\n")

	log.Print("Creating tasks table if it does not exist\n")
	_, err := repository.Exec("CREATE TABLE IF NOT EXISTS TASKS ( id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, name VARCHAR(255) NOT NULL )")

	if err != nil {
		return err
	}

	log.Printf("Initialized task database!\n")

	return nil
}

func (repository *TaskRepository) GetAllTasks() (string, error) {
	rows, err := repository.Query("SELECT * FROM TASKS")

	if err != nil {
		return "", err
	}

	defer rows.Close()

	return common.RowsToJson(rows), nil
}

func (repository *TaskRepository) CreateTask(name string) (int64, error) {
	res, err := repository.Exec(fmt.Sprintf("INSERT INTO TASKS (name) VALUES('%s')", name))

	if err != nil {
		return -1, err
	}

	id, err := res.LastInsertId()

	if err != nil {
		return -1, err
	}

	return id, nil
}
