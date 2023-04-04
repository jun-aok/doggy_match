package service

import (
	"app/internal/model/repository"
)

func IsDogOwner(userId int, dogIds []int, repository repository.IDogRepository) bool {
	// repositoryから犬一覧を取得
	// dogIdsと照らして合っているか
	return true
}
