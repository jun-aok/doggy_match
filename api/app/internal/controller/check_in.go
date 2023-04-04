package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	. "app/internal/model"
	"app/internal/model/entity/orm"
	"app/internal/model/repository"
	"app/internal/model/service"
	"app/internal/model/value"
	"app/pkg/db"

	"github.com/labstack/echo/v4"
	//"google.golang.org/grpc/credentials/alts/internal/handshaker/service"
)

// チェックインしているユーザーを取得
func (a *AuthedController) GetCheckIn(c echo.Context) error {
	return nil
}

func (a *AuthedController) CheckOut(c echo.Context) error {
	dbConnection := db.GetMySQLConnection().GetDbConnection()
	tx, error := dbConnection.Begin()
	if error != nil {
		return c.String(http.StatusInternalServerError, error.Error())
	}
	error = service.CheckOut(
		a.CurrentUser.UserId,
		repository.NewCheckInRepositoryWithTransaction(tx),
	)
	if error != nil {
		tx.Rollback()
		return c.String(http.StatusInternalServerError, error.Error())
	}
	tx.Commit()
	return c.JSON(http.StatusOK, "")
}

func (a *AuthedController) CheckIn(c echo.Context) error {
	dbConnection := db.GetMySQLConnection().GetDbConnection()
	var form checkInFormModel
	binded, err := form.Bind(c, a.CurrentUser, repository.NewDogRepository(dbConnection))
	if !binded {
		return c.JSON(http.StatusBadRequest, err)
	}

	dogs := make([]*orm.Dog, 0, 5)

	if len(form.DogIds) > 0 {
		// 犬のチェックインがない場合
		r := repository.NewDogRepository(dbConnection)
		var error error
		dogs, error = r.QueryByUserId(a.CurrentUser.UserId)
		if error != nil {
			return c.String(http.StatusInternalServerError, error.Error())
		}
	}

	checkIn, error := value.NewCheckIn(
		a.CurrentUser.UserId,
		form.DogIds,
		dogs,
		form.Latitude,
		form.Longitude,
	)
	if error != nil {
		return c.String(http.StatusInternalServerError, error.Error())
	}
	if !checkIn.IsValid() {
		return c.String(404, error.Error())
	}
	tx, error := dbConnection.Begin()
	if error != nil {
		return c.String(http.StatusInternalServerError, error.Error())
	}
	error = service.CheckIn(
		a.CurrentUser.UserId,
		checkIn,
		repository.NewCheckInRepositoryWithTransaction(tx),
	)
	if error != nil {
		tx.Rollback()
		return c.String(http.StatusInternalServerError, error.Error())
	}
	tx.Commit()
	return c.JSON(http.StatusOK, map[string]*value.CheckIn{"checkIn": checkIn})
}

type checkInFormModel struct {
	// 緯度
	Latitude Position `form:"latitude"`
	// 経度
	Longitude Position `form:"longitude"`
	// type
	DogIds []int `form:"dog_ids"`
}

func (m *checkInFormModel) Bind(c echo.Context, user *orm.User, dr *repository.DogRepository) (bool, map[string]string) {
	errors := echo.FormFieldBinder(c).
		MustBindUnmarshaler("latitude", &m.Latitude).
		MustBindUnmarshaler("longitude", &m.Longitude).
		BindErrors()
	fmt.Println(&m.Latitude)
	fmt.Println(&m.Longitude)
	if errors != nil {
		errorMap := ParseBindingErrorsToErrorDictionary(errors)
		return len(errorMap) == 0, errorMap
	}
	array := make([]int, 0, 5)
	dogIds := c.FormValue("dog_ids")
	if dogIds != "0" {
		for _, v := range strings.Split(dogIds, ",") {
			dogId, error := strconv.Atoi(v)
			if error != nil {

			}
			array = append(array, dogId)
		}
	}
	m.DogIds = array

	errorMap := ParseBindingErrorsToErrorDictionary(errors)
	return len(errorMap) == 0, errorMap
}
