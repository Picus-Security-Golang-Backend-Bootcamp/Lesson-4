package models

type Customer struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
	Email       string `json:"email"`
	Age         int    `json:"age"`
}
