package entity

import "app/internal/model/value"

type UserInfo struct {
	Profile *Profile       `json:"profile"`
	CheckIn *value.CheckIn `json:"checkIn"`
}
