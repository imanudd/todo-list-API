package activity

import (
	"log"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type ActivityRepositoryImpl struct {
	db *gorm.DB
}

func NewActivityRepositoryImpl(connectionStr string) ActivityRepository {
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
	return &ActivityRepositoryImpl{
		db: db,
	}
}

func (repo *ActivityRepositoryImpl) GetAllActivity(c echo.Context) (res Response, err error) {
	activities := []Activity{}
	result := repo.db.Find(&activities)
	if result.Error != nil {
		panic("Gagal Show All Activities")
	}
	res.Data = activities
	return res, nil
}

func (repo *ActivityRepositoryImpl) GetActivityByID(c echo.Context, id int) (res Response, err error) {
	var activity Activity
	result := repo.db.Where("activity_id = ?", id).Find(&activity)
	if result.Error != nil {
		return res, err
	}

	if activity.Activity_id == 0 {
		return res, err
	} else {
		res.Data = activity
		return res, nil
	}

}

func (repo *ActivityRepositoryImpl) CreateActivity(c echo.Context, req Activity) (res Response, err error) {
	activity := Activity{
		Title:      req.Title,
		Email:      req.Email,
		Created_at: time.Now().Local(),
		Updated_at: time.Now().Local(),
	}
	result := repo.db.Omit("activity_id").Create(&activity)
	if result.Error != nil {
		panic(err)
	}
	// I Dont Know Why Gorm Is Not Returning Back Of activity_id
	// I Think Gorm Doesn't support for returing ID ?
	// Iam so sorry for cannot fix it yet
	res.Data = activity
	return res, nil

}

func (repo *ActivityRepositoryImpl) UpdateActivity(c echo.Context, req Activity) (res Response, err error) {
	activity := Activity{
		Title:      req.Title,
		Updated_at: time.Now().Local(),
	}

	result := repo.db.Model(&activity).
		Where("activity_id", req.Activity_id).
		Select("title", "updated_at").
		Updates(activity)

	if result.Error != nil {
		return res, err
	}
	getData, _ := repo.GetActivityByID(c, req.Activity_id)
	res.Data = getData.Data
	return res, nil
}

func (repo *ActivityRepositoryImpl) DeleteActivity(c echo.Context, id int) (res Response, err error) {
	var activity Activity
	checkId, _ := repo.GetActivityByID(c, id)
	if checkId.Data == nil {
		res.Status = "Not Found"
		return res, err
	} else {
		result := repo.db.Where("activity_id = ?", id).Delete(&activity)
		if result.Error != nil {
			panic(err)
		}
		return res, nil
	}
}
