package model

type CreateCartDto struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email,omitempty"`
	BirdDate    string `json:"bird_date,omitempty"`
}

type CardElementDto struct {
	Hash string `json:"hash"`
}

type CardElementResult struct {
	CreatedAt      string    `json:"created_at"`
	UpdatedAt      string    `json:"updated_at"`
	CreatedVia     string    `json:"created_via"`
	SerialNumber   string    `json:"serial_number"`
	ProjectId      int       `json:"project_id"`
	PassNumber     int       `json:"pass_number"`
	ExpirationDate string    `json:"expiration_date"`
	Voided         string    `json:"voided"`
	Fields         FieldCard `json:"fields"`
}

type FieldCard struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	BirdDate    string `json:"bird_date"`
}
