package output

import "github.com/pramot5525/reward-management/internal/core/domain"

type RewardCodeRepository interface {
	CreateBatch(codes []domain.RewardCode) error
	Update(code domain.RewardCode) (*domain.RewardCode, error)
	FindByRewardID(rewardID uint) ([]domain.RewardCode, error)
	FindAvailableByRewardID(rewardID uint) (*domain.RewardCode, error)
}
