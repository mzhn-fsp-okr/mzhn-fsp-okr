// Code generated by mockery v2.46.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// SessionsStorage is an autogenerated mock type for the SessionsStorage type
type SessionsStorage struct {
	mock.Mock
}

// Check provides a mock function with given fields: ctx, userId, token
func (_m *SessionsStorage) Check(ctx context.Context, userId string, token string) error {
	ret := _m.Called(ctx, userId, token)

	if len(ret) == 0 {
		panic("no return value specified for Check")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, userId, token)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: ctx, userId
func (_m *SessionsStorage) Delete(ctx context.Context, userId string) error {
	ret := _m.Called(ctx, userId)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, userId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Save provides a mock function with given fields: ctx, userId, token
func (_m *SessionsStorage) Save(ctx context.Context, userId string, token string) error {
	ret := _m.Called(ctx, userId, token)

	if len(ret) == 0 {
		panic("no return value specified for Save")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, userId, token)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewSessionsStorage creates a new instance of SessionsStorage. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSessionsStorage(t interface {
	mock.TestingT
	Cleanup(func())
}) *SessionsStorage {
	mock := &SessionsStorage{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
