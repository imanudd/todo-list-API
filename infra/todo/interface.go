package todo

import "github.com/labstack/echo/v4"

type TodoRepository interface {
	GetAllTodo(c echo.Context, req Todo) (res Response, err error)
	GetTodoById(c echo.Context, id int) (res Response, err error)
	CreateTodo(c echo.Context, req Todo) (res Response, err error)
	UpdateTodo(c echo.Context, req Todo) (res Response, err error)
	DeleteTodo(c echo.Context, req Todo) (res Response, err error)
}

type TodoService interface {
	GetAllTodo(c echo.Context, req Todo) (res Response, err error)
	GetTodoById(c echo.Context, id int) (res Response, err error)
	CreateTodo(c echo.Context, req Todo) (res Response, err error)
	UpdateTodo(c echo.Context, req Todo) (res Response, err error)
	DeleteTodo(c echo.Context, req Todo) (res Response, err error)
}
