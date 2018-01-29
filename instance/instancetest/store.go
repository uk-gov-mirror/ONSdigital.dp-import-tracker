// Code generated by moq; DO NOT EDIT
// github.com/matryer/moq

package instancetest

import (
	"context"
	"sync"
)

var (
	lockStoreMockUpdateInstanceWithHierarchyBuilt   sync.RWMutex
	lockStoreMockUpdateInstanceWithSearchIndexBuilt sync.RWMutex
)

// StoreMock is a mock implementation of Store.
//
//     func TestSomethingThatUsesStore(t *testing.T) {
//
//         // make and configure a mocked Store
//         mockedStore := &StoreMock{
//             UpdateInstanceWithHierarchyBuiltFunc: func(ctx context.Context, instanceID string, dimensionID string) (bool, error) {
// 	               panic("TODO: mock out the UpdateInstanceWithHierarchyBuilt method")
//             },
//             UpdateInstanceWithSearchIndexBuiltFunc: func(ctx context.Context, instanceID string, dimensionID string) (bool, error) {
// 	               panic("TODO: mock out the UpdateInstanceWithSearchIndexBuilt method")
//             },
//         }
//
//         // TODO: use mockedStore in code that requires Store
//         //       and then make assertions.
//
//     }
type StoreMock struct {
	// UpdateInstanceWithHierarchyBuiltFunc mocks the UpdateInstanceWithHierarchyBuilt method.
	UpdateInstanceWithHierarchyBuiltFunc func(ctx context.Context, instanceID string, dimensionID string) (bool, error)

	// UpdateInstanceWithSearchIndexBuiltFunc mocks the UpdateInstanceWithSearchIndexBuilt method.
	UpdateInstanceWithSearchIndexBuiltFunc func(ctx context.Context, instanceID string, dimensionID string) (bool, error)

	// calls tracks calls to the methods.
	calls struct {
		// UpdateInstanceWithHierarchyBuilt holds details about calls to the UpdateInstanceWithHierarchyBuilt method.
		UpdateInstanceWithHierarchyBuilt []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// InstanceID is the instanceID argument value.
			InstanceID string
			// DimensionID is the dimensionID argument value.
			DimensionID string
		}
		// UpdateInstanceWithSearchIndexBuilt holds details about calls to the UpdateInstanceWithSearchIndexBuilt method.
		UpdateInstanceWithSearchIndexBuilt []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// InstanceID is the instanceID argument value.
			InstanceID string
			// DimensionID is the dimensionID argument value.
			DimensionID string
		}
	}
}

// UpdateInstanceWithHierarchyBuilt calls UpdateInstanceWithHierarchyBuiltFunc.
func (mock *StoreMock) UpdateInstanceWithHierarchyBuilt(ctx context.Context, instanceID string, dimensionID string) (bool, error) {
	if mock.UpdateInstanceWithHierarchyBuiltFunc == nil {
		panic("moq: StoreMock.UpdateInstanceWithHierarchyBuiltFunc is nil but Store.UpdateInstanceWithHierarchyBuilt was just called")
	}
	callInfo := struct {
		Ctx         context.Context
		InstanceID  string
		DimensionID string
	}{
		Ctx:         ctx,
		InstanceID:  instanceID,
		DimensionID: dimensionID,
	}
	lockStoreMockUpdateInstanceWithHierarchyBuilt.Lock()
	mock.calls.UpdateInstanceWithHierarchyBuilt = append(mock.calls.UpdateInstanceWithHierarchyBuilt, callInfo)
	lockStoreMockUpdateInstanceWithHierarchyBuilt.Unlock()
	return mock.UpdateInstanceWithHierarchyBuiltFunc(ctx, instanceID, dimensionID)
}

// UpdateInstanceWithHierarchyBuiltCalls gets all the calls that were made to UpdateInstanceWithHierarchyBuilt.
// Check the length with:
//     len(mockedStore.UpdateInstanceWithHierarchyBuiltCalls())
func (mock *StoreMock) UpdateInstanceWithHierarchyBuiltCalls() []struct {
	Ctx         context.Context
	InstanceID  string
	DimensionID string
} {
	var calls []struct {
		Ctx         context.Context
		InstanceID  string
		DimensionID string
	}
	lockStoreMockUpdateInstanceWithHierarchyBuilt.RLock()
	calls = mock.calls.UpdateInstanceWithHierarchyBuilt
	lockStoreMockUpdateInstanceWithHierarchyBuilt.RUnlock()
	return calls
}

// UpdateInstanceWithSearchIndexBuilt calls UpdateInstanceWithSearchIndexBuiltFunc.
func (mock *StoreMock) UpdateInstanceWithSearchIndexBuilt(ctx context.Context, instanceID string, dimensionID string) (bool, error) {
	if mock.UpdateInstanceWithSearchIndexBuiltFunc == nil {
		panic("moq: StoreMock.UpdateInstanceWithSearchIndexBuiltFunc is nil but Store.UpdateInstanceWithSearchIndexBuilt was just called")
	}
	callInfo := struct {
		Ctx         context.Context
		InstanceID  string
		DimensionID string
	}{
		Ctx:         ctx,
		InstanceID:  instanceID,
		DimensionID: dimensionID,
	}
	lockStoreMockUpdateInstanceWithSearchIndexBuilt.Lock()
	mock.calls.UpdateInstanceWithSearchIndexBuilt = append(mock.calls.UpdateInstanceWithSearchIndexBuilt, callInfo)
	lockStoreMockUpdateInstanceWithSearchIndexBuilt.Unlock()
	return mock.UpdateInstanceWithSearchIndexBuiltFunc(ctx, instanceID, dimensionID)
}

// UpdateInstanceWithSearchIndexBuiltCalls gets all the calls that were made to UpdateInstanceWithSearchIndexBuilt.
// Check the length with:
//     len(mockedStore.UpdateInstanceWithSearchIndexBuiltCalls())
func (mock *StoreMock) UpdateInstanceWithSearchIndexBuiltCalls() []struct {
	Ctx         context.Context
	InstanceID  string
	DimensionID string
} {
	var calls []struct {
		Ctx         context.Context
		InstanceID  string
		DimensionID string
	}
	lockStoreMockUpdateInstanceWithSearchIndexBuilt.RLock()
	calls = mock.calls.UpdateInstanceWithSearchIndexBuilt
	lockStoreMockUpdateInstanceWithSearchIndexBuilt.RUnlock()
	return calls
}
