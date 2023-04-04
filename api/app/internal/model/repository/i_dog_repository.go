package repository

import (
	. "app/internal/model/entity/orm"
)

type IDogRepository interface {
	Regist(dog *Dog) (int, error)
	// 指定したユーザーIDのdog一覧を取得する
	QueryByUserId(userId int) ([]*Dog, error)
}
