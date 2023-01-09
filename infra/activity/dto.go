package activity

import (
	"time"
)

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Activity struct {
	Activity_id int       `grom:"column:activity_id; primary_key" json:"activityId"`
	Title       string    `json:"title"`
	Email       string    `json:"email"`
	Created_at  time.Time `json:"createdAt"`
	Updated_at  time.Time `json:"updatedAt"`
}

type Todo struct {
	Todo_id     int       `json:"todo_id"`
	Activity_id int       `json:"activity_id"`
	Title       string    `json:"title_id"`
	Priority    string    `json:"priority"`
	Is_active   bool      `json:"is_active"`
	Created_at  time.Time `json:"createdAt"`
	Updated_at  time.Time `json:"updatedAt"`
}
