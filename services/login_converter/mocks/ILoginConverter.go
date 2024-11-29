// Code generated by mockery v2.35.3. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// ILoginConverter is an autogenerated mock type for the ILoginConverter type
type ILoginConverter struct {
	mock.Mock
}

// IsFullID provides a mock function with given fields: id
func (_m *ILoginConverter) IsFullID(id string) bool {
	ret := _m.Called(id)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// ToFullID provides a mock function with given fields: id
func (_m *ILoginConverter) ToFullID(id string) string {
	ret := _m.Called(id)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// NewILoginConverter creates a new instance of ILoginConverter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewILoginConverter(t interface {
	mock.TestingT
	Cleanup(func())
}) *ILoginConverter {
	mock := &ILoginConverter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
