package controller

import (
	"app/internal/model/authenticator"
	"app/pkg/logger"
	"net/http"

	"firebase.google.com/go/v4/auth"

	"github.com/labstack/echo/v4"
)

type GoogleOnlyAuthedController struct {
	FirebaseToken *auth.Token
}

func (a *GoogleOnlyAuthedController) Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token, error := authenticator.Auth(c)
		if error != nil {
			logger.Error(error)
			// 401
			return c.String(http.StatusUnauthorized, "")
		}
		a.FirebaseToken = token
		return next(c)
	}
}
