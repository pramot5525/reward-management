package input

import "github.com/pramot5525/reward-management/internal/core/domain"

type RedeemUsecase interface {
	RedeemReward(userID string, rewardID uint) (*domain.RewardCode, error)
	GetUserRedeemed(userID string) ([]domain.RewardTransaction, error)
}
