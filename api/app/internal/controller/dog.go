package controller

import (
	"net/http"
	"time"

	. "app/internal/model"
	"app/internal/model/entity/orm"
	. "app/internal/model/enum"
	"app/internal/model/repository"
	"app/pkg/db"

	"github.com/labstack/echo/v4"
)

func (a *AuthedController) RegisterDog(c echo.Context) error {
	var form dogRegisterFormModel
	binded, err := form.Bind(c)
	if !binded {
		return c.JSON(http.StatusBadRequest, err)
	}

	dbConnection := db.GetMySQLConnection().GetDbConnection()
	ur := repository.NewUserRepository(dbConnection)
	user, error := ur.QueryByUserCode(a.FirebaseToken.Claims["user_id"].(string))
	if error != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	r := repository.NewDogRepository(dbConnection)
	dog := &orm.Dog{
		UserId:      user.UserId,
		Name:        form.Name,
		BirthDate:   form.BirthDate,
		Gender:      form.Gender,
		Personality: form.Personality,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	dogId, error := r.Regist(dog)
	if error != nil {
		return c.String(http.StatusInternalServerError, error.Error())
	}
	dog.DogId = dogId
	return c.JSON(http.StatusOK, "")

}

type dogRegisterFormModel struct {
	Name        Name        `form:"name"`
	BirthDate   Date        `form:"birthDate"`
	Gender      Gender      `form:"gender"`
	Personality Personality `form:"personality"`
}

func (m *dogRegisterFormModel) Bind(c echo.Context) (bool, map[string]string) {
	errors := echo.FormFieldBinder(c).
		MustBindUnmarshaler("name", &m.Name).
		MustBindUnmarshaler("birthDate", &m.BirthDate).
		MustBindUnmarshaler("gender", &m.Gender).
		MustBindUnmarshaler("personality", &m.Personality).
		BindErrors()

	errorMap := ParseBindingErrorsToErrorDictionary(errors)
	return len(errorMap) == 0, errorMap
}
