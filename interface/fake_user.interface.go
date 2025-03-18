package _interface

import "time"

type Movie struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	ReleaseYear   int       `json:"releaseYear"`
	Duration      int       `json:"duration"`
	ThumbnailUrl  string    `json:"thumbnailUrl"`
	HlsFilePathS3 string    `json:"hlsFilePathS3"`
	Views         int       `json:"views"`
	Likes         int       `json:"likes"`
	Dislikes      int       `json:"dislikes"`
	Status        string    `json:"status"` // Adjust this field if 'status' is a custom type (e.g., an enum)
	IsPublished   bool      `json:"isPublished"`
	TypeID        string    `json:"typeId"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

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
