// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	"acp-iam-api/business/roles"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

func (_m *Repository) GetRoles(id uint) (*roles.Roles, error) {
	ret := _m.Called(id)

	var r0 *roles.Roles
	if rf, ok := ret.Get(0).(func(uint) *roles.Roles); ok {
		r0 = rf(uint(id))
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*roles.Roles)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(uint(id))
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *Repository) GetAllRoles() ([]roles.Roles, error) {
	ret := _m.Called()

	var r0 []roles.Roles
	if rf, ok := ret.Get(0).(func() []roles.Roles); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]roles.Roles)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *Repository) AddRoles(name string, isActive bool) error {
	ret := _m.Called(name, isActive)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, bool) error); ok {
		r0 = rf(name, isActive)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *Repository) UpdateRoles(id uint, rolesParam roles.Roles) error {
	ret := _m.Called(id, rolesParam)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, roles.Roles) error); ok {
		r0 = rf(id, rolesParam)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *Repository) DeleteRoles(id uint) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
