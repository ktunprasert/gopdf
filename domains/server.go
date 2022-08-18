package domains

type JsonResponse[T any] struct {
	Success bool   `json:"success"`
	Message string `json:"message"`

	Data T `json:"data"`
}

type ErrorResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
