// Code generated by mockery v2.34.1. DO NOT EDIT.

package mocks

import (
	context "context"

	model "github.com/betawulan/pizza-hub/model"
	mock "github.com/stretchr/testify/mock"
)

// MenuRepository is an autogenerated mock type for the MenuRepository type
type MenuRepository struct {
	mock.Mock
}

// GetMenus provides a mock function with given fields: ctx
func (_m *MenuRepository) GetMenus(ctx context.Context) ([]model.Menu, error) {
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

// NewMenuRepository creates a new instance of MenuRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMenuRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MenuRepository {
	mock := &MenuRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
