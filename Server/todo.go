package main

type Todo struct {
	Id          int32  `json:"id"`
	Title       string `json:"title"`
	IsCompleted bool   `json:"iscompleted"`
}

type Todos []Todo

type TodoR struct {
	Data Todos `json:"todos"`
}

type TodoC struct {
	Data Todo `json:"todo"`
}
