// Code generated by mockery v2.43.2. DO NOT EDIT.

package core

import (
	context "context"

	core "github.com/kotopesp/sos-kotopes/internal/core"
	mock "github.com/stretchr/testify/mock"
)

// MockUserService is an autogenerated mock type for the UserService type
type MockUserService struct {
	mock.Mock
}

type MockUserService_Expecter struct {
	mock *mock.Mock
}

func (_m *MockUserService) EXPECT() *MockUserService_Expecter {
	return &MockUserService_Expecter{mock: &_m.Mock}
}

// GetUser provides a mock function with given fields: ctx, id
func (_m *MockUserService) GetUser(ctx context.Context, id int) (core.User, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetUser")
	}

	var r0 core.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) (core.User, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) core.User); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(core.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUserService_GetUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUser'
type MockUserService_GetUser_Call struct {
	*mock.Call
}

// GetUser is a helper method to define mock.On call
//   - ctx context.Context
//   - id int
func (_e *MockUserService_Expecter) GetUser(ctx interface{}, id interface{}) *MockUserService_GetUser_Call {
	return &MockUserService_GetUser_Call{Call: _e.mock.On("GetUser", ctx, id)}
}

func (_c *MockUserService_GetUser_Call) Run(run func(ctx context.Context, id int)) *MockUserService_GetUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int))
	})
	return _c
}

func (_c *MockUserService_GetUser_Call) Return(user core.User, err error) *MockUserService_GetUser_Call {
	_c.Call.Return(user, err)
	return _c
}

func (_c *MockUserService_GetUser_Call) RunAndReturn(run func(context.Context, int) (core.User, error)) *MockUserService_GetUser_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateUser provides a mock function with given fields: ctx, id, update
func (_m *MockUserService) UpdateUser(ctx context.Context, id int, update core.UpdateUser) (core.User, error) {
	ret := _m.Called(ctx, id, update)

	if len(ret) == 0 {
		panic("no return value specified for UpdateUser")
	}

	var r0 core.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int, core.UpdateUser) (core.User, error)); ok {
		return rf(ctx, id, update)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int, core.UpdateUser) core.User); ok {
		r0 = rf(ctx, id, update)
	} else {
		r0 = ret.Get(0).(core.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int, core.UpdateUser) error); ok {
		r1 = rf(ctx, id, update)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUserService_UpdateUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateUser'
type MockUserService_UpdateUser_Call struct {
	*mock.Call
}

// UpdateUser is a helper method to define mock.On call
//   - ctx context.Context
//   - id int
//   - update core.UpdateUser
func (_e *MockUserService_Expecter) UpdateUser(ctx interface{}, id interface{}, update interface{}) *MockUserService_UpdateUser_Call {
	return &MockUserService_UpdateUser_Call{Call: _e.mock.On("UpdateUser", ctx, id, update)}
}

func (_c *MockUserService_UpdateUser_Call) Run(run func(ctx context.Context, id int, update core.UpdateUser)) *MockUserService_UpdateUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int), args[2].(core.UpdateUser))
	})
	return _c
}

func (_c *MockUserService_UpdateUser_Call) Return(updatedUser core.User, err error) *MockUserService_UpdateUser_Call {
	_c.Call.Return(updatedUser, err)
	return _c
}

func (_c *MockUserService_UpdateUser_Call) RunAndReturn(run func(context.Context, int, core.UpdateUser) (core.User, error)) *MockUserService_UpdateUser_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockUserService creates a new instance of MockUserService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockUserService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockUserService {
	mock := &MockUserService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
