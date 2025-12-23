package middleware

import (
	"pbi/internal/helper"
	"pbi/internal/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func AuthChecker(strict bool) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			if strict {
				return helper.Unauthorized(c, "Token tidak ditemukan")
			}
			return helper.BadRequest(c, "Authorization header is missing", nil)
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := utils.ValidateToken(tokenString)
		if err != nil || claims.UserID <= 0 {
			return helper.Unauthorized(c, "Token tidak valid atau expired")
		}

		c.Locals("user_id", claims.UserID)
		c.Locals("is_admin", claims.IsAdmin)
		return c.Next()
	}
}

func AdminChecker() fiber.Handler {
	panic("unimplemented")
}