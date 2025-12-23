package request

type PayRequest struct {
	Login         string `json:"login"`
	Password      string `json:"password"`
	PaymentNumber string `json:"payment_number"`
}

type PayRefund struct {
	Login         string  `json:"login"`
	Password      string  `json:"password"`
	PaymentNumber string  `json:"payment_number"`
	RefundAmount  float64 `json:"refund_amount"`
}
