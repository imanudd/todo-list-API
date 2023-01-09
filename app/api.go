package app

import (
	"to-do-list/infra/activity"
	"to-do-list/infra/todo"

	"github.com/labstack/echo/v4"
)

func Init() {
	e := echo.New()
	registerProductAPI(e, "mysql")
	e.Logger.Fatal(e.Start(":3030"))
}

func registerProductAPI(e *echo.Echo, db string) {
	var activityRepository activity.ActivityRepository
	var todoRepository todo.TodoRepository
	switch db {
	case "mysql":
		activityRepository = activity.NewActivityRepositoryImpl("root:root@tcp(127.0.0.1:3306)/todolist?charset=utf8mb4&parseTime=True&loc=Local")
		todoRepository = todo.NewTodoRepositoryImpl("root:root@tcp(127.0.0.1:3306)/todolist?charset=utf8mb4&parseTime=True&loc=Local")
	default:
		panic(`unknown db selections!`)
	}

	activityService := activity.NewActivityServiceImpl(activityRepository)
	activityApi := activity.NewActivityAPIImpl(activityService)
	activity.RegisterRoute(e, activityApi)

	todoService := todo.NewTodoServiceImpl(todoRepository)
	todoApi := todo.NewTodoAPIImpl(todoService)
	todo.RegisterRoute(e, todoApi)
}
