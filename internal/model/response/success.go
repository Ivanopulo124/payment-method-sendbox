package response

type Success struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}
