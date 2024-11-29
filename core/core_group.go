package core

import (
	"context"
	"errors"
)

type ServiceCoreGroup struct {
	// external dependencies
	// ...
	cache map[string]Group
}

func NewCoreGroup(someDependencies ...any) ICoreGroup {
	return &ServiceCoreGroup{
		cache: make(map[string]Group),
	}
}

func (s *ServiceCoreGroup) Get(_ context.Context, req GroupGetRequest) (*GroupGetResult, error) {
	res, ok := s.cache[req.ID]
	if !ok {
		return nil, errors.New("core.group_not_found")
	}

	return &GroupGetResult{
		Group: res,
	}, nil
}

func (s *ServiceCoreGroup) Create(_ context.Context, req GroupCreateRequest) (*GroupCreateResult, error) {
	if !s.validateCreateRequest(req) {
		return nil, errors.New("core.group.invalid_create_request")
	}

	res := Group{
		ID:     req.ID,
		Name:   req.Name,
		Rights: req.Rights,
	}

	s.cache[res.ID] = res

	return &GroupCreateResult{
		Group: res,
	}, nil
}

func (s *ServiceCoreGroup) validateCreateRequest(req GroupCreateRequest) bool {
	if req.Name == "" {
		return false
	}

	if len(req.Rights) == 0 {
		return false
	}

	return true
}
