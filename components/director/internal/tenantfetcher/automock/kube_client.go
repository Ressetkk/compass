// Code generated by mockery v2.10.5. DO NOT EDIT.

package automock

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// KubeClient is an autogenerated mock type for the KubeClient type
type KubeClient struct {
	mock.Mock
}

// GetTenantFetcherConfigMapData provides a mock function with given fields: ctx
func (_m *KubeClient) GetTenantFetcherConfigMapData(ctx context.Context) (string, string, error) {
	ret := _m.Called(ctx)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context) string); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(context.Context) string); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context) error); ok {
		r2 = rf(ctx)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// UpdateTenantFetcherConfigMapData provides a mock function with given fields: ctx, lastRunTimestamp, lastResyncTimestamp
func (_m *KubeClient) UpdateTenantFetcherConfigMapData(ctx context.Context, lastRunTimestamp string, lastResyncTimestamp string) error {
	ret := _m.Called(ctx, lastRunTimestamp, lastResyncTimestamp)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, lastRunTimestamp, lastResyncTimestamp)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
