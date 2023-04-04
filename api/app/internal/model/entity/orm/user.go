package orm

import (
	. "app/internal/model"
	. "app/internal/model/enum"
	"time"
)

type User struct {
	UserId    int        `json:"userId"`
	UserCode  string     `json:"userCode"`
	Email     string     `json:"email"`
	Name      Name       `json:"name"`
	BirthDate Date       `json:"birthDate"`
	Gender    Gender     `json:"gender"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt` // nullable
}
