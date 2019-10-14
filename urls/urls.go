package urls

import (
	"database/sql"
	"net/http"

	get_tasks_ssr "../controllers/ssr/get_tasks"
)

func SetUrlHandlers(Db *sql.DB) {
	getTasksHandlerSSR := get_tasks_ssr.GetTasks(Db)
	http.HandleFunc("/", getTasksHandlerSSR.Handler)
}
