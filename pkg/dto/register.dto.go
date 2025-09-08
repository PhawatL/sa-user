package dto

type PostRegisterPatientRequestDto struct {
	HospitalID       string  `json:"hospital_id" validate:"required"`
	Email            string  `json:"email" validate:"required,email"`
	Password         string  `json:"password" validate:"required,min=6"`
	FirstName        string  `json:"first_name" validate:"required"`
	LastName         string  `json:"last_name" validate:"required"`
	PhoneNumber      string  `json:"phone_number"`
	Address          *string `json:"address"`
	Allergies        *string `json:"allergies"`
	EmergencyContact *string `json:"emergency_contact"`
	BloodType        *string `json:"blood_type" validate:"omitempty,oneof='A-' 'A+' 'B-' 'B+' 'AB-' 'AB+' 'O-' 'O+' 'A' 'B' 'AB' 'O'"`
}

type PostRegisterResponseDto struct {
	Message string `json:"message"`
}
