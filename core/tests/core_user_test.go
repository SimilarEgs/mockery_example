package tests

import (
	"context"
	"errors"
	"github.com/SimilarEgs/mockery_example/core"
	groupMock "github.com/SimilarEgs/mockery_example/core/mocks"
	loginConverterMock "github.com/SimilarEgs/mockery_example/services/login_converter/mocks"
	notiferMock "github.com/SimilarEgs/mockery_example/services/notifer/mocks"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestServiceCore_Create(t *testing.T) {
	tests := []struct {
		name    string
		start   func() (core.ICoreUser, *core.UserCreateRequest, error)
		wantErr bool
	}{
		{
			name: "core.UserCreate | success | with groups",
			start: func() (core.ICoreUser, *core.UserCreateRequest, error) {
				groupCore := groupMock.NewICoreGroup(t)
				loginConverter := loginConverterMock.NewILoginConverter(t)
				notifier := notiferMock.NewINotifier(t)

				createRequest := &core.UserCreateRequest{
					ID:     "test-user1",
					Name:   "alexey",
					Age:    20,
					Groups: []string{"group1"},
				}

				groupCore.
					On("Get", mock.Anything, core.GroupGetRequest{ID: "group1"}).
					Once().
					Return(&core.GroupGetResult{}, nil)

				loginConverter.
					On("ToFullID", createRequest.ID).
					Twice().
					Return(createRequest.ID)

				notifier.
					On("NotifyOfUserCreated", createRequest.ID).
					Once().
					Return(nil)

				return core.NewCoreUser(groupCore, loginConverter, notifier), createRequest, nil
			},
			wantErr: false,
		},
		{
			name: "core.UserCreate | fail | group get error",
			start: func() (core.ICoreUser, *core.UserCreateRequest, error) {
				groupCore := groupMock.NewICoreGroup(t)
				loginConverter := loginConverterMock.NewILoginConverter(t)
				notifier := notiferMock.NewINotifier(t)

				createRequest := &core.UserCreateRequest{
					ID:     "test-user-fail",
					Name:   "errorUser",
					Age:    30,
					Groups: []string{"group2"},
				}

				groupCore.
					On("Get", mock.Anything, core.GroupGetRequest{ID: "group2"}).
					Once().
					Return(nil, errors.New("core.group_not_found"))

				return core.NewCoreUser(groupCore, loginConverter, notifier), createRequest, nil
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, request, err := tt.start()
			if err != nil {
				t.Error(err)
				return
			}

			_, err = s.Create(context.TODO(), *request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
