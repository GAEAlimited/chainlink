// Code generated by mockery v2.42.2. DO NOT EDIT.

package client

import (
	context "context"

	types "github.com/smartcontractkit/chainlink/v2/common/types"
	mock "github.com/stretchr/testify/mock"
)

// mockSendOnlyClient is an autogenerated mock type for the sendOnlyClient type
type mockSendOnlyClient[CHAIN_ID types.ID] struct {
	mock.Mock
}

// ChainID provides a mock function with given fields: _a0
func (_m *mockSendOnlyClient[CHAIN_ID]) ChainID(_a0 context.Context) (CHAIN_ID, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for ChainID")
	}

	var r0 CHAIN_ID
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (CHAIN_ID, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(context.Context) CHAIN_ID); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(CHAIN_ID)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Close provides a mock function with given fields:
func (_m *mockSendOnlyClient[CHAIN_ID]) Close() {
	_m.Called()
}

// Dial provides a mock function with given fields: ctx
func (_m *mockSendOnlyClient[CHAIN_ID]) Dial(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Dial")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// newMockSendOnlyClient creates a new instance of mockSendOnlyClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockSendOnlyClient[CHAIN_ID types.ID](t interface {
	mock.TestingT
	Cleanup(func())
}) *mockSendOnlyClient[CHAIN_ID] {
	mock := &mockSendOnlyClient[CHAIN_ID]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
