package mysql

import (
	"github.com/pramot5525/reward-management/internal/adapters/secondary/mysql/model"
	"github.com/pramot5525/reward-management/internal/core/domain"
	"gorm.io/gorm"
)

type tierRepository struct {
	db *gorm.DB
}

func NewTierRepository(db *gorm.DB) *tierRepository {
	return &tierRepository{db: db}
}

func (r *tierRepository) FindAll() ([]domain.LoyaltyTier, int64, error) {
	var tiers []model.LoyaltyTier
	var total int64

	r.db.Model(&model.LoyaltyTier{}).Count(&total)

	if err := r.db.Find(&tiers).Error; err != nil {
		return nil, 0, err
	}

	result := make([]domain.LoyaltyTier, len(tiers))
	for i, t := range tiers {
		result[i] = domain.LoyaltyTier{
			ID:       t.ID,
			Code:     t.Code,
			Name:     t.Name,
			ImageURL: t.ImageURL,
		}
	}
	return result, total, nil
}
