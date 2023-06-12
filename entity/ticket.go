package entity

type SendEmailRequest struct {
	User  string `json:"user"`
	Email string `json:"email"`
	Item  string `json:"item"`
}
