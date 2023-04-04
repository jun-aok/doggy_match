package repository

import (
	"app/internal/model/entity"
)

type IProfileRepository interface {
	Query(userId int) (*entity.Profile, error)
}
