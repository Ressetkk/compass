// Code generated by mockery v1.0.0. DO NOT EDIT.

package automock

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	v1alpha1 "github.com/kyma-incubator/compass/components/operations-controller/api/v1alpha1"
)

// K8SClient is an autogenerated mock type for the K8SClient type
type K8SClient struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, operation
func (_m *K8SClient) Create(ctx context.Context, operation *v1alpha1.Operation) (*v1alpha1.Operation, error) {
	ret := _m.Called(ctx, operation)

	var r0 *v1alpha1.Operation
	if rf, ok := ret.Get(0).(func(context.Context, *v1alpha1.Operation) *v1alpha1.Operation); ok {
		r0 = rf(ctx, operation)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.Operation)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1alpha1.Operation) error); ok {
		r1 = rf(ctx, operation)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, name, opts
func (_m *K8SClient) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	ret := _m.Called(ctx, name, opts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, v1.DeleteOptions) error); ok {
		r0 = rf(ctx, name, opts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, name, options
func (_m *K8SClient) Get(ctx context.Context, name string, options v1.GetOptions) (*v1alpha1.Operation, error) {
	ret := _m.Called(ctx, name, options)

	var r0 *v1alpha1.Operation
	if rf, ok := ret.Get(0).(func(context.Context, string, v1.GetOptions) *v1alpha1.Operation); ok {
		r0 = rf(ctx, name, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.Operation)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, v1.GetOptions) error); ok {
		r1 = rf(ctx, name, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, operation
func (_m *K8SClient) Update(ctx context.Context, operation *v1alpha1.Operation) (*v1alpha1.Operation, error) {
	ret := _m.Called(ctx, operation)

	var r0 *v1alpha1.Operation
	if rf, ok := ret.Get(0).(func(context.Context, *v1alpha1.Operation) *v1alpha1.Operation); ok {
		r0 = rf(ctx, operation)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.Operation)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1alpha1.Operation) error); ok {
		r1 = rf(ctx, operation)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
