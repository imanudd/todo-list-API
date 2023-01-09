package activity

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func RegisterRoute(e *echo.Echo, svc *API) {
	e.GET("/activity-groups", svc.GetAllActivity)
	e.GET("/activity-groups/:id", svc.GetActivityByID)
	e.POST("/activity-groups", svc.CreateActivity)
	e.PATCH("/activity-groups/:id", svc.UpdateActivity)
	e.DELETE("/activity-groups/:id", svc.DeleteActivity)
}

type API struct {
	svc ActivityService
}

func NewActivityAPIImpl(svc ActivityService) *API {
	return &API{
		svc: svc,
	}
}

func (api *API) GetAllActivity(c echo.Context) error {
	res, err := api.svc.GetAllActivity(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, res)
}

func (api *API) GetActivityByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	res, err := api.svc.GetActivityByID(c, id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": err.Error(),
		})
	}
	if res.Data == nil {
		conv := strconv.Itoa(id)
		return c.JSON(http.StatusNotFound, map[string]string{
			"Status":  "Not Found",
			"Message": "Activity with ID " + conv + " Not Found",
		})
	}
	return c.JSON(http.StatusOK, res)

}

func (api *API) CreateActivity(c echo.Context) error {
	var json = new(Activity)
	if err := c.Bind(json); err != nil {
		return err
	}
	if json.Title == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Status":  "Bad Request",
			"Message": "Title Cannot Be Null",
		})
	}
	res, err := api.svc.CreateActivity(c, *json)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, res)
}

func (api *API) UpdateActivity(c echo.Context) error {
	var json = new(Activity)
	var res Response
	json.Activity_id, _ = strconv.Atoi(c.Param("id"))

	if err := c.Bind(json); err != nil {
		return c.JSON(http.StatusNotFound, res)
	}
	if json.Title == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Status":  "Bad Request",
			"Message": "Title Cannot Be Null",
		})
	}
	res, err := api.svc.UpdateActivity(c, *json)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}
	if res.Data == nil {
		conv := strconv.Itoa(json.Activity_id)
		return c.JSON(http.StatusNotFound, map[string]string{
			"Status":  "Not Found",
			"Message": "Activity with ID " + conv + " Not Found",
		})
	}
	return c.JSON(http.StatusCreated, res)
}

func (api *API) DeleteActivity(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	res, err := api.svc.DeleteActivity(c, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	if res.Status == "Not Found" {
		conv := strconv.Itoa(id)
		return c.JSON(http.StatusNotFound, map[string]string{
			"Status":  "Not Found",
			"Message": "Activity with ID " + conv + " Not Found",
		})
	}
	return c.JSON(http.StatusOK, res)
}
