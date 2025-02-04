package httputil

type Response struct {
	Data interface{} `json:"data"`
}

type SimpleMessageResponse struct {
	Message string `json:"message"`
}

func NewResponse(data interface{}) *Response {
	return &Response{data}
}

func NewMessageResponse(message string) *SimpleMessageResponse {
	return &SimpleMessageResponse{Message: message}
}
