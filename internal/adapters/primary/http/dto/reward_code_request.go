package dto

import "time"

type CreateRewardCodesRequest struct {
	RewardID    uint       `json:"rewardId" validate:"required"`
	Codes       []string   `json:"codes" validate:"required,min=1,max=1000"`
	ExpiredDate *time.Time `json:"expiredDate"`
}

type UpdateRewardCodeRequest struct {
	ID          uint `json:"id" validate:"required"`
	IsAvailable bool `json:"isAvailable"`
}

type RedeemRequest struct {
	RewardID uint `json:"rewardId" validate:"required"`
}
