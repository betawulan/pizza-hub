// Code generated by mockery v2.34.1. DO NOT EDIT.

package mocks

import (
	context "context"

	model "github.com/betawulan/pizza-hub/model"
	mock "github.com/stretchr/testify/mock"
)

// MenuService is an autogenerated mock type for the MenuService type
type MenuService struct {
	mock.Mock
}

// GetMenus provides a mock function with given fields: ctx
func (_m *MenuService) GetMenus(ctx context.Context) ([]model.Menu, error) {
	ret := _m.Called(ctx)

	var r0 []model.Menu
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]model.Menu, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []model.Menu); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Menu)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMenuService creates a new instance of MenuService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMenuService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MenuService {
	mock := &MenuService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
