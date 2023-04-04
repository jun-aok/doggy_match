package controller

import (
	"net/http"
	// "strconv"
	// "time"

	// "app/internal/model"

	"app/internal/model/entity"
	"app/internal/model/repository"
	"app/internal/model/service"
	"app/pkg/db"

	"github.com/labstack/echo/v4"
)

func (a *GoogleOnlyAuthedController) GetUserInfo(c echo.Context) error {
	db := db.GetMySQLConnection().GetDbConnection()
	tx, error := db.Begin()
	if error != nil {
		return c.String(http.StatusInternalServerError, error.Error())
	}
	ur := repository.NewUserRepository(db)
	user, error := ur.QueryByUserCode(a.FirebaseToken.Claims["user_id"].(string))
	if error != nil {
		return error
	}
	if user == nil {
		// 登録されていないユーザー
		return c.JSON(http.StatusOK, map[string]*entity.Profile{"profile": nil})
	}
	pr := repository.NewProfileRepository(db)
	cr := repository.NewCheckInRepositoryWithTransaction(tx)
	userInfo, error := service.QueryUserInfo(user.UserId, pr, cr)
	if error != nil {
		// 500
		return c.JSON(http.StatusInternalServerError, error.Error())
	}
	return c.JSON(http.StatusOK, map[string]*entity.UserInfo{"userInfo": userInfo})
}
