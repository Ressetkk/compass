// Code generated by mockery v2.5.1. DO NOT EDIT.

package automock

import (
	context "context"

	model "github.com/kyma-incubator/compass/components/director/internal/model"
	mock "github.com/stretchr/testify/mock"
)

// SystemAuthService is an autogenerated mock type for the SystemAuthService type
type SystemAuthService struct {
	mock.Mock
}

// ListForObject provides a mock function with given fields: ctx, objectType, objectID
func (_m *SystemAuthService) ListForObject(ctx context.Context, objectType model.SystemAuthReferenceObjectType, objectID string) ([]model.SystemAuth, error) {
	ret := _m.Called(ctx, objectType, objectID)

	var r0 []model.SystemAuth
	if rf, ok := ret.Get(0).(func(context.Context, model.SystemAuthReferenceObjectType, string) []model.SystemAuth); ok {
		r0 = rf(ctx, objectType, objectID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.SystemAuth)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.SystemAuthReferenceObjectType, string) error); ok {
		r1 = rf(ctx, objectType, objectID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
