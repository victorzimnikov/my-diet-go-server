package structures

type errorData struct {
	message string
}

type AppError interface {
	Error() errorData
}

type appError struct {
	message string
}

func (c appError) Error() errorData {
	res := errorData{}

	res.message = c.message

	return res
}

func MakeAppError(message string) AppError {
	var ts AppError = appError{message: message}

	return ts
}
