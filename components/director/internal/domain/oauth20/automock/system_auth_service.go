// Code generated by mockery v2.10.5. DO NOT EDIT.

package automock

import (
	context "context"

	internalmodel "github.com/kyma-incubator/compass/components/director/internal/model"
	mock "github.com/stretchr/testify/mock"

	model "github.com/kyma-incubator/compass/components/director/pkg/model"
)

// SystemAuthService is an autogenerated mock type for the SystemAuthService type
type SystemAuthService struct {
	mock.Mock
}

// CreateWithCustomID provides a mock function with given fields: ctx, id, objectType, objectID, authInput
func (_m *SystemAuthService) CreateWithCustomID(ctx context.Context, id string, objectType model.SystemAuthReferenceObjectType, objectID string, authInput *internalmodel.AuthInput) (string, error) {
	ret := _m.Called(ctx, id, objectType, objectID, authInput)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, string, model.SystemAuthReferenceObjectType, string, *internalmodel.AuthInput) string); ok {
		r0 = rf(ctx, id, objectType, objectID, authInput)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, model.SystemAuthReferenceObjectType, string, *internalmodel.AuthInput) error); ok {
		r1 = rf(ctx, id, objectType, objectID, authInput)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByIDForObject provides a mock function with given fields: ctx, objectType, authID
func (_m *SystemAuthService) GetByIDForObject(ctx context.Context, objectType model.SystemAuthReferenceObjectType, authID string) (*model.SystemAuth, error) {
	ret := _m.Called(ctx, objectType, authID)

	var r0 *model.SystemAuth
	if rf, ok := ret.Get(0).(func(context.Context, model.SystemAuthReferenceObjectType, string) *model.SystemAuth); ok {
		r0 = rf(ctx, objectType, authID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.SystemAuth)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.SystemAuthReferenceObjectType, string) error); ok {
		r1 = rf(ctx, objectType, authID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
