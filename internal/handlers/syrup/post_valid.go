package syrup

import (
	"discountdb-api/internal/models/syrup"
	"github.com/gofiber/fiber/v2"
)

// PostCouponValid godoc
// @Summary Report Valid Coupon
// @Description Report that a coupon code was successfully used
// @Tags syrup
// @Produce json
// @Param X-Syrup-API-Key header string false "Optional API key for authentication"
// @Param id path string true "The ID of the coupon"
// @Success 200 {object} syrup.Success "Successful response"
// @Header 200 {string} X-RateLimit-Limit "The maximum number of requests allowed per time window"
// @Header 200 {string} X-RateLimit-Remaining "The number of requests remaining in the time window"
// @Header 200 {string} X-RateLimit-Reset "The time when the rate limit window resets (Unix timestamp)"
// @Failure 400 {object} syrup.ErrorResponse "Bad Request"
// @Failure 401 {object} syrup.ErrorResponse "Unauthorized"
// @Failure 429 {object} syrup.ErrorResponse "Too Many Requests"
// @Header 429 {integer} X-RateLimit-RetryAfter "Time to wait before retrying (seconds)"
// @Failure 500 {object} syrup.ErrorResponse "Internal Server Error"
// @Router /syrup/coupons/valid/{id} [post]
func PostCouponValid(ctx *fiber.Ctx) error {
	// TODO: Implement this handler

	return ctx.Status(fiber.StatusInternalServerError).JSON(
		syrup.ErrorResponse{
			Error:   "NotImplemented",
			Message: "The API endpoint is not yet implemented",
		},
	)
}
