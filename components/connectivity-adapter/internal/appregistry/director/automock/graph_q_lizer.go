// Code generated by mockery v2.5.1. DO NOT EDIT.

package automock

import (
	graphql "github.com/kyma-incubator/compass/components/director/pkg/graphql"
	mock "github.com/stretchr/testify/mock"
)

// GraphQLizer is an autogenerated mock type for the GraphQLizer type
type GraphQLizer struct {
	mock.Mock
}

// APIDefinitionInputToGQL provides a mock function with given fields: in
func (_m *GraphQLizer) APIDefinitionInputToGQL(in graphql.APIDefinitionInput) (string, error) {
	ret := _m.Called(in)

	var r0 string
	if rf, ok := ret.Get(0).(func(graphql.APIDefinitionInput) string); ok {
		r0 = rf(in)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(graphql.APIDefinitionInput) error); ok {
		r1 = rf(in)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BundleCreateInputToGQL provides a mock function with given fields: in
func (_m *GraphQLizer) BundleCreateInputToGQL(in graphql.BundleCreateInput) (string, error) {
	ret := _m.Called(in)

	var r0 string
	if rf, ok := ret.Get(0).(func(graphql.BundleCreateInput) string); ok {
		r0 = rf(in)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(graphql.BundleCreateInput) error); ok {
		r1 = rf(in)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BundleUpdateInputToGQL provides a mock function with given fields: in
func (_m *GraphQLizer) BundleUpdateInputToGQL(in graphql.BundleUpdateInput) (string, error) {
	ret := _m.Called(in)

	var r0 string
	if rf, ok := ret.Get(0).(func(graphql.BundleUpdateInput) string); ok {
		r0 = rf(in)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(graphql.BundleUpdateInput) error); ok {
		r1 = rf(in)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DocumentInputToGQL provides a mock function with given fields: in
func (_m *GraphQLizer) DocumentInputToGQL(in *graphql.DocumentInput) (string, error) {
	ret := _m.Called(in)

	var r0 string
	if rf, ok := ret.Get(0).(func(*graphql.DocumentInput) string); ok {
		r0 = rf(in)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*graphql.DocumentInput) error); ok {
		r1 = rf(in)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EventDefinitionInputToGQL provides a mock function with given fields: in
func (_m *GraphQLizer) EventDefinitionInputToGQL(in graphql.EventDefinitionInput) (string, error) {
	ret := _m.Called(in)

	var r0 string
	if rf, ok := ret.Get(0).(func(graphql.EventDefinitionInput) string); ok {
		r0 = rf(in)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(graphql.EventDefinitionInput) error); ok {
		r1 = rf(in)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
