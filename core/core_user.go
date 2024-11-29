package core

import (
	"context"
	"errors"
	"github.com/SimilarEgs/mockery_example/services/login_converter"
	"github.com/SimilarEgs/mockery_example/services/notifer"
)

type ServiceCoreUser struct {
	cache map[string]User

	Group          ICoreGroup
	LoginConverter login_converter.ILoginConverter
	Notifier       notifer.INotifier
}

func NewCoreUser(coreGroup ICoreGroup, lc login_converter.ILoginConverter, notifier notifer.INotifier) ICoreUser {
	return &ServiceCoreUser{
		cache:          make(map[string]User),
		Group:          coreGroup,
		LoginConverter: lc,
		Notifier:       notifier,
	}
}

func (s *ServiceCoreUser) Get(_ context.Context, req UserGetRequest) (*UserGetResult, error) {
	res, ok := s.cache[s.LoginConverter.ToFullID(req.ID)]
	if !ok {
		return nil, errors.New("core.user_not_found")
	}

	return &UserGetResult{
		User: res,
	}, nil
}

func (s *ServiceCoreUser) Create(_ context.Context, req UserCreateRequest) (*UserCreateResult, error) {
	if !s.validateCreateRequest(req) {
		return nil, errors.New("core.user.invalid_create_request")
	}

	usr := User{
		ID:     s.LoginConverter.ToFullID(req.ID),
		Name:   req.Name,
		Age:    req.Age,
		Groups: req.Groups,
	}

	s.cache[usr.ID] = usr

	if err := s.Notifier.NotifyOfUserCreated(usr.ID); err != nil {
		return nil, err
	}

	res, err := s.Get(context.TODO(), UserGetRequest{ID: req.ID})
	if err != nil {
		return nil, err
	}

	return &UserCreateResult{User: res.User}, nil
}

func (s *ServiceCoreUser) validateCreateRequest(req UserCreateRequest) bool {
	for _, grp := range req.Groups {
		if _, err := s.Group.Get(context.TODO(), GroupGetRequest{ID: grp}); err != nil {
			return false
		}
	}

	return true
}
