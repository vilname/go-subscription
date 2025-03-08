package model

type WebCitrus struct {
	ProjectId int         `json:"project_id"`
	Fields    interface{} `json:"fields"`
}

type WebCitrusSuccessResponse struct {
	Data    WebCitrusCartResponse `json:"data"`
	Status  int                   `json:"status"`
	Success bool                  `json:"success"`
}

type WebCitrusCartResponse struct {
	CreatedAd      string        `json:"created_at"`
	CreatedVia     string        `json:"created_via"`
	ExpirationDate string        `json:"expiration_date"`
	Fields         CreateCartDto `json:"fields"`
	Link           string        `json:"link"`
	PassNumber     string        `json:"pass_number"`
	ProductId      int           `json:"product_id"`
	SerialNumber   string        `json:"serial_number"`
	UpdatedAt      string        `json:"updated_at"`
	Voided         bool          `json:"voided"`
}
