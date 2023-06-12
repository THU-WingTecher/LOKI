// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	comm "github.com/hyperledger/fabric/core/comm"

	mock "github.com/stretchr/testify/mock"
)

// EndpointUpdater is an autogenerated mock type for the EndpointUpdater type
type EndpointUpdater struct {
	mock.Mock
}

// GetEndpoint provides a mock function with given fields:
func (_m *EndpointUpdater) GetEndpoint() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// UpdateEndpoints provides a mock function with given fields: endpoints
func (_m *EndpointUpdater) UpdateEndpoints(endpoints []comm.EndpointCriteria) {
	_m.Called(endpoints)
}
