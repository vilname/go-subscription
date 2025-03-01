package model

type CreateSubscriptionDto struct {
	Email string `json:"email" validate:"required"`
}
