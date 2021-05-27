package exception

type ValidationError struct {
	Message string
}

func (validatorError ValidationError) Error() string {
	return validatorError.Message
}