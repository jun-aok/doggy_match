package entity

import (
	"app/internal/model/entity/orm"
	model "app/internal/model/enum"
)

type ActiveUser struct {
	Name   string
	Gender model.Gender
	Dogs   []*orm.Dog
}
