// Code generated by mockery v2.45.1. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	models "github.com/syedvasil/shaffra_skill_assessment_golang/internal/app/repositories/models"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// CreateUser provides a mock function with given fields: _a0
func (_m *Service) CreateUser(_a0 models.User) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for CreateUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(models.User) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteUser provides a mock function with given fields: id
func (_m *Service) DeleteUser(id primitive.ObjectID) error {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(primitive.ObjectID) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetUserByID provides a mock function with given fields: id
func (_m *Service) GetUserByID(id primitive.ObjectID) (models.User, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetUserByID")
	}

	var r0 models.User
	var r1 error
	if rf, ok := ret.Get(0).(func(primitive.ObjectID) (models.User, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(primitive.ObjectID) models.User); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(models.User)
	}

	if rf, ok := ret.Get(1).(func(primitive.ObjectID) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUsers provides a mock function with given fields: page, limit
func (_m *Service) GetUsers(page int, limit int) ([]models.User, error) {
	ret := _m.Called(page, limit)

	if len(ret) == 0 {
		panic("no return value specified for GetUsers")
	}

	var r0 []models.User
	var r1 error
	if rf, ok := ret.Get(0).(func(int, int) ([]models.User, error)); ok {
		return rf(page, limit)
	}
	if rf, ok := ret.Get(0).(func(int, int) []models.User); ok {
		r0 = rf(page, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.User)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(page, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUser provides a mock function with given fields: id, _a1
func (_m *Service) UpdateUser(id primitive.ObjectID, _a1 models.UserUpdateReq) error {
	ret := _m.Called(id, _a1)

	if len(ret) == 0 {
		panic("no return value specified for UpdateUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(primitive.ObjectID, models.UserUpdateReq) error); ok {
		r0 = rf(id, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewService creates a new instance of Service. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewService(t interface {
	mock.TestingT
	Cleanup(func())
}) *Service {
	mock := &Service{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
