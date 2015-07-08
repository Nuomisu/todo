package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name       string
	Method     string
	Pattern    string
	HandleFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(route.HandleFunc)
	}
	return router
}

var routes = Routes{
	Route{
		"TodoIndex",
		"Get",
		"/todos",
		TodoIndex,
	},
	Route{
		"TodoCreate",
		"Post",
		"/todos",
		TodoCreate,
	},
	Route{
		"TodoUpdate",
		"Put",
		"/todos/{todoId}",
		TodoUpdate,
	},
	Route{
		"TodoDelete",
		"Delete",
		"/todos/{todoId}",
		TodoDelete,
	},
}
