package v1

type GetUserParams struct {
	Id int64 `from:id`
}

type UpdateUserParams struct {
	Id   int64  `from:id`
	Name string `form:"name"`
	Age  int    `form:"age"`
}
