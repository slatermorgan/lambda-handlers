// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import (
	address "bitbucket.org/oneiota/mesh-connect-models/address"

	mock "github.com/stretchr/testify/mock"
)

// Connector is an autogenerated mock type for the Connector type
type Connector struct {
	mock.Mock
}

// Authorize provides a mock function with given fields: token
func (_m *Connector) Authorize(token string) error {
	ret := _m.Called(token)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(token)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Find provides a mock function with given fields: postcode
func (_m *Connector) Find(postcode string) ([]*address.Address, error) {
	ret := _m.Called(postcode)

	var r0 []*address.Address
	if rf, ok := ret.Get(0).(func(string) []*address.Address); ok {
		r0 = rf(postcode)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*address.Address)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(postcode)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Retrieve provides a mock function with given fields: id
func (_m *Connector) Retrieve(id string) (*address.LookupRetrieve, error) {
	ret := _m.Called(id)

	var r0 *address.LookupRetrieve
	if rf, ok := ret.Get(0).(func(string) *address.LookupRetrieve); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*address.LookupRetrieve)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type NewConnectorT interface {
	mock.TestingT
	Cleanup(func())
}

// NewConnector creates a new instance of Connector. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewConnector(t NewConnectorT) *Connector {
	mock := &Connector{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
