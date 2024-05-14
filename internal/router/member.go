package router

import (
	"go-clean-arch-game-server/internal/api"

	"github.com/gofiber/contrib/casbin"
	"github.com/gofiber/fiber/v2"
)

type MemberRouter interface {
	Init(root *fiber.Router, authzMiddleware *casbin.Middleware)
}

type memberRouter struct {
	api api.MemberHttpApi
}

func NewMemberRouter(api api.MemberHttpApi) MemberRouter {
	return &memberRouter{api: api}
}

func (mr *memberRouter) Init(root *fiber.Router, authzMiddleware *casbin.Middleware) {
	memberRouter := (*root).Group("/member")
	{
		// commands

		memberRouter.Post("", mr.api.AddMember)
		// cragRouter.Put("/:id", mr.api.UpdateCrag)
		// cragRouter.Delete("/:id", mr.api.DeleteCrag)
		// queries
		// cragRouter.Get("", mr.api.GetCrags)
		// cragRouter.Get("/:id", mr.api.GetCrag)
	}

}
