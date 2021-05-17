package v1

import (
	"context"
	"errors"
	"net/http"
	"w4work/configs"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type UserServiceHandler interface {
	UpdateUser(context.Context, *UpdateUserParams) error
	GetUser(context.Context, *GetUserParams) (interface{}, error)
}

func NewUserServiceHandler(srv UserServiceHandler) http.Handler {
	router := gin.Default()
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("useridValid", useridValid)
		v.RegisterValidation("ageValid", ageValid)
	}

	router.GET("/user/query", func(ctx *gin.Context) {
		var userQuery GetUserParams
		var result interface{}
		if err := ctx.ShouldBindWith(&userQuery,binding.Query); err == nil {
			if result, err = srv.GetUser(ctx, &userQuery); err != nil {
				if errors.Is(err,configs.ErrNotFound) {
					ctx.JSON(http.StatusOK, gin.H{"msg": "no data found!"})
					return
				}
				ctx.JSON(http.StatusInternalServerError, gin.H{"msg": configs.ErrInteralFound.Error()})
				return
			}
			ctx.JSON(http.StatusOK, result)
		}else {
			ctx.JSON(http.StatusBadRequest, gin.H{"msg": configs.ErrParamsInvalid.Error()})
		}
	})

	router.GET("/user/update", func(ctx *gin.Context) {
		var userUpdate UpdateUserParams
		if err := ctx.ShouldBindWith(&userUpdate,binding.Query); err == nil {
			if err = srv.UpdateUser(ctx, &userUpdate); err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"msg": configs.ErrInteralFound.Error()})
				return
			}
			ctx.JSON(http.StatusOK, gin.H{"msg": "update success"})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{"msg": configs.ErrParamsInvalid.Error()})
		}
	})
	return router
}
