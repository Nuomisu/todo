package main

var currentId int32

var todos Todos

func init() {
	RepoCreateTodo(Todo{Id: 1, Title: "Write presentation"})
	RepoCreateTodo(Todo{Id: 2, Title: "Write presentation"})
}

func RepoCreateTodo(t Todo) Todo {
	currentId += 1
	t.Id = currentId
	todos = append(todos, t)
	return t
}

func RepoUpdateTodo(id int32, to Todo) Todo {
	for i, t := range todos {
		if t.Id == id {
			todos[i].Title = to.Title
			todos[i].IsCompleted = to.IsCompleted
			return todos[i]
		}
	}
	// return empty Todo if not found
	return Todo{}
}

func RepoDeleteTodo(id int32) {
	for i, t := range todos {
		if t.Id == id {
			todos = append(todos[:i], todos[i+1:]...)
		}
	}
}

func TodoToC(t Todo) TodoC {
	return TodoC{Data: t}
}
