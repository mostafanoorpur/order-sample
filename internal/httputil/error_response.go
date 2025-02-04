package httputil

type ErrorResponse struct {
	Status  int    `json:"status_code"`
	Message string `json:"message"`
}
