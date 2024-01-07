package taskdataqueries

import "github.com/golang/mock/gomock"

// mock_taskdataqueries.go

// MockTaskDataQuery is a mock of TaskDataQuery interface
type MockTaskDataQuery struct {
	ctrl     *gomock.Controller
	recorder *MockTaskDataQueryMockRecorder
}

// MockTaskDataQueryMockRecorder is the mock recorder for MockTaskDataQuery
type MockTaskDataQueryMockRecorder struct {
	mock *MockTaskDataQuery
}

// NewMockTaskDataQuery creates a new mock instance
func NewMockTaskDataQuery(ctrl *gomock.Controller) *MockTaskDataQuery {
	mock := &MockTaskDataQuery{ctrl: ctrl}
	mock.recorder = &MockTaskDataQueryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTaskDataQuery) EXPECT() *MockTaskDataQueryMockRecorder {
	return m.recorder
}

// GetDataQuery mocks base method
func (m *MockTaskDataQuery) GetDataQuery(param GetTaskDataParam) ([]TaskDataResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDataQuery", param)
	ret0, _ := ret[0].([]TaskDataResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateDataQuery mocks base method
func (m *MockTaskDataQuery) CreateDataQuery(param CreateTaskDataParam) (error, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDataQuery", param)
	ret0, _ := ret[0].(error)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateDataQuery mocks base method
func (m *MockTaskDataQuery) UpdateDataQuery(param UpdateTaskDataParam) (error, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateDataQuery", param)
	ret0, _ := ret[0].(error)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteDataQuery mocks base method
func (m *MockTaskDataQuery) DeleteDataQuery(param DeleteTaskDataParam) (error, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteDataQuery", param)
	ret0, _ := ret[0].(error)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckSameTasksDataQuery mocks base method
func (m *MockTaskDataQuery) CheckSameTasksDataQuery(param CreateTaskDataParam) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckSameTasksDataQuery", param)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckTaskExistDataQuery mocks base method
func (m *MockTaskDataQuery) CheckTaskExistDataQuery(param DeleteTaskDataParam) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckTaskExistDataQuery", param)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}
