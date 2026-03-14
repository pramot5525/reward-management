package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nocnoc-thailand/reward-management/internal/adapters/primary/http/dto"
	"github.com/nocnoc-thailand/reward-management/internal/core/domain"
	"github.com/nocnoc-thailand/reward-management/internal/core/ports/input"
)

type RewardCodeHandler struct {
	usecase input.RewardCodeUsecase
}

func NewRewardCodeHandler(usecase input.RewardCodeUsecase) *RewardCodeHandler {
	return &RewardCodeHandler{usecase: usecase}
}

func (h *RewardCodeHandler) Create(c *fiber.Ctx) error {
	var req dto.CreateRewardCodesRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if err := validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	codes := make([]domain.RewardCode, len(req.Codes))
	for i, code := range req.Codes {
		codes[i] = domain.RewardCode{
			Code:        code,
			IsAvailable: true,
			ExpiredDate: req.ExpiredDate,
		}
	}

	if err := h.usecase.CreateRewardCodes(req.RewardID, codes); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "codes created", "count": len(codes)})
}

func (h *RewardCodeHandler) Update(c *fiber.Ctx) error {
	var req dto.UpdateRewardCodeRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if err := validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	updated, err := h.usecase.UpdateRewardCode(domain.RewardCode{
		ID:          req.ID,
		IsAvailable: req.IsAvailable,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"id": updated.ID, "isAvailable": updated.IsAvailable})
}

func (h *RewardCodeHandler) GetByRewardID(c *fiber.Ctx) error {
	rewardID := c.QueryInt("rewardId", 0)
	if rewardID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "rewardId is required"})
	}

	codes, err := h.usecase.GetRewardCodes(uint(rewardID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	res := make([]fiber.Map, len(codes))
	for i, code := range codes {
		res[i] = fiber.Map{
			"id":          code.ID,
			"code":        code.Code,
			"isAvailable": code.IsAvailable,
			"expiredDate": code.ExpiredDate,
		}
	}
	return c.JSON(res)
}
