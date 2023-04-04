package orm

import (
	. "app/internal/model"
	. "app/internal/model/enum"
	"time"
)

type Dog struct {
	DogId       int         `json:"dogId"`
	UserId      int         `json:"userId"`
	Name        Name        `json:"name"`
	BirthDate   Date        `json:"birthDate"`
	Gender      Gender      `json:"gender"`
	Personality Personality `json:"personality"`
	CreatedAt   time.Time   `json:"createdAt"`
	UpdatedAt   time.Time   `json:"updatedAt"`
	DeletedAt   *time.Time  `json:"deletedAt` // nullable
}
