package mysql

import (
	"github.com/nocnoc-thailand/reward-management/internal/core/domain"
	"gorm.io/gorm"
)

type rewardCodeRepository struct {
	db *gorm.DB
}

func NewRewardCodeRepository(db *gorm.DB) *rewardCodeRepository {
	return &rewardCodeRepository{db: db}
}

func (r *rewardCodeRepository) CreateBatch(codes []domain.RewardCode) error {
	// TODO: implement
	return nil
}

func (r *rewardCodeRepository) Update(code domain.RewardCode) (*domain.RewardCode, error) {
	// TODO: implement
	return &code, nil
}

func (r *rewardCodeRepository) FindByRewardID(rewardID uint) ([]domain.RewardCode, error) {
	// TODO: implement
	return nil, nil
}

func (r *rewardCodeRepository) FindAvailableByRewardID(rewardID uint) (*domain.RewardCode, error) {
	// TODO: implement
	return nil, nil
}
