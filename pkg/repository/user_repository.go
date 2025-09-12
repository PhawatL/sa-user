package repository

import (
	"context"
	"user-service/pkg/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

// type UserRepositoryI interface {
// 	FindByID(ctx context.Context, id string) (*models.User, error)
// 	FindByEmail(ctx context.Context, email string) (*models.User, error)
// 	Create(ctx context.Context, user *models.User) error
// }

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) FindByHospitalID(ctx context.Context, hospitalID string) (*models.User, error) {
	var user models.User
	if err := r.db.WithContext(ctx).Joins("JOIN patients ON patients.user_id = users.id").Where("patients.hospital_id = ?", hospitalID).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// type GetProfileResponseDto struct {
// 	FirstName        string     `json:"first_name"`
// 	LastName         string     `json:"last_name"`
// 	Gender           string     `json:"gender"`
// 	PhoneNumber      string     `json:"phone_number"`
// 	HospitalID       string     `json:"hospital_id"`
// 	BirthDate        *time.Time `json:"birth_date"`
// 	IDCardNumber     *string    `json:"id_card_number"`
// 	Address          *string    `json:"address"`
// 	Allergies        *string    `json:"allergies"`
// 	EmergencyContact *string    `json:"emergency_contact"`
// 	BloodType        *string    `json:"blood_type"`
// }

func (r *UserRepository) FindPatientByID(ctx context.Context, id string) (*models.User, error) {
	var user models.User
	if err := r.db.WithContext(ctx).Preload("Patient").Where("id = ?", id).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

	}
	return &user, nil

}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Create(ctx context.Context, user *models.User) error {
	if err := r.db.WithContext(ctx).Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) CreateTx(tx *gorm.DB, u *models.User) error {
	return tx.Create(u).Error
}
