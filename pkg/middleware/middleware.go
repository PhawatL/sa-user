package middleware

import (
	"user-service/pkg/jwt"

	"github.com/gofiber/fiber/v2"
)

func JwtMiddleware(jwtService *jwt.JwtService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// get cookie named "access_token"
		token := c.Cookies("access_token")
		if token == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Missing or malformed JWT",
			})
		}

		claims, err := jwtService.Parse(token)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid JWT",
			})
		}

		// ผูก context ให้ handler ถัดไป
		c.Locals("userID", claims.UserID)
		c.Locals("role", claims.Role)
		// ใส่ scopes/roles ถ้าจำเป็น: c.Locals("scopes", claims.Scopes)

		return c.Next()
	}
}
