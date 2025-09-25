package validator

import (
	"github.com/go-playground/validator/v10"
)
var _validate *validator.Validate

func init(){
	_validate = validator.New()
}