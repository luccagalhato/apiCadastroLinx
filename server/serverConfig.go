package server

import (
	"apiCadastro/config"

	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// NewServer ...
func NewServer(conf config.Config, routes map[string]Handler) *http.Server {
	r := mux.NewRouter()

	for path, handler := range routes {
		r.Path(path).Methods(handler.Method).HandlerFunc(handler.Fn)
	}

	return &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf(":%d", conf.APIPort),
		WriteTimeout: 150 * time.Second,
		ReadTimeout:  150 * time.Second,
	}
}
