// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	"context"
	"github.com/ONSdigital/dp-api-clients-go/dataset"
	"github.com/ONSdigital/dp-import-tracker/api"
	"net/url"
	"sync"
)

var (
	lockDatasetClientMockGetInstance                  sync.RWMutex
	lockDatasetClientMockGetInstances                 sync.RWMutex
	lockDatasetClientMockPutInstanceImportTasks       sync.RWMutex
	lockDatasetClientMockPutInstanceState             sync.RWMutex
	lockDatasetClientMockUpdateInstanceWithNewInserts sync.RWMutex
)

// Ensure, that DatasetClientMock does implement api.DatasetClient.
// If this is not the case, regenerate this file with moq.
var _ api.DatasetClient = &DatasetClientMock{}

// DatasetClientMock is a mock implementation of api.DatasetClient.
//
//     func TestSomethingThatUsesDatasetClient(t *testing.T) {
//
//         // make and configure a mocked api.DatasetClient
//         mockedDatasetClient := &DatasetClientMock{
//             GetInstanceFunc: func(ctx context.Context, userAuthToken string, serviceAuthToken string, collectionID string, instanceID string) (dataset.Instance, error) {
// 	               panic("mock out the GetInstance method")
//             },
//             GetInstancesFunc: func(ctx context.Context, userAuthToken string, serviceAuthToken string, collectionID string, vars url.Values) (dataset.Instances, error) {
// 	               panic("mock out the GetInstances method")
//             },
//             PutInstanceImportTasksFunc: func(ctx context.Context, serviceAuthToken string, instanceID string, data dataset.InstanceImportTasks) error {
// 	               panic("mock out the PutInstanceImportTasks method")
//             },
//             PutInstanceStateFunc: func(ctx context.Context, serviceAuthToken string, instanceID string, state dataset.State) error {
// 	               panic("mock out the PutInstanceState method")
//             },
//             UpdateInstanceWithNewInsertsFunc: func(ctx context.Context, serviceAuthToken string, instanceID string, observationsInserted int32) error {
// 	               panic("mock out the UpdateInstanceWithNewInserts method")
//             },
//         }
//
//         // use mockedDatasetClient in code that requires api.DatasetClient
//         // and then make assertions.
//
//     }
type DatasetClientMock struct {
	// GetInstanceFunc mocks the GetInstance method.
	GetInstanceFunc func(ctx context.Context, userAuthToken string, serviceAuthToken string, collectionID string, instanceID string) (dataset.Instance, error)

	// GetInstancesFunc mocks the GetInstances method.
	GetInstancesFunc func(ctx context.Context, userAuthToken string, serviceAuthToken string, collectionID string, vars url.Values) (dataset.Instances, error)

	// PutInstanceImportTasksFunc mocks the PutInstanceImportTasks method.
	PutInstanceImportTasksFunc func(ctx context.Context, serviceAuthToken string, instanceID string, data dataset.InstanceImportTasks) error

	// PutInstanceStateFunc mocks the PutInstanceState method.
	PutInstanceStateFunc func(ctx context.Context, serviceAuthToken string, instanceID string, state dataset.State) error

	// UpdateInstanceWithNewInsertsFunc mocks the UpdateInstanceWithNewInserts method.
	UpdateInstanceWithNewInsertsFunc func(ctx context.Context, serviceAuthToken string, instanceID string, observationsInserted int32) error

	// calls tracks calls to the methods.
	calls struct {
		// GetInstance holds details about calls to the GetInstance method.
		GetInstance []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// UserAuthToken is the userAuthToken argument value.
			UserAuthToken string
			// ServiceAuthToken is the serviceAuthToken argument value.
			ServiceAuthToken string
			// CollectionID is the collectionID argument value.
			CollectionID string
			// InstanceID is the instanceID argument value.
			InstanceID string
		}
		// GetInstances holds details about calls to the GetInstances method.
		GetInstances []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// UserAuthToken is the userAuthToken argument value.
			UserAuthToken string
			// ServiceAuthToken is the serviceAuthToken argument value.
			ServiceAuthToken string
			// CollectionID is the collectionID argument value.
			CollectionID string
			// Vars is the vars argument value.
			Vars url.Values
		}
		// PutInstanceImportTasks holds details about calls to the PutInstanceImportTasks method.
		PutInstanceImportTasks []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ServiceAuthToken is the serviceAuthToken argument value.
			ServiceAuthToken string
			// InstanceID is the instanceID argument value.
			InstanceID string
			// Data is the data argument value.
			Data dataset.InstanceImportTasks
		}
		// PutInstanceState holds details about calls to the PutInstanceState method.
		PutInstanceState []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ServiceAuthToken is the serviceAuthToken argument value.
			ServiceAuthToken string
			// InstanceID is the instanceID argument value.
			InstanceID string
			// State is the state argument value.
			State dataset.State
		}
		// UpdateInstanceWithNewInserts holds details about calls to the UpdateInstanceWithNewInserts method.
		UpdateInstanceWithNewInserts []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ServiceAuthToken is the serviceAuthToken argument value.
			ServiceAuthToken string
			// InstanceID is the instanceID argument value.
			InstanceID string
			// ObservationsInserted is the observationsInserted argument value.
			ObservationsInserted int32
		}
	}
}

// GetInstance calls GetInstanceFunc.
func (mock *DatasetClientMock) GetInstance(ctx context.Context, userAuthToken string, serviceAuthToken string, collectionID string, instanceID string) (dataset.Instance, error) {
	if mock.GetInstanceFunc == nil {
		panic("DatasetClientMock.GetInstanceFunc: method is nil but DatasetClient.GetInstance was just called")
	}
	callInfo := struct {
		Ctx              context.Context
		UserAuthToken    string
		ServiceAuthToken string
		CollectionID     string
		InstanceID       string
	}{
		Ctx:              ctx,
		UserAuthToken:    userAuthToken,
		ServiceAuthToken: serviceAuthToken,
		CollectionID:     collectionID,
		InstanceID:       instanceID,
	}
	lockDatasetClientMockGetInstance.Lock()
	mock.calls.GetInstance = append(mock.calls.GetInstance, callInfo)
	lockDatasetClientMockGetInstance.Unlock()
	return mock.GetInstanceFunc(ctx, userAuthToken, serviceAuthToken, collectionID, instanceID)
}

// GetInstanceCalls gets all the calls that were made to GetInstance.
// Check the length with:
//     len(mockedDatasetClient.GetInstanceCalls())
func (mock *DatasetClientMock) GetInstanceCalls() []struct {
	Ctx              context.Context
	UserAuthToken    string
	ServiceAuthToken string
	CollectionID     string
	InstanceID       string
} {
	var calls []struct {
		Ctx              context.Context
		UserAuthToken    string
		ServiceAuthToken string
		CollectionID     string
		InstanceID       string
	}
	lockDatasetClientMockGetInstance.RLock()
	calls = mock.calls.GetInstance
	lockDatasetClientMockGetInstance.RUnlock()
	return calls
}

// GetInstances calls GetInstancesFunc.
func (mock *DatasetClientMock) GetInstances(ctx context.Context, userAuthToken string, serviceAuthToken string, collectionID string, vars url.Values) (dataset.Instances, error) {
	if mock.GetInstancesFunc == nil {
		panic("DatasetClientMock.GetInstancesFunc: method is nil but DatasetClient.GetInstances was just called")
	}
	callInfo := struct {
		Ctx              context.Context
		UserAuthToken    string
		ServiceAuthToken string
		CollectionID     string
		Vars             url.Values
	}{
		Ctx:              ctx,
		UserAuthToken:    userAuthToken,
		ServiceAuthToken: serviceAuthToken,
		CollectionID:     collectionID,
		Vars:             vars,
	}
	lockDatasetClientMockGetInstances.Lock()
	mock.calls.GetInstances = append(mock.calls.GetInstances, callInfo)
	lockDatasetClientMockGetInstances.Unlock()
	return mock.GetInstancesFunc(ctx, userAuthToken, serviceAuthToken, collectionID, vars)
}

// GetInstancesCalls gets all the calls that were made to GetInstances.
// Check the length with:
//     len(mockedDatasetClient.GetInstancesCalls())
func (mock *DatasetClientMock) GetInstancesCalls() []struct {
	Ctx              context.Context
	UserAuthToken    string
	ServiceAuthToken string
	CollectionID     string
	Vars             url.Values
} {
	var calls []struct {
		Ctx              context.Context
		UserAuthToken    string
		ServiceAuthToken string
		CollectionID     string
		Vars             url.Values
	}
	lockDatasetClientMockGetInstances.RLock()
	calls = mock.calls.GetInstances
	lockDatasetClientMockGetInstances.RUnlock()
	return calls
}

// PutInstanceImportTasks calls PutInstanceImportTasksFunc.
func (mock *DatasetClientMock) PutInstanceImportTasks(ctx context.Context, serviceAuthToken string, instanceID string, data dataset.InstanceImportTasks) error {
	if mock.PutInstanceImportTasksFunc == nil {
		panic("DatasetClientMock.PutInstanceImportTasksFunc: method is nil but DatasetClient.PutInstanceImportTasks was just called")
	}
	callInfo := struct {
		Ctx              context.Context
		ServiceAuthToken string
		InstanceID       string
		Data             dataset.InstanceImportTasks
	}{
		Ctx:              ctx,
		ServiceAuthToken: serviceAuthToken,
		InstanceID:       instanceID,
		Data:             data,
	}
	lockDatasetClientMockPutInstanceImportTasks.Lock()
	mock.calls.PutInstanceImportTasks = append(mock.calls.PutInstanceImportTasks, callInfo)
	lockDatasetClientMockPutInstanceImportTasks.Unlock()
	return mock.PutInstanceImportTasksFunc(ctx, serviceAuthToken, instanceID, data)
}

// PutInstanceImportTasksCalls gets all the calls that were made to PutInstanceImportTasks.
// Check the length with:
//     len(mockedDatasetClient.PutInstanceImportTasksCalls())
func (mock *DatasetClientMock) PutInstanceImportTasksCalls() []struct {
	Ctx              context.Context
	ServiceAuthToken string
	InstanceID       string
	Data             dataset.InstanceImportTasks
} {
	var calls []struct {
		Ctx              context.Context
		ServiceAuthToken string
		InstanceID       string
		Data             dataset.InstanceImportTasks
	}
	lockDatasetClientMockPutInstanceImportTasks.RLock()
	calls = mock.calls.PutInstanceImportTasks
	lockDatasetClientMockPutInstanceImportTasks.RUnlock()
	return calls
}

// PutInstanceState calls PutInstanceStateFunc.
func (mock *DatasetClientMock) PutInstanceState(ctx context.Context, serviceAuthToken string, instanceID string, state dataset.State) error {
	if mock.PutInstanceStateFunc == nil {
		panic("DatasetClientMock.PutInstanceStateFunc: method is nil but DatasetClient.PutInstanceState was just called")
	}
	callInfo := struct {
		Ctx              context.Context
		ServiceAuthToken string
		InstanceID       string
		State            dataset.State
	}{
		Ctx:              ctx,
		ServiceAuthToken: serviceAuthToken,
		InstanceID:       instanceID,
		State:            state,
	}
	lockDatasetClientMockPutInstanceState.Lock()
	mock.calls.PutInstanceState = append(mock.calls.PutInstanceState, callInfo)
	lockDatasetClientMockPutInstanceState.Unlock()
	return mock.PutInstanceStateFunc(ctx, serviceAuthToken, instanceID, state)
}

// PutInstanceStateCalls gets all the calls that were made to PutInstanceState.
// Check the length with:
//     len(mockedDatasetClient.PutInstanceStateCalls())
func (mock *DatasetClientMock) PutInstanceStateCalls() []struct {
	Ctx              context.Context
	ServiceAuthToken string
	InstanceID       string
	State            dataset.State
} {
	var calls []struct {
		Ctx              context.Context
		ServiceAuthToken string
		InstanceID       string
		State            dataset.State
	}
	lockDatasetClientMockPutInstanceState.RLock()
	calls = mock.calls.PutInstanceState
	lockDatasetClientMockPutInstanceState.RUnlock()
	return calls
}

// UpdateInstanceWithNewInserts calls UpdateInstanceWithNewInsertsFunc.
func (mock *DatasetClientMock) UpdateInstanceWithNewInserts(ctx context.Context, serviceAuthToken string, instanceID string, observationsInserted int32) error {
	if mock.UpdateInstanceWithNewInsertsFunc == nil {
		panic("DatasetClientMock.UpdateInstanceWithNewInsertsFunc: method is nil but DatasetClient.UpdateInstanceWithNewInserts was just called")
	}
	callInfo := struct {
		Ctx                  context.Context
		ServiceAuthToken     string
		InstanceID           string
		ObservationsInserted int32
	}{
		Ctx:                  ctx,
		ServiceAuthToken:     serviceAuthToken,
		InstanceID:           instanceID,
		ObservationsInserted: observationsInserted,
	}
	lockDatasetClientMockUpdateInstanceWithNewInserts.Lock()
	mock.calls.UpdateInstanceWithNewInserts = append(mock.calls.UpdateInstanceWithNewInserts, callInfo)
	lockDatasetClientMockUpdateInstanceWithNewInserts.Unlock()
	return mock.UpdateInstanceWithNewInsertsFunc(ctx, serviceAuthToken, instanceID, observationsInserted)
}

// UpdateInstanceWithNewInsertsCalls gets all the calls that were made to UpdateInstanceWithNewInserts.
// Check the length with:
//     len(mockedDatasetClient.UpdateInstanceWithNewInsertsCalls())
func (mock *DatasetClientMock) UpdateInstanceWithNewInsertsCalls() []struct {
	Ctx                  context.Context
	ServiceAuthToken     string
	InstanceID           string
	ObservationsInserted int32
} {
	var calls []struct {
		Ctx                  context.Context
		ServiceAuthToken     string
		InstanceID           string
		ObservationsInserted int32
	}
	lockDatasetClientMockUpdateInstanceWithNewInserts.RLock()
	calls = mock.calls.UpdateInstanceWithNewInserts
	lockDatasetClientMockUpdateInstanceWithNewInserts.RUnlock()
	return calls
}
