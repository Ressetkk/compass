// Code generated by mockery v2.9.4. DO NOT EDIT.

package automock

import (
	context "context"

	model "github.com/kyma-incubator/compass/components/director/internal/model"
	mock "github.com/stretchr/testify/mock"
)

// BusinessTenantMappingService is an autogenerated mock type for the BusinessTenantMappingService type
type BusinessTenantMappingService struct {
	mock.Mock
}

// DeleteMany provides a mock function with given fields: ctx, tenantInputs
func (_m *BusinessTenantMappingService) DeleteMany(ctx context.Context, tenantInputs []string) error {
	ret := _m.Called(ctx, tenantInputs)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []string) error); ok {
		r0 = rf(ctx, tenantInputs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetTenantByExternalID provides a mock function with given fields: ctx, externalID
func (_m *BusinessTenantMappingService) GetTenantByExternalID(ctx context.Context, externalID string) (*model.BusinessTenantMapping, error) {
	ret := _m.Called(ctx, externalID)

	var r0 *model.BusinessTenantMapping
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.BusinessTenantMapping); ok {
		r0 = rf(ctx, externalID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.BusinessTenantMapping)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, externalID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ctx
func (_m *BusinessTenantMappingService) List(ctx context.Context) ([]*model.BusinessTenantMapping, error) {
	ret := _m.Called(ctx)

	var r0 []*model.BusinessTenantMapping
	if rf, ok := ret.Get(0).(func(context.Context) []*model.BusinessTenantMapping); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.BusinessTenantMapping)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListLabels provides a mock function with given fields: ctx, tenantID
func (_m *BusinessTenantMappingService) ListLabels(ctx context.Context, tenantID string) (map[string]*model.Label, error) {
	ret := _m.Called(ctx, tenantID)

	var r0 map[string]*model.Label
	if rf, ok := ret.Get(0).(func(context.Context, string) map[string]*model.Label); ok {
		r0 = rf(ctx, tenantID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]*model.Label)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, tenantID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListPageBySearchTerm provides a mock function with given fields: ctx, searchTerm, pageSize, cursor
func (_m *BusinessTenantMappingService) ListPageBySearchTerm(ctx context.Context, searchTerm string, pageSize int, cursor string) (*model.BusinessTenantMappingPage, error) {
	ret := _m.Called(ctx, searchTerm, pageSize, cursor)

	var r0 *model.BusinessTenantMappingPage
	if rf, ok := ret.Get(0).(func(context.Context, string, int, string) *model.BusinessTenantMappingPage); ok {
		r0 = rf(ctx, searchTerm, pageSize, cursor)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.BusinessTenantMappingPage)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, int, string) error); ok {
		r1 = rf(ctx, searchTerm, pageSize, cursor)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, id, tenantInput
func (_m *BusinessTenantMappingService) Update(ctx context.Context, id string, tenantInput model.BusinessTenantMappingInput) error {
	ret := _m.Called(ctx, id, tenantInput)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, model.BusinessTenantMappingInput) error); ok {
		r0 = rf(ctx, id, tenantInput)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpsertMany provides a mock function with given fields: ctx, tenantInputs
func (_m *BusinessTenantMappingService) UpsertMany(ctx context.Context, tenantInputs ...model.BusinessTenantMappingInput) error {
	_va := make([]interface{}, len(tenantInputs))
	for _i := range tenantInputs {
		_va[_i] = tenantInputs[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, ...model.BusinessTenantMappingInput) error); ok {
		r0 = rf(ctx, tenantInputs...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
