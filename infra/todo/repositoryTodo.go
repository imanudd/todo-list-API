package todo

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type TodoRepositoryImpl struct {
	db *gorm.DB
}

func NewTodoRepositoryImpl(connectionStr string) TodoRepository {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,
			Colorful:                  true,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      false,
			LogLevel:                  logger.Silent,
		},
	)
	db, err := gorm.Open(mysql.Open(connectionStr), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
	return &TodoRepositoryImpl{
		db: db,
	}
}
func (repo *TodoRepositoryImpl) GetAllTodo(c echo.Context, req Todo) (res Response, err error) {
	todos := []Todo{}
	result := repo.db.Where("activity_group_id = ?", req.Activity_group_id).Find(&todos)
	if result.Error != nil {
		panic(err)
	}
	fmt.Println(req.Activity_group_id)
	res.Data = todos
	return res, nil
}
func (repo *TodoRepositoryImpl) GetTodoById(c echo.Context, id int) (res Response, err error) {
	var todo Todo
	result := repo.db.Where("todo_id = ?", id).Find(&todo)
	if result.Error != nil {
		return res, err
	}

	if todo.Todo_id == 0 {
		return res, err
	} else {
		res.Data = todo
		return res, nil
	}

}

func (repo *TodoRepositoryImpl) CreateTodo(c echo.Context, req Todo) (res Response, err error) {
	// var id int64
	todo := Todo{
		Activity_group_id: req.Activity_group_id,
		Title:             req.Title,
		Is_active:         req.Is_active,
		Priority:          "very-high",
		Created_at:        time.Now().Local(),
		Updated_at:        time.Now().Local(),
	}
	// sqlStatement := "INSERT INTO todos (activity_group_id,title,priority,is_active,created_at,updated_at) values (?,?,?,?,?,?);"
	result := repo.db.Debug().Omit("todo_id").Create(&todo)
	if result.Error != nil {
		panic(err)
	}
	res.Data = todo
	return res, nil
}

func (repo *TodoRepositoryImpl) DeleteTodo(c echo.Context, req Todo) (res Response, err error) {
	var todo Todo
	checkId, _ := repo.GetTodoById(c, req.Todo_id)
	if checkId.Data == nil {
		res.Status = "Not Found"
	}
	result := repo.db.Where("todo_id = ? && title = ?", req.Todo_id, req.Title).Delete(&todo)
	if result.Error != nil {
		panic(err)
	}
	return res, nil
}

func (repo *TodoRepositoryImpl) UpdateTodo(c echo.Context, req Todo) (res Response, err error) {
	todo := Todo{
		Title:     req.Title,
		Priority:  req.Priority,
		Is_active: req.Is_active,
	}
	result := repo.db.Model(&todo).
		Where("todo_id = ?", req.Todo_id).
		Select("title", "priority", "is_active").
		Updates(todo)
	if result.Error != nil {
		return res, err
	}
	getData, _ := repo.GetTodoById(c, req.Todo_id)
	res.Data = getData.Data
	return res, nil

}
