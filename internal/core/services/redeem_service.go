package services

import (
	"errors"

	"github.com/nocnoc-thailand/reward-management/internal/core/domain"
	"github.com/nocnoc-thailand/reward-management/internal/core/ports/output"
)

type redeemService struct {
	rewardRepo      output.RewardRepository
	rewardCodeRepo  output.RewardCodeRepository
	transactionRepo output.RewardTransactionRepository
	cache           output.CachePort
}

func NewRedeemService(
	rewardRepo output.RewardRepository,
	rewardCodeRepo output.RewardCodeRepository,
	transactionRepo output.RewardTransactionRepository,
	cache output.CachePort,
) *redeemService {
	return &redeemService{
		rewardRepo:      rewardRepo,
		rewardCodeRepo:  rewardCodeRepo,
		transactionRepo: transactionRepo,
		cache:           cache,
	}
}

func (s *redeemService) RedeemReward(userID string, rewardID uint) (*domain.RewardCode, error) {
	reward, err := s.rewardRepo.FindByID(rewardID)
	if err != nil {
		return nil, err
	}
	if reward == nil {
		return nil, errors.New("reward not found")
	}

	// TODO: validate expiry, check limit, atomic redeem
	code, err := s.rewardCodeRepo.FindAvailableByRewardID(rewardID)
	if err != nil {
		return nil, err
	}

	_ = userID
	return code, nil
}

func (s *redeemService) GetUserRedeemed(userID string) ([]domain.RewardTransaction, error) {
	return s.transactionRepo.FindByUserID(userID)
}
