// Code generated by mockery v1.0.0. DO NOT EDIT.

package automock

import mock "github.com/stretchr/testify/mock"

// TokenGenerator is an autogenerated mock type for the TokenGenerator type
type TokenGenerator struct {
	mock.Mock
}

// NewToken provides a mock function with given fields:
func (_m *TokenGenerator) NewToken() (string, error) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
