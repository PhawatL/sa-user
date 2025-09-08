package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Role type
type Role string

const (
	PatientRole Role = "patient"
	DoctorRole  Role = "doctor"
	AdminRole   Role = "admin"
)

// User model
type User struct {
	ID          uuid.UUID      `gorm:"type:uuid;primary_key;" json:"id"`
	HospitalID  string         `gorm:"not null;unique" json:"hospital_id"`
	Email       string         `gorm:"not null;unique" json:"email"`
	Password    string         `gorm:"not null" json:"-"`
	FirstName   string         `gorm:"not null" json:"first_name"`
	LastName    string         `gorm:"not null" json:"last_name"`
	Gender      string         `gorm:"not null" json:"gender"`
	PhoneNumber string         `json:"phone_number"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

// Patient model
type Patient struct {
	UserID           uuid.UUID      `gorm:"type:uuid;primary_key" json:"user_id"`
	User             User           `gorm:"foreignKey:UserID;references:ID"`
	Address          *string         `json:"address,omitempty"`
	Allergies        *string         `json:"allergies,omitempty"`
	EmergencyContact *string         `json:"emergency_contact,omitempty"`
	BloodType        *string         `gorm:"size:5" json:"blood_type,omitempty"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"-"`
}

// Doctor model
type Doctor struct {
	UserID          uuid.UUID      `gorm:"type:uuid;primary_key" json:"user_id"`
	User            User           `gorm:"foreignKey:UserID;references:ID"`
	Specialty       string         `json:"specialty,omitempty"`
	Bio             string         `json:"bio,omitempty"`
	YearsExperience int            `json:"years_experience,omitempty"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`
}

type UserRole struct {
	UserID uuid.UUID `gorm:"type:uuid;primary_key" json:"user_id"`
	User   User      `gorm:"foreignKey:UserID;references:ID"`
	Role   Role      `gorm:"type:roles;primary_key" json:"role"`
}
