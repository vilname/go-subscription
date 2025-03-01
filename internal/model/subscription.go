package model

type CreateSubscriptionDto struct {
	Email string `json:"email" validate:"required"`
}

type CreateSubscriptionResult struct {
	Hash string `json:"hash"`
}
