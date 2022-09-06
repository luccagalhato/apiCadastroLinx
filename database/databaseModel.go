package database

import (
	"apiCadastro/config"
	"database/sql"
	"net/http"
	"net/url"
)

// ConectionSqlConfig ...
type ConectionSqlConfig struct {
	linx   *linx
	server *http.Server
}

// SQLStr ...
type SQLStr struct {
	conf *config.SQL
	url  *url.URL
	db   *sql.DB
}

// linx ...
type linx SQLStr
