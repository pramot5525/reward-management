package input

import "github.com/pramot5525/reward-management/internal/core/domain"

type TierUsecase interface {
	GetTiers() ([]domain.LoyaltyTier, int64, error)
}
