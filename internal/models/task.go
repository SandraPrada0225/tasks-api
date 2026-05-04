package models

// define como luce una tarea
type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}
