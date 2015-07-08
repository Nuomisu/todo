package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func TodoIndex(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("todo come in")

	todor := TodoR{Data: todos}

	if err := json.NewEncoder(rw).Encode(todor); err != nil {
		panic(err)
	}
}

func TodoCreate(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create todo")
	var todo TodoC
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	fmt.Println(body)
	if err := json.Unmarshal(body, &todo); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	fmt.Println(todo.Data.Id)
	fmt.Println(todo.Data.Title)
	fmt.Println(todo.Data.IsCompleted)

	t := RepoCreateTodo(todo.Data)
	tc := TodoToC(t)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(tc); err != nil {
		panic(err)
	}
}

func TodoUpdate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	toID, err := strconv.Atoi(todoId)
	if err != nil {
		panic(err)
	}
	fmt.Println(todoId)
	fmt.Println("update todo")
	var todo TodoC
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	fmt.Println(body)
	if err := json.Unmarshal(body, &todo); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	fmt.Println(todo.Data.Id)
	fmt.Println(todo.Data.Title)
	fmt.Println(todo.Data.IsCompleted)

	t := RepoUpdateTodo(int32(toID), todo.Data)
	tc := TodoToC(t)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(tc); err != nil {
		panic(err)
	}
}

func TodoDelete(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	todoId := vars["todoId"]
	toID, err := strconv.Atoi(todoId)

	fmt.Println(toID)

	if err != nil {
		panic(err)
	}

	RepoDeleteTodo(int32(toID))
}
