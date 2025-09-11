package dto

type PatientUpdateProfileRequestDto struct {
	// user table
	FirstName   string `json:"first_name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	// patient table
	Address          *string `json:"address,omitempty"`
	Allergies        *string `json:"allergies,omitempty"`
	EmergencyContact *string `json:"emergency_contact,omitempty"`
}

type PatientUpdateProfileResponeDto struct {
	Message string `json:"message"`
}
