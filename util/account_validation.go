package util

import (
	"fiber-ordering/exception"
	"fiber-ordering/model"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"regexp"
)
func Validate(request *model.AccountRequest) {
	err := validation.ValidateStruct(request,
		validation.Field(&request.FullName, validation.Required),
		validation.Field(&request.Email, validation.Required),
		validation.Field(&request.PhoneNumber, validation.Required, validation.Match(regexp.MustCompile("^[0-9]{12}$"))),
		validation.Field(&request.Username, validation.Required),
		validation.Field(&request.Password, validation.Required, validation.Length(6,100)),
	)

	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}
}