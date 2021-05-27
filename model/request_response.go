package model

import "time"

// AccountRequest for create/update account request
type AccountRequest struct {
	FullName      string `json:"full_name"`
	Email         string `json:"email"`
	PhoneNumber   string `json:"phone_number"`
	Username      string `json:"username"`
	Password      string `json:"password"`
}

// AccountResponse for create account response
type AccountResponse struct {
	FullName      string 	`json:"full_name"`
	Email         string 	`json:"email"`
	PhoneNumber   string 	`json:"phone_number"`
	Username      string 	`json:"username"`
	CreatedAt 	  time.Time `json:"created_at"`
}

// UpdateAccountRequest for update account request
type UpdateAccountRequest struct {
	FullName      string 	`json:"full_name"`
	Email         string 	`json:"email"`
	PhoneNumber   string 	`json:"phone_number"`
}