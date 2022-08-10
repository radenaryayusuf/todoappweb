package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/radenaryayusuf/todoappweb/config"
	"github.com/radenaryayusuf/todoappweb/models"
)

var (
	id       int
	content  string
	isdone   bool
	assignee string
	deadline string
	created  time.Time
	view     = template.Must(template.ParseFiles("./web/index.html"))
	database = config.Database()
)

func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	stmt := `SELECT id, content, isdone, assignee, cast(deadline as varchar(30)) as deadline, created FROM tasks`

	rows, err := database.Query(stmt)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	var tasks []models.Task

	for rows.Next() {

		err = rows.Scan(&id, &content, &isdone, &assignee, &deadline, &created)
		if err != nil {
			fmt.Println(err)
		}

		// fmt.Println("my deadline date: ", timeT)
		task := models.Task{
			ID:       id,
			Content:  content,
			IsDone:   isdone,
			Assignee: assignee,
			Deadline: deadline,
			Created:  created,
		}

		tasks = append(tasks, task)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err)
	}

	data := models.View{
		Tasks: tasks,
	}

	_ = view.Execute(w, data)
}

func Insert(w http.ResponseWriter, r *http.Request) {

	content := r.FormValue("content")
	assignee := r.FormValue("assignee")
	deadline := r.FormValue("deadline")

	stmt := `INSERT INTO tasks(content, assignee, deadline) VALUES($1, $2, $3)`

	_, err := database.Exec(stmt, content, assignee, deadline)

	if err != nil {
		fmt.Println(err)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	stmt := `DELETE FROM tasks WHERE id = $1`

	_, err := database.Exec(stmt, id)

	if err != nil {
		fmt.Println(err)
	}

	http.Redirect(w, r, "/", 301)
}

func Complete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	fmt.Println(id)
	stmt := `UPDATE tasks SET isdone = true WHERE id = $1`

	_, err := database.Exec(stmt, id)

	// fmt.Println(stmt)

	if err != nil {
		fmt.Println(err)
	}

	http.Redirect(w, r, "/", 301)
}
