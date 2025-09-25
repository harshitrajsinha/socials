package validator

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func Users( s interface{}) error {
	err := _validate.Struct(s)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); !ok {
			panic(err)
		}
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.Field())
			fmt.Println(err.Tag())

		}
		return err
	}
	return nil
}
func Payload(s interface{}) error {
	err := _validate.Struct(s)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); !ok {
			panic(err)
		}
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.Field())
			fmt.Println(err.Tag())
		}
		return err
	}
	return nil
}