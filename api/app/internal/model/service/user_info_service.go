package service

import (
	"app/internal/model/entity"
	"app/internal/model/repository"
)

func QueryUserInfo(userId int, pr repository.IProfileRepository, cr repository.ICheckInRepository) (*entity.UserInfo, error) {
	profile, error := pr.Query(userId)
	if error != nil {
		return nil, error
	}
	checkIn, error := cr.Query(userId)
	if error != nil {
		return nil, error
	}
	return &entity.UserInfo{
		Profile: profile,
		CheckIn: checkIn,
	}, nil
}
