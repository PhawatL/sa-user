package service

import (
	"context"
	"errors"
	"fmt"
	"user-service/pkg/constants"
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

func (s *UserService) Register(ctx context.Context, body *dto.PatientRegisterPatientRequestDto) (*dto.PatientRegisterResponseDto, error) {
	user := &models.User{
		ID:          utils.GenerateUUIDv7(),
		FirstName:   body.FirstName,
		LastName:    body.LastName,
		Gender:      body.Gender,
		Role:        constants.RolePatient,
		PhoneNumber: body.PhoneNumber,
	}
	patient := &models.Patient{
		UserID:           user.ID,
		HospitalID:       body.HospitalID,
		BirthDate:        utils.ParseNullableTime(body.BirthDate),
		IDCardNumber:     body.IDCardNumber,
		Address:          body.Address,
		Allergies:        body.Allergies,
		EmergencyContact: body.EmergencyContact,
		BloodType:        body.BloodType,
	}
	fmt.Println("Hello From service Update")

	hashedPassword, err := utils.HashPassword(body.Password)
	if err != nil {
		return &dto.PatientRegisterResponseDto{}, err
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
		return &dto.PatientRegisterResponseDto{}, err
	}

	return &dto.PatientRegisterResponseDto{Message: "User registered successfully"}, nil
}

func (s *UserService) PatientLogin(ctx context.Context, body *dto.PatientLoginRequestDto) (*dto.PatientLoginResponseDto, error) {

	user, err := s.userRepository.FindByHospitalID(ctx, body.HospitalID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}
	ok, err := utils.VerifyPassword(body.Password, user.Password)
	if !ok || err != nil {
		fmt.Println(ok, err)
		return nil, errors.New("invalid credentials")
	}
	// sign token
	token, err := s.jwtService.GenerateToken(user.ID.String(), constants.RolePatient)
	if err != nil {
		return nil, err
	}
	// set token in cookie
	return &dto.PatientLoginResponseDto{
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

func (s *UserService) UpdatePatientProfile(ctx context.Context, userID string, role string, body *dto.PatientUpdateProfileRequestDto) (*dto.PatientUpdateProfileResponeDto, error) {
	if role != "patient" {
		return &dto.PatientUpdateProfileResponeDto{}, nil
	}

	user := &models.User{
		FirstName:   body.FirstName,
		LastName:    body.LastName,
		PhoneNumber: body.PhoneNumber,
	}
	patient := &models.Patient{
		Address:          body.Address,
		Allergies:        body.Allergies,
		EmergencyContact: body.EmergencyContact,
	}

	fmt.Println("Hello from service update.")

	if err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := s.userRepository.UpdateUser(ctx, userID, user); err != nil {
			return err
		}
		if err := s.patientRepository.UpdatePatient(ctx, userID, patient); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return &dto.PatientUpdateProfileResponeDto{}, err
	}

	return &dto.PatientUpdateProfileResponeDto{Message: "User updated successfully"}, nil
}
