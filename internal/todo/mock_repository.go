// Code generated by mockery v2.3.0. DO NOT EDIT.

package todo

import mock "github.com/stretchr/testify/mock"

// MockRepository is an autogenerated mock type for the Repository type
type MockRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: todo
func (_m *MockRepository) Create(todo Todo) (Todo, error) {
	ret := _m.Called(todo)

	var r0 Todo
	if rf, ok := ret.Get(0).(func(Todo) Todo); ok {
		r0 = rf(todo)
	} else {
		r0 = ret.Get(0).(Todo)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(Todo) error); ok {
		r1 = rf(todo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAll provides a mock function with given fields:
func (_m *MockRepository) FindAll() ([]Todo, error) {
	ret := _m.Called()

	var r0 []Todo
	if rf, ok := ret.Get(0).(func() []Todo); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]Todo)
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
