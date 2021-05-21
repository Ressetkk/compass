// Code generated by mockery v2.5.1. DO NOT EDIT.

package automock

import (
	context "context"

	model "github.com/kyma-incubator/compass/components/director/internal/model"
	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, item
func (_m *Repository) Create(ctx context.Context, item model.SystemAuth) error {
	ret := _m.Called(ctx, item)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.SystemAuth) error); ok {
		r0 = rf(ctx, item)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteByIDForObject provides a mock function with given fields: ctx, tenant, id, objType
func (_m *Repository) DeleteByIDForObject(ctx context.Context, tenant string, id string, objType model.SystemAuthReferenceObjectType) error {
	ret := _m.Called(ctx, tenant, id, objType)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, model.SystemAuthReferenceObjectType) error); ok {
		r0 = rf(ctx, tenant, id, objType)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteByIDForObjectGlobal provides a mock function with given fields: ctx, id, objType
func (_m *Repository) DeleteByIDForObjectGlobal(ctx context.Context, id string, objType model.SystemAuthReferenceObjectType) error {
	ret := _m.Called(ctx, id, objType)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, model.SystemAuthReferenceObjectType) error); ok {
		r0 = rf(ctx, id, objType)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetByID provides a mock function with given fields: ctx, tenant, id
func (_m *Repository) GetByID(ctx context.Context, tenant string, id string) (*model.SystemAuth, error) {
	ret := _m.Called(ctx, tenant, id)

	var r0 *model.SystemAuth
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *model.SystemAuth); ok {
		r0 = rf(ctx, tenant, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.SystemAuth)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, tenant, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByIDGlobal provides a mock function with given fields: ctx, id
func (_m *Repository) GetByIDGlobal(ctx context.Context, id string) (*model.SystemAuth, error) {
	ret := _m.Called(ctx, id)

	var r0 *model.SystemAuth
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.SystemAuth); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.SystemAuth)
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

// GetByJSONValue provides a mock function with given fields: ctx, value
func (_m *Repository) GetByJSONValue(ctx context.Context, value map[string]interface{}) (*model.SystemAuth, error) {
	ret := _m.Called(ctx, value)

	var r0 *model.SystemAuth
	if rf, ok := ret.Get(0).(func(context.Context, map[string]interface{}) *model.SystemAuth); ok {
		r0 = rf(ctx, value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.SystemAuth)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, map[string]interface{}) error); ok {
		r1 = rf(ctx, value)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListForObject provides a mock function with given fields: ctx, tenant, objectType, objectID
func (_m *Repository) ListForObject(ctx context.Context, tenant string, objectType model.SystemAuthReferenceObjectType, objectID string) ([]model.SystemAuth, error) {
	ret := _m.Called(ctx, tenant, objectType, objectID)

	var r0 []model.SystemAuth
	if rf, ok := ret.Get(0).(func(context.Context, string, model.SystemAuthReferenceObjectType, string) []model.SystemAuth); ok {
		r0 = rf(ctx, tenant, objectType, objectID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.SystemAuth)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, model.SystemAuthReferenceObjectType, string) error); ok {
		r1 = rf(ctx, tenant, objectType, objectID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListForObjectGlobal provides a mock function with given fields: ctx, objectType, objectID
func (_m *Repository) ListForObjectGlobal(ctx context.Context, objectType model.SystemAuthReferenceObjectType, objectID string) ([]model.SystemAuth, error) {
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

// Update provides a mock function with given fields: ctx, item
func (_m *Repository) Update(ctx context.Context, item *model.SystemAuth) error {
	ret := _m.Called(ctx, item)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.SystemAuth) error); ok {
		r0 = rf(ctx, item)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
