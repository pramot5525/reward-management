package mysql

import (
	"github.com/pramot5525/reward-management/internal/core/domain"
	"gorm.io/gorm"
)

type rewardTransactionRepository struct {
	db *gorm.DB
}

func NewRewardTransactionRepository(db *gorm.DB) *rewardTransactionRepository {
	return &rewardTransactionRepository{db: db}
}

func (r *rewardTransactionRepository) Create(tx domain.RewardTransaction) (*domain.RewardTransaction, error) {
	// TODO: implement
	return &tx, nil
}

func (r *rewardTransactionRepository) FindByUserID(userID string) ([]domain.RewardTransaction, error) {
	// TODO: implement
	return nil, nil
}

func (r *rewardTransactionRepository) CountByUserIDAndRewardID(userID string, rewardID uint) (int64, error) {
	// TODO: implement
	return 0, nil
}
