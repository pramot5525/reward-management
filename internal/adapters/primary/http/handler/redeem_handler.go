package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nocnoc-thailand/reward-management/internal/adapters/primary/http/dto"
	"github.com/nocnoc-thailand/reward-management/internal/adapters/primary/http/middleware"
	"github.com/nocnoc-thailand/reward-management/internal/core/ports/input"
)

type RedeemHandler struct {
	usecase input.RedeemUsecase
}

func NewRedeemHandler(usecase input.RedeemUsecase) *RedeemHandler {
	return &RedeemHandler{usecase: usecase}
}

func (h *RedeemHandler) Redeem(c *fiber.Ctx) error {
	var req dto.RedeemRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if err := validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	userID := c.Locals(middleware.UserIDKey).(string)

	code, err := h.usecase.RedeemReward(userID, req.RewardID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"code":        code.Code,
		"expiredDate": code.ExpiredDate,
	})
}

func (h *RedeemHandler) GetUserRedeemed(c *fiber.Ctx) error {
	userID := c.Locals(middleware.UserIDKey).(string)

	transactions, err := h.usecase.GetUserRedeemed(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	res := make([]fiber.Map, len(transactions))
	for i, tx := range transactions {
		res[i] = fiber.Map{
			"id":           tx.ID,
			"rewardCodeId": tx.RewardCodeID,
			"status":       tx.Status,
		}
	}
	return c.JSON(res)
}
