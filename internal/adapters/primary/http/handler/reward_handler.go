package handler

import (
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/pramot5525/reward-management/internal/adapters/primary/http/dto"
	"github.com/pramot5525/reward-management/internal/adapters/primary/http/middleware"
	"github.com/pramot5525/reward-management/internal/core/domain"
	"github.com/pramot5525/reward-management/internal/core/ports/input"
)

var validate = validator.New()

type RewardHandler struct {
	usecase input.RewardUsecase
}

func NewRewardHandler(usecase input.RewardUsecase) *RewardHandler {
	return &RewardHandler{usecase: usecase}
}

func (h *RewardHandler) Create(c *fiber.Ctx) error {
	var req dto.CreateRewardRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if err := validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	startDate, err := time.Parse("2006-01-02", req.StartDate)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid startDate, use YYYY-MM-DD"})
	}
	endDate, err := time.Parse("2006-01-02", req.EndDate)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid endDate, use YYYY-MM-DD"})
	}

	tiers := make([]domain.Tier, len(req.Tiers))
	for i, t := range req.Tiers {
		tiers[i] = domain.Tier{TierID: t}
	}
	rewardUsers := make([]domain.RewardUser, len(req.UserIDs))
	for i, u := range req.UserIDs {
		rewardUsers[i] = domain.RewardUser{UserID: u}
	}

	reward := domain.Reward{
		RedeemPoint:      req.RedeemPoint,
		StartDate:        startDate,
		EndDate:          endDate,
		ConditionBy:      domain.ConditionBy(req.ConditionBy),
		IsLimitRedeem:    req.IsLimitRedeem,
		CountLimitRedeem: req.CountLimitRedeem,
		Tiers:            tiers,
		RewardUsers:      rewardUsers,
		RewardInfo: domain.RewardInfo{
			TitleTH:            req.InfoTH.Title,
			SubtitleTH:         req.InfoTH.Subtitle,
			DescriptionTH:      req.InfoTH.Description,
			TermConditionTH:    req.InfoTH.TermCondition,
			BannerURL:          req.InfoTH.BannerURL,
			LogoURL:            req.InfoTH.LogoURL,
			IsCodeDisplayed:    req.InfoTH.IsCodeDisplayed,
			IsQrDisplayed:      req.InfoTH.IsQrDisplayed,
			IsBarcodeDisplayed: req.InfoTH.IsBarcodeDisplayed,
			TitleEN:            req.InfoEN.Title,
			SubtitleEN:         req.InfoEN.Subtitle,
			DescriptionEN:      req.InfoEN.Description,
			TermConditionEN:    req.InfoEN.TermCondition,
		},
	}

	created, err := h.usecase.CreateReward(reward)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(dto.ToRewardResponse(*created, c.Get("Accept-Language", "th")))
}

func (h *RewardHandler) Update(c *fiber.Ctx) error {
	var req dto.UpdateRewardRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if err := validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	reward := domain.Reward{
		ID:               req.ID,
		RedeemPoint:      req.RedeemPoint,
		IsLimitRedeem:    req.IsLimitRedeem,
		CountLimitRedeem: req.CountLimitRedeem,
	}
	if req.StartDate != "" {
		t, err := time.Parse("2006-01-02", req.StartDate)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid startDate"})
		}
		reward.StartDate = t
	}
	if req.EndDate != "" {
		t, err := time.Parse("2006-01-02", req.EndDate)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid endDate"})
		}
		reward.EndDate = t
	}
	if req.InfoTH != nil {
		reward.RewardInfo.TitleTH = req.InfoTH.Title
		reward.RewardInfo.SubtitleTH = req.InfoTH.Subtitle
		reward.RewardInfo.DescriptionTH = req.InfoTH.Description
		reward.RewardInfo.TermConditionTH = req.InfoTH.TermCondition
		reward.RewardInfo.BannerURL = req.InfoTH.BannerURL
		reward.RewardInfo.LogoURL = req.InfoTH.LogoURL
	}
	if req.InfoEN != nil {
		reward.RewardInfo.TitleEN = req.InfoEN.Title
		reward.RewardInfo.SubtitleEN = req.InfoEN.Subtitle
		reward.RewardInfo.DescriptionEN = req.InfoEN.Description
		reward.RewardInfo.TermConditionEN = req.InfoEN.TermCondition
	}

	updated, err := h.usecase.UpdateReward(reward)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(dto.ToRewardResponse(*updated, c.Get("Accept-Language", "th")))
}

func (h *RewardHandler) GetAvailable(c *fiber.Ctx) error {
	userID := c.Locals(middleware.UserIDKey).(string)

	rewards, err := h.usecase.GetRewards(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	lang := c.Get("Accept-Language", "th")
	res := make([]dto.RewardResponse, len(rewards))
	for i, r := range rewards {
		res[i] = dto.ToRewardResponse(r, lang)
	}
	return c.JSON(res)
}

func (h *RewardHandler) GetByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("rewardId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid rewardId"})
	}

	reward, err := h.usecase.GetRewardByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	if reward == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "reward not found"})
	}

	return c.JSON(dto.ToRewardResponse(*reward, c.Get("Accept-Language", "th")))
}

func (h *RewardHandler) GetList(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	rewards, total, err := h.usecase.GetRewardList(page, limit)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	lang := c.Get("Accept-Language", "th")
	items := make([]dto.RewardResponse, len(rewards))
	for i, r := range rewards {
		items[i] = dto.ToRewardResponse(r, lang)
	}
	return c.JSON(dto.PaginatedRewardResponse{Items: items, Total: total, Page: page, Limit: limit})
}

func (h *RewardHandler) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("rewardId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid rewardId"})
	}

	if err := h.usecase.DeleteReward(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
