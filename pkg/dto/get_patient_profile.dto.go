package dto

type GetProfileResponseDto struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
}