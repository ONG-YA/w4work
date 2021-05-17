package v1

import (
	"github.com/go-playground/validator/v10"
)

type GetUserParams struct {
	Id int64 `form:"id" binding:"required,useridValid"`
}

type UpdateUserParams struct {
	Id   int64 `form:"id" binding:"required,useridValid"`
	Name string `form:"name"`
	Age  int `form:"age" binding:"required,useridValid"`
}

var useridValid validator.Func = func(fl validator.FieldLevel) bool {
	id, ok := fl.Field().Interface().(int64)
	if ok && id >= 0{
		return true
	}
	return false
}

var ageValid validator.Func = func(fl validator.FieldLevel) bool {
	age, ok := fl.Field().Interface().(int)
	if ok && age >= 0{
		return true
	}
	return false
}
