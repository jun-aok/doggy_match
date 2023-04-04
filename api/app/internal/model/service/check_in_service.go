package service

import (
	"app/internal/model/repository"
	"app/internal/model/value"
)

func CheckIn(
	userId int,
	checkIn *value.CheckIn,
	checkInRepository repository.ICheckInRepository,
) error {
	//　指定したdogが
	// チェックインする前に現在のステートチェック
	// チェックイン済みだった場合、チェックアウトしてから新しくチェックインする
	existCheckIn, error := checkInRepository.Query(userId)
	if error != nil {
		return error
	}
	if existCheckIn != nil {
		// チェックアウト処理
		error := CheckOut(userId, checkInRepository)
		if error != nil {
			return error
		}
	}
	return checkInRepository.Regist(userId, checkIn)
}

func CheckOut(userId int, checkInRepository repository.ICheckInRepository) error {
	error := checkInRepository.Delete(userId)
	if error != nil {
		return error
	}
	// logの登録
	return nil
}

func GetNearCheckIn(userId int, court int, checkInRepository repository.ICheckInRepository) error {
	// 近くにいる人を取得する
	_, error := checkInRepository.Query(userId)
	if error != nil {
		return error
	}
	return nil
}
