package server

import (
	"net/http"
)

type router struct {
	rules map[string]map[string]http.HandlerFunc
}

func NewRouter() *router {
	return &router{
		rules: make(map[string]map[string]http.HandlerFunc),
	}
}

func (rt *router) FindHandler(method, path string) (http.HandlerFunc, bool, bool) {
	_, isThisPathExists := rt.rules[path]
	handlerFc, thisMethodExists := rt.rules[path][method]
	return handlerFc, thisMethodExists, isThisPathExists
}

func (rt *router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handlerFc, thisMethodExists, isThisPathExists := rt.FindHandler(r.Method, r.URL.Path)
	if !isThisPathExists {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if !thisMethodExists {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	handlerFc(w, r)
}
