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

// Handler functions
// PatientRegister godoc
// @Summary Register a new patient
// @Description Register a new patient in the system
// @Tags patients
// @Accept  json
// @Produce  json
// @Param patient body dto.PatientRegisterPatientRequestDto true "Patient registration data"
// @Success 201 {object} dto.PatientRegisterResponseDto "Patient registered successfully"
// @Failure 400 {object} response.ErrorResponse "Invalid request body"
// @Failure 500 {object} response.ErrorResponse "Failed to register user"
// @Router /patient/register [post]
func (h *UserHandler) PatientRegister(ctx *fiber.Ctx) error {
	fmt.Println("Register endpoint hit")
	var body dto.PatientRegisterPatientRequestDto
	if err := ctx.BodyParser(&body); err != nil {
		return response.BadRequest(ctx, "Invalid request body "+err.Error())
	}

	res, err := h.userService.Register(ctx.Context(), &body)

	if err != nil {
		return response.InternalServerError(ctx, "Failed to register user: "+err.Error())
	}

	return response.Created(ctx, res)
}

func (h *UserHandler) PatientLogin(ctx *fiber.Ctx) error {
	var body dto.PatientLoginRequestDto
	if err := ctx.BodyParser(&body); err != nil {
		return response.BadRequest(ctx, "Invalid request body "+err.Error())
	}
	res, err := h.userService.PatientLogin(ctx.Context(), &body)
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
