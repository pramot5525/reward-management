package output

import "github.com/nocnoc-thailand/reward-management/internal/core/domain"

type RewardRepository interface {
	Create(reward domain.Reward) (*domain.Reward, error)
	Update(reward domain.Reward) (*domain.Reward, error)
	FindByID(id uint) (*domain.Reward, error)
	FindAvailableForUser(userID string) ([]domain.Reward, error)
	FindAll(offset, limit int) ([]domain.Reward, int64, error)
	SoftDelete(id uint) error
}
