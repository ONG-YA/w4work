package biz

import "context"

type User struct {
	Id   int64
	Name string
	Age  int
}

type UserRepo interface {
	GetUser(ctx context.Context, id int64) (*User, error)
	UpdateUser(ctx context.Context, id int64, user *User) error
}

type UserUsecase struct {
	repo UserRepo
}

func NewUserUsecase(repo UserRepo) *UserUsecase {
	return &UserUsecase{repo: repo}
}

func (us *UserUsecase) Get(ctx context.Context, id int64) (*User, error) {
	return us.repo.GetUser(ctx, id)
}

func (us *UserUsecase) Update(ctx context.Context, id int64, user *User) error {
	return us.repo.UpdateUser(ctx, id, user)
}
