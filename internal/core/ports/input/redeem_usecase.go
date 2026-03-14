package input

import "github.com/nocnoc-thailand/reward-management/internal/core/domain"

type RedeemUsecase interface {
	RedeemReward(userID string, rewardID uint) (*domain.RewardCode, error)
	GetUserRedeemed(userID string) ([]domain.RewardTransaction, error)
}
