package todo

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func RegisterRoute(e *echo.Echo, svc *API) {
	e.GET("/todo-items:activity_group_id", svc.GetAllTodo)
	e.GET("/todo-items/:id", svc.GetTodoByID)
	e.POST("/todo-items", svc.CreateTodo)
	e.PATCH("/todo-items/:id", svc.UpdateTodo)
	e.DELETE("/todo-items/:id", svc.DeleteTodo)
}

type API struct {
	svc TodoService
}

func NewTodoAPIImpl(svc TodoService) *API {
	return &API{
		svc: svc,
	}
}

func (api *API) GetAllTodo(c echo.Context) error {
	var req = new(Todo)
	id, _ := strconv.Atoi(c.QueryParam("activity_group_id"))
	req.Activity_group_id = id
	fmt.Println("Ini di API", req.Activity_group_id)
	res, err := api.svc.GetAllTodo(c, *req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, res)
}

func (api *API) GetTodoByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	res, err := api.svc.GetTodoById(c, id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": err.Error(),
		})
	}
	if res.Data == nil {
		conv := strconv.Itoa(id)
		return c.JSON(http.StatusNotFound, map[string]string{
			"Status":  "Not Found",
			"Message": "Todo with ID " + conv + " Not Found",
		})
	} else {
		return c.JSON(http.StatusOK, res)
	}
}

func (api *API) CreateTodo(c echo.Context) error {
	var json = new(Todo)
	if err := c.Bind(json); err != nil {
		return err
	}
	if json.Title == "" && json.Activity_group_id == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Status":  "Bad Request",
			"Message": "Title Cannot Be Null",
		})
	}
	res, err := api.svc.CreateTodo(c, *json)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, res)
}

func (api *API) UpdateTodo(c echo.Context) error {
	var json = new(Todo)
	var res Response
	json.Todo_id, _ = strconv.Atoi(c.Param("id"))

	if err := c.Bind(json); err != nil {
		return c.JSON(http.StatusNotFound, res)
	}
	if json.Title == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Status":  "Bad Request",
			"Message": "Title Cannot Be Null",
		})
	}
	res, err := api.svc.UpdateTodo(c, *json)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}
	if res.Data == nil {
		conv := strconv.Itoa(json.Todo_id)
		return c.JSON(http.StatusNotFound, map[string]string{
			"Status":  "Not Found",
			"Message": "Todo with ID " + conv + " Not Found",
		})
	}
	return c.JSON(http.StatusCreated, res)
}

func (api *API) DeleteTodo(c echo.Context) error {
	var todo = new(Todo)
	todo.Todo_id, _ = strconv.Atoi(c.Param("id"))
	if err := c.Bind(todo); err != nil {
		return err
	}
	res, err := api.svc.DeleteTodo(c, *todo)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	if res.Status == "Not Found" {
		conv := strconv.Itoa(todo.Todo_id)
		return c.JSON(http.StatusNotFound, map[string]string{
			"Status":  "Not Found",
			"Message": "Todo with ID " + conv + " Not Found",
		})
	}
	return c.JSON(http.StatusOK, res)
}
