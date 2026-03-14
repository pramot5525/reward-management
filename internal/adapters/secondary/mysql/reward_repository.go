package mysql

import (
	"github.com/pramot5525/reward-management/internal/adapters/secondary/mysql/model"
	"github.com/pramot5525/reward-management/internal/core/domain"
	"gorm.io/gorm"
)

type rewardRepository struct {
	db *gorm.DB
}

func NewRewardRepository(db *gorm.DB) *rewardRepository {
	return &rewardRepository{db: db}
}

func (r *rewardRepository) Create(reward domain.Reward) (*domain.Reward, error) {
	m := toModel(reward)
	m.State = int(domain.RewardStateActive)

	if err := r.db.Create(&m).Error; err != nil {
		return nil, err
	}

	result := toDomain(m)
	return &result, nil
}

func (r *rewardRepository) Update(reward domain.Reward) (*domain.Reward, error) {
	var m model.Reward
	if err := r.db.First(&m, reward.ID).Error; err != nil {
		return nil, err
	}

	m.RedeemPoint = reward.RedeemPoint
	m.IsLimitRedeem = reward.IsLimitRedeem
	m.CountLimitRedeem = reward.CountLimitRedeem
	if !reward.StartDate.IsZero() {
		m.StartDate = reward.StartDate
	}
	if !reward.EndDate.IsZero() {
		m.EndDate = reward.EndDate
	}

	if err := r.db.Save(&m).Error; err != nil {
		return nil, err
	}

	// update reward info
	r.db.Model(&model.RewardInfo{}).Where("reward_id = ?", m.ID).Updates(map[string]any{
		"title_th":             reward.RewardInfo.TitleTH,
		"subtitle_th":          reward.RewardInfo.SubtitleTH,
		"description_th":       reward.RewardInfo.DescriptionTH,
		"term_condition_th":    reward.RewardInfo.TermConditionTH,
		"title_en":             reward.RewardInfo.TitleEN,
		"subtitle_en":          reward.RewardInfo.SubtitleEN,
		"description_en":       reward.RewardInfo.DescriptionEN,
		"term_condition_en":    reward.RewardInfo.TermConditionEN,
		"banner_url":           reward.RewardInfo.BannerURL,
		"logo_url":             reward.RewardInfo.LogoURL,
	})

	result := toDomain(m)
	return &result, nil
}

func (r *rewardRepository) FindByID(id uint) (*domain.Reward, error) {
	var m model.Reward
	err := r.db.Preload("RewardInfo").Preload("Tiers").Preload("RewardUsers").
		Where("id = ? AND state != ?", id, int(domain.RewardStateDeleted)).
		First(&m).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	result := toDomain(m)
	return &result, nil
}

func (r *rewardRepository) FindAvailableForUser(userID string) ([]domain.Reward, error) {
	var rewards []model.Reward
	err := r.db.Preload("RewardInfo").Preload("Tiers").Preload("RewardUsers").
		Where("state = ?", int(domain.RewardStateActive)).
		Find(&rewards).Error
	if err != nil {
		return nil, err
	}
	return toDomainList(rewards), nil
}

func (r *rewardRepository) FindAll(offset, limit int) ([]domain.Reward, int64, error) {
	var rewards []model.Reward
	var total int64

	r.db.Model(&model.Reward{}).Where("state != ?", int(domain.RewardStateDeleted)).Count(&total)

	err := r.db.Preload("RewardInfo").
		Where("state != ?", int(domain.RewardStateDeleted)).
		Offset(offset).Limit(limit).
		Find(&rewards).Error
	if err != nil {
		return nil, 0, err
	}
	return toDomainList(rewards), total, nil
}

func (r *rewardRepository) SoftDelete(id uint) error {
	return r.db.Model(&model.Reward{}).Where("id = ?", id).
		Update("state", int(domain.RewardStateDeleted)).Error
}

// --- mappers ---

func toModel(d domain.Reward) model.Reward {
	tiers := make([]model.Tier, len(d.Tiers))
	for i, t := range d.Tiers {
		tiers[i] = model.Tier{TierID: t.TierID}
	}
	users := make([]model.RewardUser, len(d.RewardUsers))
	for i, u := range d.RewardUsers {
		users[i] = model.RewardUser{UserID: u.UserID}
	}
	return model.Reward{
		RedeemPoint:      d.RedeemPoint,
		StartDate:        d.StartDate,
		EndDate:          d.EndDate,
		ConditionBy:      string(d.ConditionBy),
		IsLimitRedeem:    d.IsLimitRedeem,
		CountLimitRedeem: d.CountLimitRedeem,
		Tiers:            tiers,
		RewardUsers:      users,
		RewardInfo: model.RewardInfo{
			TitleTH:            d.RewardInfo.TitleTH,
			SubtitleTH:         d.RewardInfo.SubtitleTH,
			DescriptionTH:      d.RewardInfo.DescriptionTH,
			TermConditionTH:    d.RewardInfo.TermConditionTH,
			TitleEN:            d.RewardInfo.TitleEN,
			SubtitleEN:         d.RewardInfo.SubtitleEN,
			DescriptionEN:      d.RewardInfo.DescriptionEN,
			TermConditionEN:    d.RewardInfo.TermConditionEN,
			BannerURL:          d.RewardInfo.BannerURL,
			LogoURL:            d.RewardInfo.LogoURL,
			IsCodeDisplayed:    d.RewardInfo.IsCodeDisplayed,
			IsQrDisplayed:      d.RewardInfo.IsQrDisplayed,
			IsBarcodeDisplayed: d.RewardInfo.IsBarcodeDisplayed,
		},
	}
}

func toDomain(m model.Reward) domain.Reward {
	tiers := make([]domain.Tier, len(m.Tiers))
	for i, t := range m.Tiers {
		tiers[i] = domain.Tier{ID: t.ID, RewardID: t.RewardID, TierID: t.TierID}
	}
	users := make([]domain.RewardUser, len(m.RewardUsers))
	for i, u := range m.RewardUsers {
		users[i] = domain.RewardUser{ID: u.ID, RewardID: u.RewardID, UserID: u.UserID}
	}
	return domain.Reward{
		ID:               m.ID,
		RedeemPoint:      m.RedeemPoint,
		StartDate:        m.StartDate,
		EndDate:          m.EndDate,
		State:            domain.RewardState(m.State),
		ConditionBy:      domain.ConditionBy(m.ConditionBy),
		IsLimitRedeem:    m.IsLimitRedeem,
		CountLimitRedeem: m.CountLimitRedeem,
		Tiers:            tiers,
		RewardUsers:      users,
		RewardInfo: domain.RewardInfo{
			ID:                 m.RewardInfo.ID,
			RewardID:           m.RewardInfo.RewardID,
			TitleTH:            m.RewardInfo.TitleTH,
			SubtitleTH:         m.RewardInfo.SubtitleTH,
			DescriptionTH:      m.RewardInfo.DescriptionTH,
			TermConditionTH:    m.RewardInfo.TermConditionTH,
			TitleEN:            m.RewardInfo.TitleEN,
			SubtitleEN:         m.RewardInfo.SubtitleEN,
			DescriptionEN:      m.RewardInfo.DescriptionEN,
			TermConditionEN:    m.RewardInfo.TermConditionEN,
			BannerURL:          m.RewardInfo.BannerURL,
			LogoURL:            m.RewardInfo.LogoURL,
			IsCodeDisplayed:    m.RewardInfo.IsCodeDisplayed,
			IsQrDisplayed:      m.RewardInfo.IsQrDisplayed,
			IsBarcodeDisplayed: m.RewardInfo.IsBarcodeDisplayed,
		},
	}
}

func toDomainList(ms []model.Reward) []domain.Reward {
	result := make([]domain.Reward, len(ms))
	for i, m := range ms {
		result[i] = toDomain(m)
	}
	return result
}
