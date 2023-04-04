package entity

import (
	model "app/internal/model/enum"
	"time"
)

type Profile struct {
	Email string        `json:"email"`
	Token string        `json:"token"`
	User  *ProfileUser  `json:"user"`
	Dogs  []*ProfileDog `json:"dogs"`
}

type ProfileUser struct {
	UserId    int           `json:"userId"`
	Name      string        `json:"name"`
	BirthDate time.Time     `json:"birthDate"`
	Gender    *model.Gender `json:"gender"`
}

type ProfileDog struct {
	DogId       *int               `json:"dogId"`
	Name        *string            `json:"name"`
	BirthDate   *time.Time         `json:"birthDate"`
	Gender      *model.Gender      `json:"gender"`
	Personality *model.Personality `json:"personality"`
}
