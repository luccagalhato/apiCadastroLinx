package server

import (
	"net/http"
)

// Handler ...
type Handler struct {
	Method string
	Fn     http.HandlerFunc
}
