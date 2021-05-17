// Code generated by mockery v2.2.1. DO NOT EDIT.

package automock

import (
	context "context"

	model "github.com/kyma-incubator/compass/components/director/internal/model"
	mock "github.com/stretchr/testify/mock"
)

// EventService is an autogenerated mock type for the EventService type
type EventService struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, appID, bundleID, packageID, in, specs, bundleIDs
func (_m *EventService) Create(ctx context.Context, appID string, bundleID *string, packageID *string, in model.EventDefinitionInput, specs []*model.SpecInput, bundleIDs []string) (string, error) {
	ret := _m.Called(ctx, appID, bundleID, packageID, in, specs, bundleIDs)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, string, *string, *string, model.EventDefinitionInput, []*model.SpecInput, []string) string); ok {
		r0 = rf(ctx, appID, bundleID, packageID, in, specs, bundleIDs)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, *string, *string, model.EventDefinitionInput, []*model.SpecInput, []string) error); ok {
		r1 = rf(ctx, appID, bundleID, packageID, in, specs, bundleIDs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id
func (_m *EventService) Delete(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ListByApplicationID provides a mock function with given fields: ctx, appID
func (_m *EventService) ListByApplicationID(ctx context.Context, appID string) ([]*model.EventDefinition, error) {
	ret := _m.Called(ctx, appID)

	var r0 []*model.EventDefinition
	if rf, ok := ret.Get(0).(func(context.Context, string) []*model.EventDefinition); ok {
		r0 = rf(ctx, appID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.EventDefinition)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, appID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateInManyBundles provides a mock function with given fields: ctx, id, in, specIn, bundleIDsForCreation, bundleIDsForDeletion
func (_m *EventService) UpdateInManyBundles(ctx context.Context, id string, in model.EventDefinitionInput, specIn *model.SpecInput, bundleIDsForCreation []string, bundleIDsForDeletion []string) error {
	ret := _m.Called(ctx, id, in, specIn, bundleIDsForCreation, bundleIDsForDeletion)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, model.EventDefinitionInput, *model.SpecInput, []string, []string) error); ok {
		r0 = rf(ctx, id, in, specIn, bundleIDsForCreation, bundleIDsForDeletion)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
