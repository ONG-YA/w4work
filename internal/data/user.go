package data

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"time"
	"w4work/configs"
	"w4work/internal/biz"
)

type userRepo struct {
	data *Data
}

func NewUserRepo(data *Data) biz.UserRepo {
	return &userRepo{data: data}
}

func (us *userRepo) GetUser(ctx context.Context, id int64) (*biz.User, error) {
	key := fmt.Sprintf("user_%d", id)
	val, err := us.data.rd.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil,  errors.Wrapf(configs.ErrNotFound, fmt.Sprintf("key: %s error: %v", key, err))
	} else if err != nil {
		return nil, errors.Wrapf(configs.ErrInteralFound, fmt.Sprintf("key: %s error: %v", key, err))
	}
	user := &biz.User{}
	errr := json.Unmarshal([]byte(val), user)
	return user, errr
}

func (us *userRepo) UpdateUser(ctx context.Context, id int64, user *biz.User) error {
	key := fmt.Sprintf("user_%d", id)
	v, err := json.Marshal(&user)
	if err != nil {
		return err
	}
	return us.data.rd.Set(ctx, key, string(v), 24*time.Hour).Err()
}
