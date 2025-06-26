package v1

import "github.com/ividernvi/ivicompose/pkg/options"

type UserLogin struct {
	UserUUID string `json:"user_uuid" validate:"uuid;required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"email;required"`
	Phone    string `json:"phone" validate:"phone"`
}

type UserList struct {
	options.ListMeta
	Users []UserLogin `json:"users"`
}
