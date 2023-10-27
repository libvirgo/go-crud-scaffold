package validator

import (
	"github.com/go-playground/validator/v10"
	"frame/utils"
)

var WalletValidator validator.Func = func(fl validator.FieldLevel) bool {
	s, ok := fl.Field().Interface().(string)
	if ok {
		if utils.String(s).IsWalletAddress() {
			return true
		}
	}
	return false
}
