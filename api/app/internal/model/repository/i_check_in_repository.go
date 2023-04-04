package repository

import (
	"app/internal/model/value"
)

// checkin dog log全てこのレポジトリで行う
type ICheckInRepository interface {
	Query(userId int) (*value.CheckIn, error)
	Regist(userId int, checkIn *value.CheckIn) error
	Delete(userId int) error
	RegistLog(userId int, log *value.CheckIn) error
	CompleteLog(userId int) error
}
