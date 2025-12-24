// Di httproute.go
package http

import (
	"fmt"
	"pbi/internal/config/container"
	rest "pbi/internal/server/http/handler"

	"github.com/gofiber/fiber/v2"
)

func HttpRouteInit(r *fiber.App, containerConf *container.Container) {
	fmt.Println("=== Initializing HTTP Routes ===")

	// Auth Route
	api := r.Group("/api/v1")
	fmt.Println("API Group created: /api/v1")

	rest.AuthRoute(api, containerConf.UserUsc)
	fmt.Println("Auth routes registered")
}
