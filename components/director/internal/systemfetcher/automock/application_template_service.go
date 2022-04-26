// Code generated by mockery v2.10.5. DO NOT EDIT.

package automock

import (
	context "context"

	model "github.com/kyma-incubator/compass/components/director/internal/model"
	mock "github.com/stretchr/testify/mock"
)

// ApplicationTemplateService is an autogenerated mock type for the applicationTemplateService type
type ApplicationTemplateService struct {
	mock.Mock
}

// Get provides a mock function with given fields: ctx, id
func (_m *ApplicationTemplateService) Get(ctx context.Context, id string) (*model.ApplicationTemplate, error) {
	ret := _m.Called(ctx, id)

	var r0 *model.ApplicationTemplate
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.ApplicationTemplate); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.ApplicationTemplate)
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

// PrepareApplicationCreateInputJSON provides a mock function with given fields: appTemplate, values
func (_m *ApplicationTemplateService) PrepareApplicationCreateInputJSON(appTemplate *model.ApplicationTemplate, values model.ApplicationFromTemplateInputValues) (string, error) {
	ret := _m.Called(appTemplate, values)

	var r0 string
	if rf, ok := ret.Get(0).(func(*model.ApplicationTemplate, model.ApplicationFromTemplateInputValues) string); ok {
		r0 = rf(appTemplate, values)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.ApplicationTemplate, model.ApplicationFromTemplateInputValues) error); ok {
		r1 = rf(appTemplate, values)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
