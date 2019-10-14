package controllers

import (
	"database/sql"
	"net/http"
)

type UrlHandler func(writer http.ResponseWriter, request *http.Request)

type BindUrlHandler struct {
	Handler UrlHandler
	Db      *sql.DB
}
