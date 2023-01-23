package api

import (
	"github.com/go-playground/validator"
	"github.com/renaldyhidayatt/ImplSimpleBank/util"
)

var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		return util.IsSupportedCurrency(currency)
	}
	return false
}
