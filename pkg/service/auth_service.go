package service

import (
	"context"
	"errors"
	"fmt"
	"user-service/pkg/dto"
	"user-service/pkg/jwt"
	"user-service/pkg/models"
	"user-service/pkg/repository"
	"user-service/pkg/utils"

	"gorm.io/gorm"
)

type UserService struct {
	db                *gorm.DB
	userRepository    *repository.UserRepository
	patientRepository *repository.PatientRepository
	jwtService        *jwt.JwtService
}

func NewUserService(db *gorm.DB, userRepo *repository.UserRepository, patientRepo *repository.PatientRepository, jwtService *jwt.JwtService) *UserService {
	return &UserService{
		db:                db,
		userRepository:    userRepo,
		patientRepository: patientRepo,
		jwtService:        jwtService,
	}
}

func (s *UserService) Register(ctx context.Context, body *dto.PostRegisterPatientRequestDto) (dto.PostRegisterResponseDto, error) {
	user := &models.User{
		ID:          utils.GenerateUUIDv7(),
		Email:       body.Email,
		FirstName:   body.FirstName,
		LastName:    body.LastName,
		HospitalID:  body.HospitalID,
		PhoneNumber: body.PhoneNumber,
	}
	patient := &models.Patient{
		UserID:           user.ID,
		Address:          body.Address,
		Allergies:        body.Allergies,
		EmergencyContact: body.EmergencyContact,
		BloodType:        body.BloodType,
	}

	hashedPassword, err := utils.HashPassword(body.Password)
	if err != nil {
		return dto.PostRegisterResponseDto{}, err
	}
	user.Password = hashedPassword

	err = s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := s.userRepository.CreateTx(tx, user); err != nil {
			return err
		}
		if err := s.patientRepository.CreateTx(tx, patient); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return dto.PostRegisterResponseDto{}, err
	}

	return dto.PostRegisterResponseDto{Message: "User registered successfully"}, nil
}

func (s *UserService) Login(ctx context.Context, body *dto.PostLoginRequestDto) (dto.PostLoginResponseDto, error) {

	user, err := s.userRepository.FindByEmail(ctx, body.Email)
	if err != nil {
		return dto.PostLoginResponseDto{}, err
	}

	if user == nil {
		return dto.PostLoginResponseDto{}, errors.New("user not found")
	}

	ok, err := utils.VerifyPassword(body.Password, user.Password)
	if !ok || err != nil {
		fmt.Println(ok, err)
		return dto.PostLoginResponseDto{}, errors.New("invalid credentials")
	}

	// sign token
	token, err := s.jwtService.GenerateToken(user.ID.String())
	if err != nil {
		return dto.PostLoginResponseDto{}, err
	}
	// set token in cookie
	return dto.PostLoginResponseDto{
		AccessToken: token,
	}, nil
}

func (s *UserService) GetProfileByID(ctx context.Context, userID string) (*models.User, error) {
	user, err := s.userRepository.FindByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}
