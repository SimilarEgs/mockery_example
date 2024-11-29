package core

import (
	"context"
)

type ICoreUser interface {
	Get(context.Context, UserGetRequest) (*UserGetResult, error)
	Create(context.Context, UserCreateRequest) (*UserCreateResult, error)
}

type (
	User struct {
		ID     string
		Name   string
		Age    int
		Groups []string
	}

	UserCreateRequest struct {
		ID     string
		Name   string
		Age    int
		Groups []string
	}

	UserCreateResult struct {
		User User
	}

	UserGetRequest struct {
		ID string
	}

	UserGetResult struct {
		User User
	}
)
