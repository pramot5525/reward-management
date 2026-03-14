package output

import "github.com/nocnoc-thailand/reward-management/internal/core/domain"

type RewardCodeRepository interface {
	CreateBatch(codes []domain.RewardCode) error
	Update(code domain.RewardCode) (*domain.RewardCode, error)
	FindByRewardID(rewardID uint) ([]domain.RewardCode, error)
	FindAvailableByRewardID(rewardID uint) (*domain.RewardCode, error)
}
