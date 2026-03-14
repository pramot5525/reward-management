package input

import "github.com/nocnoc-thailand/reward-management/internal/core/domain"

type RewardCodeUsecase interface {
	CreateRewardCodes(rewardID uint, codes []domain.RewardCode) error
	UpdateRewardCode(code domain.RewardCode) (*domain.RewardCode, error)
	GetRewardCodes(rewardID uint) ([]domain.RewardCode, error)
}
