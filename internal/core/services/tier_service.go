package services

import (
	"github.com/pramot5525/reward-management/internal/core/domain"
	"github.com/pramot5525/reward-management/internal/core/ports/output"
)

type tierService struct {
	tierRepo output.TierRepository
}

func NewTierService(tierRepo output.TierRepository) *tierService {
	return &tierService{tierRepo: tierRepo}
}

func (s *tierService) GetTiers() ([]domain.LoyaltyTier, int64, error) {
	return s.tierRepo.FindAll()
}
