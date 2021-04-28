package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"path/filepath"
	"strings"
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
				"error":        "Invalid request parameters. See invalid_args",
				"invalid_args": invalidArgs}, false
		}
	}
	return obj, true
}

func UploadHandler(ctx *gin.Context, field string) string {
	file, err := ctx.FormFile(field)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	extension := filepath.Ext(file.Filename)
	newFileName := GetUuid() + extension
	err = ctx.SaveUploadedFile(file, GetMediaRoot()+newFileName)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return newFileName
}

func GetUuid() string {
	uuidWithHyphen := uuid.New()
	newUuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	return newUuid
}

func GetMediaRoot() string {
	mediaRoot := "/go/src/app"
	return fmt.Sprintf("%v/%v/", mediaRoot, "media")
}
