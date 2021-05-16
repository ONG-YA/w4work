package v1

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserServiceHandler interface {
	UpdateUser(context.Context, *UpdateUserParams) error
	GetUser(context.Context, *GetUserParams) (interface{}, error)
}

func NewUserServiceHandler(srv UserServiceHandler) http.Handler {
	router := gin.Default()
	router.GET("/user/query", func(ctx *gin.Context) {
		var userQuery GetUserParams
		var result interface{}
		if err := ctx.ShouldBind(&userQuery); err == nil {
			if result, err = srv.GetUser(ctx, &userQuery); err != nil {
				ctx.JSON(500, gin.H{"msg": err.Error()})
				return
			}
			_, errr := json.Marshal(result)
			if errr != nil {
				ctx.JSON(500, gin.H{"msg": err.Error()})
				return
			}
			ctx.JSON(200, result)
		}
	})

	router.GET("/user/update", func(ctx *gin.Context) {
		var userUpdate UpdateUserParams
		if err := ctx.ShouldBind(&userUpdate); err == nil {
			if err = srv.UpdateUser(ctx, &userUpdate); err != nil {
				ctx.JSON(500, gin.H{"msg": err.Error()})
				return
			}
			ctx.JSON(200, gin.H{"msg": "update success"})
		} else {
			ctx.JSON(500, gin.H{"msg": err.Error()})
		}
	})
	return router
}
