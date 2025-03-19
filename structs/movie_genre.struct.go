package _structs

import "time"

type Genre struct {
	Name      string    `json:"name"`
	Desc      string    `json:"description"`
	Slug      string    `json:"slug"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
