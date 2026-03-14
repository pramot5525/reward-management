package errs

import "github.com/gofiber/fiber/v2"

type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *AppError) Error() string { return e.Message }

var (
	ErrFullyRedeemed        = &AppError{Code: 400001, Message: "fully redeemed"}
	ErrRewardExpired        = &AppError{Code: 400002, Message: "reward expired"}
	ErrFullyRedeemedForUser = &AppError{Code: 400003, Message: "fully redeemed for user"}
	ErrNoCodeRemaining      = &AppError{Code: 400004, Message: "no reward code remaining"}
	ErrRewardNotFound       = &AppError{Code: 404001, Message: "reward not found"}
)

func RespondError(c *fiber.Ctx, status int, err *AppError) error {
	return c.Status(status).JSON(err)
}
