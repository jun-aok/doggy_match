package repository

import . "app/internal/model/entity/orm"

type IUserRepository interface {
	Regist(user *User) (int, error)
	Query(userId int) (*User, error)
	QueryByUserCode(userCode string) (*User, error)
	Update(userId int, user *User) error
}
