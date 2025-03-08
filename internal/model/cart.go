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
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	BirdDate    string `json:"bird_date"`
	Link        string `json:"link"`
	PassNumber  string `json:"pass_number"`
}

type FieldCard struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	BirdDate    string `json:"bird_date"`
}
