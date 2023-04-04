package controller

import (
	"net/http"
	"time"

	. "app/internal/model"
	. "app/internal/model/entity/orm"
	. "app/internal/model/enum"
	"app/internal/model/repository"
	"app/pkg/db"

	"github.com/labstack/echo/v4"
)

func (a *GoogleOnlyAuthedController) RegisterUser(c echo.Context) error {
	var form userRegisterFormModel
	binded, err := form.Bind(c)
	if !binded {
		//　バリデーションエラー
		return c.JSON(http.StatusBadRequest, err)
	}
	r := repository.NewUserRepository(db.GetMySQLConnection().GetDbConnection())
	token := a.FirebaseToken
	user := &User{
		UserCode:  token.Claims["user_id"].(string),
		Email:     token.Claims["email"].(string),
		Name:      form.Name,
		BirthDate: form.BirthDate,
		Gender:    form.Gender,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	userId, error := r.Regist(user)
	if error != nil {
		return c.String(http.StatusInternalServerError, error.Error())
	}
	user.UserId = userId
	return c.JSON(http.StatusOK, &userRegisterViewModel{
		Name:      user.Name,
		BirthDate: user.BirthDate,
		Gender:    user.Gender,
	})
}

// form model
type userRegisterFormModel struct {
	Name      Name   `form:"name"`
	BirthDate Date   `form:"birthDate"`
	Gender    Gender `form:"gender"`
}

func (m *userRegisterFormModel) Bind(c echo.Context) (bool, map[string]string) {
	errors := echo.FormFieldBinder(c).
		MustBindUnmarshaler("name", &m.Name).
		MustBindUnmarshaler("birthDate", &m.BirthDate).
		MustBindUnmarshaler("gender", &m.Gender).
		BindErrors()

	errorMap := ParseBindingErrorsToErrorDictionary(errors)
	return len(errorMap) == 0, errorMap
}

// view model
type userRegisterViewModel struct {
	Name      Name   `json:"name"`
	BirthDate Date   `json:"birthDate"`
	Gender    Gender `json:"gender"`
}
