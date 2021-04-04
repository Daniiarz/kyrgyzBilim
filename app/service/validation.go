package service

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type invalidArgument struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
}

func DataBind(ctx *gin.Context, obj interface{}) (interface{}, bool) {
	if err := ctx.ShouldBind(obj); err != nil {
		var invalidArgs []invalidArgument
		if errs, ok := err.(validator.ValidationErrors); ok {
			for _, err := range errs {
				invalidArgs = append(invalidArgs, invalidArgument{
					Field: err.Field(),
					Tag:   err.Tag(),
				})
			}
			return gin.H{
				"error": "Invalid request parameters. See invalid_args",
				"invalid_args": invalidArgs}, false
		}
	}
	return obj, true
}
