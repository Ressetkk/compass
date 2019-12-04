// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	gqlschema "github.com/kyma-incubator/compass/components/provisioner/pkg/gqlschema"
	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// CleanupRuntimeData provides a mock function with given fields: id
func (_m *Service) CleanupRuntimeData(id string) (*gqlschema.CleanUpRuntimeDataResult, error) {
	ret := _m.Called(id)

	var r0 *gqlschema.CleanUpRuntimeDataResult
	if rf, ok := ret.Get(0).(func(string) *gqlschema.CleanUpRuntimeDataResult); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gqlschema.CleanUpRuntimeDataResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeprovisionRuntime provides a mock function with given fields: id
func (_m *Service) DeprovisionRuntime(id string) (string, <-chan struct{}, error) {
	ret := _m.Called(id)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 <-chan struct{}
	if rf, ok := ret.Get(1).(func(string) <-chan struct{}); ok {
		r1 = rf(id)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(<-chan struct{})
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string) error); ok {
		r2 = rf(id)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ProvisionRuntime provides a mock function with given fields: id, config
func (_m *Service) ProvisionRuntime(id string, config gqlschema.ProvisionRuntimeInput) (string, <-chan struct{}, error) {
	ret := _m.Called(id, config)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, gqlschema.ProvisionRuntimeInput) string); ok {
		r0 = rf(id, config)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 <-chan struct{}
	if rf, ok := ret.Get(1).(func(string, gqlschema.ProvisionRuntimeInput) <-chan struct{}); ok {
		r1 = rf(id, config)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(<-chan struct{})
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, gqlschema.ProvisionRuntimeInput) error); ok {
		r2 = rf(id, config)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ReconnectRuntimeAgent provides a mock function with given fields: id
func (_m *Service) ReconnectRuntimeAgent(id string) (string, error) {
	ret := _m.Called(id)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RuntimeOperationStatus provides a mock function with given fields: id
func (_m *Service) RuntimeOperationStatus(id string) (*gqlschema.OperationStatus, error) {
	ret := _m.Called(id)

	var r0 *gqlschema.OperationStatus
	if rf, ok := ret.Get(0).(func(string) *gqlschema.OperationStatus); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gqlschema.OperationStatus)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RuntimeStatus provides a mock function with given fields: id
func (_m *Service) RuntimeStatus(id string) (*gqlschema.RuntimeStatus, error) {
	ret := _m.Called(id)

	var r0 *gqlschema.RuntimeStatus
	if rf, ok := ret.Get(0).(func(string) *gqlschema.RuntimeStatus); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gqlschema.RuntimeStatus)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpgradeRuntime provides a mock function with given fields: id, config
func (_m *Service) UpgradeRuntime(id string, config *gqlschema.UpgradeRuntimeInput) (string, error) {
	ret := _m.Called(id, config)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, *gqlschema.UpgradeRuntimeInput) string); ok {
		r0 = rf(id, config)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, *gqlschema.UpgradeRuntimeInput) error); ok {
		r1 = rf(id, config)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
