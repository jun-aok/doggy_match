package controller

import (
	"app/internal/model/authenticator"
	"app/internal/model/entity/orm"
	"app/internal/model/repository"
	"app/pkg/db"
	"app/pkg/logger"
	"net/http"

	"firebase.google.com/go/v4/auth"

	"github.com/labstack/echo/v4"
)

type AuthedController struct {
	FirebaseToken *auth.Token
	CurrentUser   *orm.User
}

func (a *AuthedController) Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token, error := authenticator.Auth(c)
		if error != nil {
			logger.Error(error)
			// 401
			return c.String(http.StatusUnauthorized, "")
		}
		a.FirebaseToken = token

		dbConnection := db.GetMySQLConnection().GetDbConnection()
		ur := repository.NewUserRepository(dbConnection)
		user, error := ur.QueryByUserCode(a.FirebaseToken.Claims["user_id"].(string))
		if error != nil {
			return error
		}
		// authed_controllerはuserもないとダメ
		// if user != nil {
		//	return fmt.Errorf("userが登録されていません")
		//}
		a.CurrentUser = user
		return next(c)
	}
}
