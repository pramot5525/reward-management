package services

import (
	"github.com/nocnoc-thailand/reward-management/internal/core/domain"
	"github.com/nocnoc-thailand/reward-management/internal/core/ports/output"
)

type rewardCodeService struct {
	rewardCodeRepo output.RewardCodeRepository
	cache          output.CachePort
}

func NewRewardCodeService(rewardCodeRepo output.RewardCodeRepository, cache output.CachePort) *rewardCodeService {
	return &rewardCodeService{rewardCodeRepo: rewardCodeRepo, cache: cache}
}

func (s *rewardCodeService) CreateRewardCodes(rewardID uint, codes []domain.RewardCode) error {
	for i := range codes {
		codes[i].RewardID = rewardID
	}
	return s.rewardCodeRepo.CreateBatch(codes)
}

func (s *rewardCodeService) UpdateRewardCode(code domain.RewardCode) (*domain.RewardCode, error) {
	return s.rewardCodeRepo.Update(code)
}

func (s *rewardCodeService) GetRewardCodes(rewardID uint) ([]domain.RewardCode, error) {
	return s.rewardCodeRepo.FindByRewardID(rewardID)
}
