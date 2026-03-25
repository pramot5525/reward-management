package output

import "github.com/pramot5525/reward-management/internal/core/domain"

type TierRepository interface {
	FindAll() ([]domain.LoyaltyTier, int64, error)
}
