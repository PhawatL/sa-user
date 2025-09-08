package dto

type PostLoginRequestDto struct {
	Email    string `json:"email" required:"required"`
	Password string `json:"password" required:"required"`
}

type PostLoginResponseDto struct {
	AccessToken string `json:"access_token"`
}