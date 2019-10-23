package urls

import (
	"database/sql"
	"net/http"

	get_task_detail_ssr "../controllers/ssr/get_task_detail"
	get_tasks_ssr "../controllers/ssr/get_tasks"

	get_task_detail "../controllers/rest/get_task_detail"
	get_tasks "../controllers/rest/get_tasks"

	create_task "../controllers/rest/create_task"
)

func SetUrlHandlers(Db *sql.DB) {
	getTasksHandlerSSR := get_tasks_ssr.GetTasks(Db)
	http.HandleFunc("/", getTasksHandlerSSR.Handler) // Works!

	getTaskDetailHandlerSSR := get_task_detail_ssr.GetTaskDetail(Db, func(url string, request *http.Request) []string {
		id := request.URL.Path[len("/tasks/"):]
		params := []string{id}
		return params
	})
	http.HandleFunc("/tasks/", getTaskDetailHandlerSSR.Handler) // Works!

	createTaskHandler := create_task.CreateTask(Db)
	http.HandleFunc("/api/new/", createTaskHandler.Handler)

	getTaskHandler := get_tasks.GetTasks(Db)
	http.HandleFunc("/api/", getTaskHandler.Handler)

	getTaskDetailHandler := get_task_detail.GetTaskDetail(Db, func(url string, request *http.Request) []string {
		id := request.URL.Path[len("/api/tasks/"):]
		params := []string{id}
		return params
	})
	http.HandleFunc("/api/tasks/", getTaskDetailHandler.Handler)
}
