// Code generated by mockery v2.5.1. DO NOT EDIT.

package automock

import (
	graphql "github.com/kyma-incubator/compass/components/director/pkg/graphql"
	mock "github.com/stretchr/testify/mock"

	model "github.com/kyma-incubator/compass/components/director/internal/model"
)

// Converter is an autogenerated mock type for the Converter type
type Converter struct {
	mock.Mock
}

// FromInputGraphQL provides a mock function with given fields: in
func (_m *Converter) FromInputGraphQL(in graphql.AutomaticScenarioAssignmentSetInput) model.AutomaticScenarioAssignment {
	ret := _m.Called(in)

	var r0 model.AutomaticScenarioAssignment
	if rf, ok := ret.Get(0).(func(graphql.AutomaticScenarioAssignmentSetInput) model.AutomaticScenarioAssignment); ok {
		r0 = rf(in)
	} else {
		r0 = ret.Get(0).(model.AutomaticScenarioAssignment)
	}

	return r0
}

// LabelSelectorFromInput provides a mock function with given fields: in
func (_m *Converter) LabelSelectorFromInput(in graphql.LabelSelectorInput) model.LabelSelector {
	ret := _m.Called(in)

	var r0 model.LabelSelector
	if rf, ok := ret.Get(0).(func(graphql.LabelSelectorInput) model.LabelSelector); ok {
		r0 = rf(in)
	} else {
		r0 = ret.Get(0).(model.LabelSelector)
	}

	return r0
}

// MultipleToGraphQL provides a mock function with given fields: assignments
func (_m *Converter) MultipleToGraphQL(assignments []*model.AutomaticScenarioAssignment) []*graphql.AutomaticScenarioAssignment {
	ret := _m.Called(assignments)

	var r0 []*graphql.AutomaticScenarioAssignment
	if rf, ok := ret.Get(0).(func([]*model.AutomaticScenarioAssignment) []*graphql.AutomaticScenarioAssignment); ok {
		r0 = rf(assignments)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*graphql.AutomaticScenarioAssignment)
		}
	}

	return r0
}

// ToGraphQL provides a mock function with given fields: in
func (_m *Converter) ToGraphQL(in model.AutomaticScenarioAssignment) graphql.AutomaticScenarioAssignment {
	ret := _m.Called(in)

	var r0 graphql.AutomaticScenarioAssignment
	if rf, ok := ret.Get(0).(func(model.AutomaticScenarioAssignment) graphql.AutomaticScenarioAssignment); ok {
		r0 = rf(in)
	} else {
		r0 = ret.Get(0).(graphql.AutomaticScenarioAssignment)
	}

	return r0
}
