package main

import "time"

//Todo is a struct to store name, completed and due
type Todo struct {
	Id        int `json:"id"`
	Name      string
	Completed bool
	Due       time.Time
}

//Todos is a Todo slice
type Todos []Todo
