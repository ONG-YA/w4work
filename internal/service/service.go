package service

import (
	"github.com/google/wire"
	"w4work/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewUserService)

type UserService struct {
	user *biz.UserUsecase
}
