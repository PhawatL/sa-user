package repository

import (
	"context"
	"user-service/pkg/models"

	"gorm.io/gorm"
)

type PatientRepository struct {
	db *gorm.DB
}

// type UserRepositoryI interface {
// 	FindByID(ctx context.Context, id string) (*models.User, error)
// 	FindByEmail(ctx context.Context, email string) (*models.User, error)
// 	Create(ctx context.Context, user *models.User) error
// }

func NewPatientRepository(db *gorm.DB) *PatientRepository {
	return &PatientRepository{
		db: db,
	}
}

func (r *PatientRepository) FindByID(ctx context.Context, id string) (*models.Patient, error) {
	var patient models.Patient
	if err := r.db.WithContext(ctx).Where("user_id = ?", id).Preload("User").First(&patient).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &patient, nil
}

func (r *PatientRepository) Create(ctx context.Context, patient *models.Patient) error {
	if err := r.db.WithContext(ctx).Create(patient).Error; err != nil {
		return err
	}
	return nil
}

func (r *PatientRepository) CreateTx(tx *gorm.DB, p *models.Patient) error {
	return tx.Create(p).Error
}

func (r *PatientRepository) UpdatePatient(ctx context.Context, id string, patient *models.Patient) error {
	if err := r.db.WithContext(ctx).Where("user_id = ?", id).Updates(patient).Error; err != nil {
		return err
	}
	return nil
}
