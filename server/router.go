package server

import (
	"net/http"
	"regexp"
)

// The router represents the router object.
// It has a rules field, which is a nested map used to
// store the routing rules. The outer map uses HTTP methods
// (e.g., GET, POST, etc.) as keys, and the inner map
// uses URL paths as keys, with http.HandlerFunc
// as the corresponding handler function.
type router struct {
	rules map[string]map[string]http.HandlerFunc
}

// The NewRouter creates a new instance of the router
// struct. It initializes the rules field with an
// empty map using make. It returns a pointer to
// the created router.
func NewRouter() *router {
	return &router{
		rules: make(map[string]map[string]http.HandlerFunc),
	}
}

// The FindHandler method of the router is used to
// find the appropriate handler function for a
// given HTTP method and URL path. It takes the method
// and path as parameters and returns the matching
// http.HandlerFunc, along with boolean values
// indicating whether the method and path exist in
// the router's rules. It iterates through the rules
// for the specified method, using regular expressions
// to match the path against the stored routes.
// If a match is found, it returns the corresponding
// handler function and the boolean values.
func (rt *router) FindHandler(method, path string) (http.HandlerFunc, bool, bool) {
	_, methodExists := rt.rules[method]
	for route, handlerLogic := range rt.rules[method] {
		pathExists, _ := regexp.MatchString("^"+route+"$", path)
		if pathExists {
			return handlerLogic, methodExists, pathExists
		}
	}
	return nil, methodExists, false
}

// The ServeHTTP method of the router is the implementation
// of the http.Handler interface. It receives an incoming
// HTTP request and determines the appropriate handler
// function using FindHandler. If the method or path does
// not exist, it returns the corresponding HTTP status code.
// Otherwise, it calls the obtained handlerFunc,
// passing the response writer and request objects.
func (rt *router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handlerLogic, methodExists, pathExists := rt.FindHandler(r.Method, r.URL.Path)
	if !methodExists {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	if !pathExists {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	handlerLogic(w, r)
}
