package core

import (
	"context"
)

// для всех интерфейсов на уровне пакета: --all
// для конкретного интерфейса: --name=<НазваниеИнтерфейса>

//go:generate go run github.com/vektra/mockery/v2@v2.35.3 --name=ICoreGroup
type ICoreGroup interface {
	Get(context.Context, GroupGetRequest) (*GroupGetResult, error)
	Create(context.Context, GroupCreateRequest) (*GroupCreateResult, error)
}

type (
	Group struct {
		ID     string
		Name   string
		Rights []string
	}

	GroupCreateRequest struct {
		ID     string
		Name   string
		Rights []string
	}

	GroupCreateResult struct {
		Group Group
	}

	GroupGetRequest struct {
		ID string
	}

	GroupGetResult struct {
		Group Group
	}
)
