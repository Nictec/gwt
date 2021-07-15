package handler

import (
	"github.com/gorilla/sessions"
	"net/http"
)

type Context struct {
	W http.ResponseWriter
	Request *http.Request
	Session *sessions.CookieStore
	Config map[string]string
	Debug bool
}

type Handler func(ctx Context)