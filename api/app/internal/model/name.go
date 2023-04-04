package model

import (
	"fmt"
	"unicode/utf8"
)

type Name string

func (n *Name) UnmarshalParam(src string) error {
	if src == "" {
		return fmt.Errorf("名前を入力してください")
	}
	if utf8.RuneCountInString(src) > 64 {
		return fmt.Errorf("名前は64文字以内で入力して下さい")
	}
	*n = Name(src)
	return nil
}
