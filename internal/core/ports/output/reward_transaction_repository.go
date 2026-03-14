package output

import "github.com/pramot5525/reward-management/internal/core/domain"

type RewardTransactionRepository interface {
	Create(tx domain.RewardTransaction) (*domain.RewardTransaction, error)
	FindByUserID(userID string) ([]domain.RewardTransaction, error)
	CountByUserIDAndRewardID(userID string, rewardID uint) (int64, error)
}
