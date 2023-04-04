package model

import (
	"fmt"
	"strconv"
)

type Gender int

const (
	GenderEmpty Gender = 0 + iota
	Male
	Female
	NoAnswer
)

// intからGender変換する処理
func ItoGender(i int) (Gender, error) {
	m := map[int]Gender{
		1: Male,
		2: Female,
		3: NoAnswer,
	}
	if m[i] == 0 {
		return GenderEmpty, fmt.Errorf("Genderではありません")
	}
	return m[i], nil
}

func (g *Gender) UnmarshalParam(src string) error {
	i, error := strconv.Atoi(src)
	if error != nil {
		return error
	}
	gender, error := ItoGender(i)
	if error != nil {
		return error
	}
	*g = gender
	return error
}
