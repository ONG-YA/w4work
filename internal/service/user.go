package service

import (
	"context"
	"errors"
	"w4work/api/user/v1"
	"w4work/internal/biz"
)

var ErrIdInvalid = errors.New("id is invalid")

func NewUserService(user *biz.UserUsecase) *UserService {
	return &UserService{user: user}
}

func (u *UserService) GetUser(ctx context.Context, p *v1.GetUserParams) (interface{}, error) {
	if p.Id < 0 {
		return nil, ErrIdInvalid
	}
	return u.user.Get(ctx, p.Id)
}

func (u *UserService) UpdateUser(ctx context.Context, p *v1.UpdateUserParams) error {
	if p.Id < 0 {
		return ErrIdInvalid
	}
	user := &biz.User{Id: p.Id, Name: p.Name, Age: p.Age}
	return u.user.Update(ctx, p.Id, user)
}
