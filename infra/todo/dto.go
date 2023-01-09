package todo

import (
	"time"
)

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Todo struct {
	Todo_id           int       `json:"todo_id"`
	Activity_group_id int       `json:"activity_group_id"`
	Title             string    `json:"title"`
	Priority          string    `json:"priority"`
	Is_active         bool      `json:"is_active"`
	Created_at        time.Time `json:"createdAt"`
	Updated_at        time.Time `json:"updatedAt"`
}
