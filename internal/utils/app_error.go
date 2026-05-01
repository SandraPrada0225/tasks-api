package utils

type AppError struct {
	Message string
	Code    int
}

func (e *AppError) Error() string {
	return e.Message
}

func NewError(message string, code int) *AppError {
	return &AppError{
		Message: message,
		Code:    code,
	}

}
