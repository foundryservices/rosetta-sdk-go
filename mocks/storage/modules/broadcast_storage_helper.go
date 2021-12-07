// Code generated by mockery v1.0.0. DO NOT EDIT.

package modules

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	database "github.com/foundryservices/rosetta-sdk-go/storage/database"
	types "github.com/foundryservices/rosetta-sdk-go/types"
)

// BroadcastStorageHelper is an autogenerated mock type for the BroadcastStorageHelper type
type BroadcastStorageHelper struct {
	mock.Mock
}

// AtTip provides a mock function with given fields: _a0, _a1
func (_m *BroadcastStorageHelper) AtTip(_a0 context.Context, _a1 int64) (bool, error) {
	ret := _m.Called(_a0, _a1)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, int64) bool); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BroadcastTransaction provides a mock function with given fields: _a0, _a1, _a2
func (_m *BroadcastStorageHelper) BroadcastTransaction(_a0 context.Context, _a1 *types.NetworkIdentifier, _a2 string) (*types.TransactionIdentifier, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *types.TransactionIdentifier
	if rf, ok := ret.Get(0).(func(context.Context, *types.NetworkIdentifier, string) *types.TransactionIdentifier); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.TransactionIdentifier)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.NetworkIdentifier, string) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CurrentBlockIdentifier provides a mock function with given fields: _a0
func (_m *BroadcastStorageHelper) CurrentBlockIdentifier(_a0 context.Context) (*types.BlockIdentifier, error) {
	ret := _m.Called(_a0)

	var r0 *types.BlockIdentifier
	if rf, ok := ret.Get(0).(func(context.Context) *types.BlockIdentifier); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.BlockIdentifier)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindTransaction provides a mock function with given fields: _a0, _a1, _a2
func (_m *BroadcastStorageHelper) FindTransaction(_a0 context.Context, _a1 *types.TransactionIdentifier, _a2 database.Transaction) (*types.BlockIdentifier, *types.Transaction, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *types.BlockIdentifier
	if rf, ok := ret.Get(0).(func(context.Context, *types.TransactionIdentifier, database.Transaction) *types.BlockIdentifier); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.BlockIdentifier)
		}
	}

	var r1 *types.Transaction
	if rf, ok := ret.Get(1).(func(context.Context, *types.TransactionIdentifier, database.Transaction) *types.Transaction); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*types.Transaction)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *types.TransactionIdentifier, database.Transaction) error); ok {
		r2 = rf(_a0, _a1, _a2)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
