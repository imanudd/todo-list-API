package activity

import (
	"fmt"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ActivityServiceImpl struct {
	repo ActivityRepository
}

func NewActivityServiceImpl(repo ActivityRepository) ActivityService {
	return &ActivityServiceImpl{
		repo: repo,
	}
}

func (svc *ActivityServiceImpl) GetAllActivity(c echo.Context) (res Response, err error) {
	res, err = svc.repo.GetAllActivity(c)
	if err != nil {
		panic("error parsing data")
	}
	res.Status = "Succsess"
	res.Message = "Succsess"
	return res, nil
}

func (svc *ActivityServiceImpl) GetActivityByID(c echo.Context, id int) (res Response, err error) {
	res, err = svc.repo.GetActivityByID(c, id)
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
func (svc *ActivityServiceImpl) CreateActivity(c echo.Context, req Activity) (res Response, err error) {
	res, err = svc.repo.CreateActivity(c, req)
	if err != nil {
		fmt.Println(err)
	}
	res.Status = "Succsess"
	res.Message = "Succsess"
	return res, nil
}
func (svc *ActivityServiceImpl) UpdateActivity(c echo.Context, req Activity) (res Response, err error) {
	res, err = svc.repo.UpdateActivity(c, req)
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

func (svc *ActivityServiceImpl) DeleteActivity(c echo.Context, id int) (res Response, err error) {
	res, err = svc.repo.DeleteActivity(c, id)
	if err != nil {
		panic(err)
	}
	if res.Status == "Not Found" {
		conv := strconv.Itoa(id)
		res.Message = "Activity with ID " + conv + " Not Found"
		return res, nil
	}
	res.Status = "Succsess"
	res.Message = "Succsess"
	return res, nil

}
