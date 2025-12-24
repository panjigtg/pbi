package middleware

import (
	"pbi/internal/helper"
	"pbi/internal/utils"
	"strconv"
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
			return helper.BadRequest(c, "Authorization header is missing")
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
	return func(c *fiber.Ctx) error {
		isAdmin, ok := c.Locals("is_admin").(bool)
		if !ok || !isAdmin {
			return helper.Forbidden(c, "Akses khusus admin")
		}
		return c.Next()
	}
}


func OwnershipChecker(param string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		paramID, err := strconv.Atoi(c.Params(param))
		if err != nil {
			return helper.BadRequest(c, "ID tidak valid", err.Error())
		}

		userID := c.Locals("user_id").(int)
		isAdmin := c.Locals("is_admin").(bool)

		if !isAdmin && paramID != userID {
			return helper.Forbidden(c, "Akses ditolak")
		}
		return c.Next()
	}
}
