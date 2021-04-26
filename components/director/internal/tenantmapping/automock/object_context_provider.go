// Code generated by mockery v2.5.1. DO NOT EDIT.

package automock

import (
	context "context"

	oathkeeper "github.com/kyma-incubator/compass/components/director/internal/oathkeeper"
	tenantmapping "github.com/kyma-incubator/compass/components/director/internal/tenantmapping"
	mock "github.com/stretchr/testify/mock"
)

// ObjectContextProvider is an autogenerated mock type for the ObjectContextProvider type
type ObjectContextProvider struct {
	mock.Mock
}

// GetObjectContext provides a mock function with given fields: ctx, reqData, authDetails
func (_m *ObjectContextProvider) GetObjectContext(ctx context.Context, reqData oathkeeper.ReqData, authDetails oathkeeper.AuthDetails) (tenantmapping.ObjectContext, error) {
	ret := _m.Called(ctx, reqData, authDetails)

	var r0 tenantmapping.ObjectContext
	if rf, ok := ret.Get(0).(func(context.Context, oathkeeper.ReqData, oathkeeper.AuthDetails) tenantmapping.ObjectContext); ok {
		r0 = rf(ctx, reqData, authDetails)
	} else {
		r0 = ret.Get(0).(tenantmapping.ObjectContext)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, oathkeeper.ReqData, oathkeeper.AuthDetails) error); ok {
		r1 = rf(ctx, reqData, authDetails)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
