// Code generated by mockery v2.10.5. DO NOT EDIT.

package automock

import (
	graphql "github.com/kyma-incubator/compass/components/director/pkg/graphql"
	mock "github.com/stretchr/testify/mock"

	model "github.com/kyma-incubator/compass/components/director/internal/model"
)

// RuntimeContextConverter is an autogenerated mock type for the RuntimeContextConverter type
type RuntimeContextConverter struct {
	mock.Mock
}

// InputFromGraphQL provides a mock function with given fields: in, runtimeID
func (_m *RuntimeContextConverter) InputFromGraphQL(in graphql.RuntimeContextInput, runtimeID string) model.RuntimeContextInput {
	ret := _m.Called(in, runtimeID)

	var r0 model.RuntimeContextInput
	if rf, ok := ret.Get(0).(func(graphql.RuntimeContextInput, string) model.RuntimeContextInput); ok {
		r0 = rf(in, runtimeID)
	} else {
		r0 = ret.Get(0).(model.RuntimeContextInput)
	}

	return r0
}

// MultipleToGraphQL provides a mock function with given fields: in
func (_m *RuntimeContextConverter) MultipleToGraphQL(in []*model.RuntimeContext) []*graphql.RuntimeContext {
	ret := _m.Called(in)

	var r0 []*graphql.RuntimeContext
	if rf, ok := ret.Get(0).(func([]*model.RuntimeContext) []*graphql.RuntimeContext); ok {
		r0 = rf(in)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*graphql.RuntimeContext)
		}
	}

	return r0
}

// ToGraphQL provides a mock function with given fields: in
func (_m *RuntimeContextConverter) ToGraphQL(in *model.RuntimeContext) *graphql.RuntimeContext {
	ret := _m.Called(in)

	var r0 *graphql.RuntimeContext
	if rf, ok := ret.Get(0).(func(*model.RuntimeContext) *graphql.RuntimeContext); ok {
		r0 = rf(in)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*graphql.RuntimeContext)
		}
	}

	return r0
}
