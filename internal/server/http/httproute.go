package http

import (
	"pbi/internal/config/container"
	rest "pbi/internal/server/http/handler"

	"github.com/gofiber/fiber/v2"
)

func HttpRouteInit(r *fiber.App, containerConf *container.Container) {
	// Auth Route
	api := r.Group("/api/v1")
	
	rest.AuthRoute(api, containerConf.AuthUsc)
	rest.UserRoute(api, containerConf.UserUsc)
}
