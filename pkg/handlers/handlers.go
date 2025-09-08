package handlers

import (
	"fmt"
	"user-service/pkg/dto"
	response "user-service/pkg/response"
	"user-service/pkg/service"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService}
}

// Register godoc
// @Summary Register a new patient
// @Description Create a new user patient in the system
// @Tags User
// @Accept json
// @Produce json
// @Param register body dto.PostRegisterPatientRequestDto true "Register request body"
// @Success 201 {object} dto.PostRegisterResponseDto
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /api/v1/register [post]
func (h *UserHandler) Register(ctx *fiber.Ctx) error {
	fmt.Println("Register endpoint hit")
	var body dto.PostRegisterPatientRequestDto
	if err := ctx.BodyParser(&body); err != nil {
		return response.BadRequest(ctx, "Invalid request body")
	}

	res, err := h.userService.Register(ctx.Context(), &body)

	if err != nil {
		return response.InternalServerError(ctx, "Failed to register user: "+err.Error())
	}

	return response.Created(ctx, res)
}

func (h *UserHandler) Login(ctx *fiber.Ctx) error {
	var body dto.PostLoginRequestDto
	if err := ctx.BodyParser(&body); err != nil {
		return response.BadRequest(ctx, "Invalid request body")
	}
	res, err := h.userService.Login(ctx.Context(), &body)
	if err != nil {
		return response.Unauthorized(ctx, err.Error())
	}
	// set token in cookie
	ctx.Cookie(&fiber.Cookie{
		Name:  "access_token",
		Value: res.AccessToken,
	})
	return response.OK(ctx, res)
}

func (h *UserHandler) Profile(ctx *fiber.Ctx) error {
	userID := ctx.Locals("userID").(string)
	fmt.Println("Profile endpoint hit, userID:", userID)
	user, err := h.userService.GetProfileByID(ctx.Context(), userID)
	if err != nil {
		return response.InternalServerError(ctx, "Failed to get user profile: "+err.Error())
	}
	return response.OK(ctx, user)
}
