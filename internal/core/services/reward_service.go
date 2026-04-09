package services

import (
	"github.com/pramot5525/reward-management/internal/core/domain"
	"github.com/pramot5525/reward-management/internal/core/ports/output"
)

type rewardService struct {
	rewardRepo output.RewardRepository
	cache      output.CachePort
}

func NewRewardService(rewardRepo output.RewardRepository, cache output.CachePort) *rewardService {
	return &rewardService{rewardRepo: rewardRepo, cache: cache}
}

func (s *rewardService) CreateReward(reward domain.Reward) (*domain.Reward, error) {
	return s.rewardRepo.Create(reward)
}

func (s *rewardService) UpdateReward(reward domain.Reward) (*domain.Reward, error) {
	return s.rewardRepo.Update(reward)
}

func (s *rewardService) GetRewards(userID string) ([]domain.Reward, error) {
	return s.rewardRepo.FindAvailableForUser(userID)
}

func (s *rewardService) GetRewardByID(rewardID uint) (*domain.Reward, error) {
	return s.rewardRepo.FindByID(rewardID)
}

func (s *rewardService) GetRewardList(page, limit int) ([]domain.Reward, int64, error) {
	offset := (page - 1) * limit
	return s.rewardRepo.FindAll(offset, limit)
}

func (s *redeemService) DeleteRewardCode(rewardID uint) error {
	return s.rewardRepo.SoftDelete(rewardID)
}
