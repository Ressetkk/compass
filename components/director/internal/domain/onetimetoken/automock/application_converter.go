// Code generated by mockery v2.10.5. DO NOT EDIT.

package automock

import (
	graphql "github.com/kyma-incubator/compass/components/director/pkg/graphql"
	mock "github.com/stretchr/testify/mock"

	model "github.com/kyma-incubator/compass/components/director/internal/model"
)

// ApplicationConverter is an autogenerated mock type for the ApplicationConverter type
type ApplicationConverter struct {
	mock.Mock
}

// ToGraphQL provides a mock function with given fields: in
func (_m *ApplicationConverter) ToGraphQL(in *model.Application) *graphql.Application {
	ret := _m.Called(in)

	var r0 *graphql.Application
	if rf, ok := ret.Get(0).(func(*model.Application) *graphql.Application); ok {
		r0 = rf(in)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*graphql.Application)
		}
	}

	return r0
}
