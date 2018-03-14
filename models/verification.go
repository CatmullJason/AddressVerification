package models

type Verification struct {
	Success bool                `json:"success"`
	Errors  []FieldError        `json:"errors"`
	Details VerificationDetails `json:"details"`
}
