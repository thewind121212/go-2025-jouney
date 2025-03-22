package _structs

import "time"

type User struct {
	ID           string    `json:"id"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	Name         string    `json:"name"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	TotpSecret   string    `json:"totpSecret"`
	Bio          string    `json:"bio"`
	Country      string    `json:"country"`
	Gender       string    `json:"gender"`
	TimeZone     string    `json:"timeZone"`
	BirthDate    time.Time `json:"birthDate"`
	AvatarUrl    string    `json:"avatarUrl"`
	RecoveryCode string    `json:"recoveryCode"`
	TwoFaStatus  string    `json:"twoFaStatus"`
}
