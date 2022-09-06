package server

import (
	"apiCadastro/config"
	"database/sql"
	"net/http"
	"net/url"
)

// Handler ...
type Handler struct {
	Method string
	Fn     http.HandlerFunc
}

// SQLStr ...
type SQLStr struct {
	conf *config.SQL
	url  *url.URL
	db   *sql.DB
}
