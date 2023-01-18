package todo

import (
	"fmt"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TodoServiceImpl struct {
	repo TodoRepository
}

func NewTodoServiceImpl(repo TodoRepository) TodoService {
	return &TodoServiceImpl{
		repo: repo,
	}
}

func (svc *TodoServiceImpl) GetAllTodo(c echo.Context, req Todo) (res Response, err error) {
	res, err = svc.repo.GetAllTodo(c, req)
	if err != nil {
		panic("error parsing data")
	}
	res.Status = "Succsess"
	res.Message = "Succsess"
	return res, nil
}

func (svc *TodoServiceImpl) GetTodoById(c echo.Context, id int) (res Response, err error) {

	res, err = svc.repo.GetTodoById(c, id)
	if err != nil {
		panic(err)
	}

	if res.Data == nil {
		return res, err
	}

	res.Status = "Succsess"
	res.Message = "Succsess"
	return res, nil

}
func (svc *TodoServiceImpl) CreateTodo(c echo.Context, req Todo) (res Response, err error) {
	res, err = svc.repo.CreateTodo(c, req)
	if err != nil {
		fmt.Println(err)
	}
	res.Status = "Succsess"
	res.Message = "Succsess"
	return res, nil
}
func (svc *TodoServiceImpl) UpdateTodo(c echo.Context, req Todo) (res Response, err error) {
	res, err = svc.repo.UpdateTodo(c, req)
	if err != nil {
		panic(err)
	}
	if res.Data == nil {
		return res, nil
	}
	res.Status = "Succsess"
	res.Message = "Succsess"
	return res, nil

}

func (svc *TodoServiceImpl) DeleteTodo(c echo.Context, req Todo) (res Response, err error) {
	res, err = svc.repo.DeleteTodo(c, req)
	if err != nil {
		panic(err)
	}
	if res.Status == "Not Found" {
		conv := strconv.Itoa(req.Todo_id)
		res.Message = "Todo with ID " + conv + " Not Found"
		return res, nil
	}
	res.Status = "Succsess"
	res.Message = "Succsess"
	return res, nil

}
