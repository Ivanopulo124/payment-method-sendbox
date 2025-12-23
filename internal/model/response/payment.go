package response

type Payment struct {
	OrderNumber  string  `json:"order_number"`
	Status       string  `json:"status"`
	Amount       float64 `json:"amount"`
	CreatedAt    string  `json:"created_at"`
	PaymentUrl   string  `json:"payment_url"`
	PaymentToken string  `json:"payment_token"`
}
