package model

type CreateSubscriptionDto struct {
	Email string `json:"email" validate:"required"`
}

type CreateSubscriptionResult struct {
	Hash string `json:"hash"`
}

type SubscriptionModel struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
}
