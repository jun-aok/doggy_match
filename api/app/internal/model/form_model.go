package model

import (
	"github.com/labstack/echo/v4"
)

func ParseBindingErrorsToErrorDictionary(errors []error) map[string]string {
	errorMap := make(map[string]string)
	if len(errors) > 0 {
		// 扱いやすい形に変換する
		for _, e := range errors {
			// errorを受け取っているが本来は初めから[]echo.BindingErrorを受け取るべきだが
			// このメソッドの利便性がほとんどなくなるためこんな風な実装にしている
			be := e.(*echo.BindingError)
			if be.HTTPError.Internal == nil {
				errorMap[be.Field] = be.HTTPError.Message.(string)
			} else {
				errorMap[be.Field] = be.HTTPError.Internal.Error()
			}
		}
	}
	return errorMap
}

// import (
// 	"unicode/utf8"

// 	"github.com/labstack/echo/v4"
// )

// type FormModel interface {
// 	Binda(c echo.Context) []echo.BindingError
// 	Bind(b func(c echo.Context) []echo.BindingError)
// 	Validate(v func())
// }

// func Bind(c echo.Context, f FormModel) (bool, []error) {
// 	errorMap := make(map[string]string)
// 	errors := f.Binda(c)
// 	errors := echo.FormFieldBinder(c).
// 		MustString("name", &m.Name).
// 		MustBindUnmarshaler("birthDate", &m.BithDate).
// 		MustBindUnmarshaler("gender", &m.Gender).
// 		BindErrors()
// 	if len(errors) > 0 {
// 		// 扱いやすい形に変換する
// 		for _, e := range errors {
// 			be := e.(*echo.BindingError)
// 			errorMap[be.Field] = be.HTTPError.Internal.Error()
// 		}
// 	}
// 	// bind時点でエラーになっていたらチェックしない
// 	if _, hasError := errorMap["name"]; !hasError {
// 		if utf8.RuneCountInString(m.Name) > 64 {
// 			errorMap["name"] = "名前は64文字以内で入力して下さい"
// 		}
// 	}
// 	return len(errors) == 0, errors
// }
