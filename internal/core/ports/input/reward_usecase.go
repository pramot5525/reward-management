package input

import "github.com/nocnoc-thailand/reward-management/internal/core/domain"

type RewardUsecase interface {
	CreateReward(reward domain.Reward) (*domain.Reward, error)
	UpdateReward(reward domain.Reward) (*domain.Reward, error)
	GetRewards(userID string) ([]domain.Reward, error)
	GetRewardByID(rewardID uint) (*domain.Reward, error)
	GetRewardList(page, limit int) ([]domain.Reward, int64, error)
	DeleteReward(rewardID uint) error
}
