package api

import (
	"go-clean-arch-game-server/internal/app"
	"go-clean-arch-game-server/internal/common/errors"
	"go-clean-arch-game-server/internal/common/responses"
	"go-clean-arch-game-server/internal/common/validator"
	dto "go-clean-arch-game-server/internal/domain/dto/member"

	"github.com/gofiber/fiber/v2"
)

type MemberHttpApi interface {
	AddMember(ctx *fiber.Ctx) error
}

type memberHttpApi struct {
	memberApp app.Application
}

// NewHandler Constructor
func NewMemberHttpApi(memberApp app.Application) MemberHttpApi {
	return &memberHttpApi{memberApp: memberApp}
}

// AddMember Add a new member
// @Summary Add a new member
// @Tags Member
// @Accept json
// @Produce json
// @Param member body dto.AddMemberRequest true "The member data"
// @Success 200 {object} responses.General
// @Failure 500 {object} responses.General
// @Failure 400 {object} responses.General
// @Router /member [post]
func (cr *memberHttpApi) AddMember(ctx *fiber.Ctx) error {
	context := ctx.Context()

	req := new(dto.AddMemberRequest)

	if err := ctx.BodyParser(&req); err != nil {
		return errors.ErrBadRequest
	}

	if err := validator.GetValidator().Validate(req); err != nil {
		return errors.ErrBadRequest
	}

	if err := cr.memberApp.Commands.AddMemberHandler.Handle(context, req); err != nil {
		return err
	}
	return responses.DefaultSuccessResponse.JSON(ctx)
}
