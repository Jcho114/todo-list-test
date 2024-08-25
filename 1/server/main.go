package main

import (
	"database/sql"
	"log"
	"net/http"

	common "github.com/Jcho114/todo-list-test/1/server/common"
	tasks "github.com/Jcho114/todo-list-test/1/server/modules/tasks"
	"github.com/Jcho114/todo-list-test/1/server/modules/templates"

	_ "github.com/mattn/go-sqlite3"
)

func startDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	common.FailOnError(err, "Failed to start sqlite database")

	err = db.Ping()
	common.FailOnError(err, "Failed to connect to database")

	return db
}

func main() {
	db := startDatabase()
	defer db.Close()

	taskRouter := tasks.CreateTaskRouter(db)

	server := http.NewServeMux()

	server.Handle("/tasks/", http.StripPrefix("/tasks", taskRouter.Router))

	server.HandleFunc("/", templates.HomePageHandler)
	server.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir("./modules/templates/public"))))

	log.Print("Server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", server))
}
