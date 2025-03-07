package vaildator

import (
	"doovvvblog/utils/errmsg"

	"github.com/go-playground/validator/v10"
)

func Vaildator(data interface{}) (string, int) {
	var validate = validator.New()
	err := validate.Struct(data)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			return e.Error(), errmsg.ERROR
		}
	}
	return "", errmsg.SUCCESS
}
