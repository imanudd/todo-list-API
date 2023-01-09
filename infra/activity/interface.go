package activity

import "github.com/labstack/echo/v4"

type ActivityRepository interface {
	GetAllActivity(c echo.Context) (res Response, err error)
	GetActivityByID(c echo.Context, id int) (res Response, err error)
	CreateActivity(c echo.Context, req Activity) (res Response, err error)
	UpdateActivity(c echo.Context, req Activity) (res Response, err error)
	DeleteActivity(c echo.Context, id int) (res Response, err error)
}

type ActivityService interface {
	GetAllActivity(c echo.Context) (res Response, err error)
	GetActivityByID(c echo.Context, id int) (res Response, err error)
	CreateActivity(c echo.Context, req Activity) (res Response, err error)
	UpdateActivity(c echo.Context, req Activity) (res Response, err error)
	DeleteActivity(c echo.Context, id int) (res Response, err error)
}
