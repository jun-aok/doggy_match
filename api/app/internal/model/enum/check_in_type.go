package model

// import (
// 	"fmt"
// 	"strconv"
// )

// type CheckInType int

// const (
// 	CheckInTypeEmpty CheckInType = 0 + iota
// 	Dog
// 	Human
// )

// // intからCheckInTypeEmpty変換する処理
// func ItoCheckInType(i int) (CheckInType, error) {
// 	m := map[int]CheckInType{
// 		1: Dog,
// 		2: Human,
// 	}
// 	if m[i] == 0 {
// 		return CheckInTypeEmpty, fmt.Errorf("CheckInTypeEmptyではありません")
// 	}
// 	return m[i], nil
// }

// func (c *CheckInType) UnmarshalParam(src string) error {
// 	i, error := strconv.Atoi(src)
// 	if error != nil {
// 		return error
// 	}
// 	checkInType, error := ItoCheckInType(i)
// 	if error != nil {
// 		return error
// 	}
// 	*c = checkInType
// 	return error
// }
