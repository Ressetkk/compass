// Code generated by mockery v2.9.4. DO NOT EDIT.

package automock

import (
	context "context"

	model "github.com/kyma-incubator/compass/components/director/internal/model"
	mock "github.com/stretchr/testify/mock"
)

// ApplicationRepository is an autogenerated mock type for the ApplicationRepository type
type ApplicationRepository struct {
	mock.Mock
}

// GetGlobalByID provides a mock function with given fields: ctx, id
func (_m *ApplicationRepository) GetGlobalByID(ctx context.Context, id string) (*model.Application, error) {
	ret := _m.Called(ctx, id)

	var r0 *model.Application
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.Application); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Application)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
