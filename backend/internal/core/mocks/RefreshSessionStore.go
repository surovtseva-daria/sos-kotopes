// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	core "github.com/kotopesp/sos-kotopes/internal/core"
	mock "github.com/stretchr/testify/mock"
)

// RefreshSessionStore is an autogenerated mock type for the RefreshSessionStore type
type RefreshSessionStore struct {
	mock.Mock
}

// CountSessionsAndDelete provides a mock function with given fields: ctx, userID
func (_m *RefreshSessionStore) CountSessionsAndDelete(ctx context.Context, userID int) error {
	ret := _m.Called(ctx, userID)

	if len(ret) == 0 {
		panic("no return value specified for CountSessionsAndDelete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int) error); ok {
		r0 = rf(ctx, userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetRefreshSessionByToken provides a mock function with given fields: ctx, token
func (_m *RefreshSessionStore) GetRefreshSessionByToken(ctx context.Context, token string) (core.RefreshSession, error) {
	ret := _m.Called(ctx, token)

	if len(ret) == 0 {
		panic("no return value specified for GetRefreshSessionByToken")
	}

	var r0 core.RefreshSession
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (core.RefreshSession, error)); ok {
		return rf(ctx, token)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) core.RefreshSession); ok {
		r0 = rf(ctx, token)
	} else {
		r0 = ret.Get(0).(core.RefreshSession)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateRefreshSession provides a mock function with given fields: ctx, param, refreshSession
func (_m *RefreshSessionStore) UpdateRefreshSession(ctx context.Context, param core.UpdateRefreshSessionParam, refreshSession core.RefreshSession) error {
	ret := _m.Called(ctx, param, refreshSession)

	if len(ret) == 0 {
		panic("no return value specified for UpdateRefreshSession")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, core.UpdateRefreshSessionParam, core.RefreshSession) error); ok {
		r0 = rf(ctx, param, refreshSession)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewRefreshSessionStore creates a new instance of RefreshSessionStore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRefreshSessionStore(t interface {
	mock.TestingT
	Cleanup(func())
}) *RefreshSessionStore {
	mock := &RefreshSessionStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
