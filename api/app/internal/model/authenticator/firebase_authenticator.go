package authenticator

import (
	"context"
	"fmt"
	"strings"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/labstack/echo/v4"
)

func Auth(c echo.Context) (*auth.Token, error) {
	auths := strings.Split(c.Request().Header.Get("Authorization"), " ")
	if !(len(auths) == 2 && auths[0] == "Bearer") {
		// エラー
		return nil, fmt.Errorf("トークンがありません")
	}
	bearerToken := auths[1]
	ctx := context.Background()
	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		return nil, err
	}
	client, err := app.Auth(ctx)
	if err != nil {
		return nil, err
	}

	token, err := client.VerifyIDToken(ctx, bearerToken)
	if err != nil {
		return nil, err
	}
	return token, nil
}
