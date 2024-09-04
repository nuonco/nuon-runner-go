// Code generated by MockGen. DO NOT EDIT.
// Source: client.go

// Package nuonrunner is a generated GoMock package.
package nuonrunner

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/nuonco/nuon-runner-go/models"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// CreateHealthCheck mocks base method.
func (m *MockClient) CreateHealthCheck(ctx context.Context, req *models.ServiceCreateRunnerHealthCheckRequest) (*models.AppRunnerHealthCheck, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateHealthCheck", ctx, req)
	ret0, _ := ret[0].(*models.AppRunnerHealthCheck)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateHealthCheck indicates an expected call of CreateHealthCheck.
func (mr *MockClientMockRecorder) CreateHealthCheck(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateHealthCheck", reflect.TypeOf((*MockClient)(nil).CreateHealthCheck), ctx, req)
}

// CreateHeartBeat mocks base method.
func (m *MockClient) CreateHeartBeat(ctx context.Context, req *models.ServiceCreateRunnerHeartBeatRequest) (*models.AppRunnerHeartBeat, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateHeartBeat", ctx, req)
	ret0, _ := ret[0].(*models.AppRunnerHeartBeat)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateHeartBeat indicates an expected call of CreateHeartBeat.
func (mr *MockClientMockRecorder) CreateHeartBeat(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateHeartBeat", reflect.TypeOf((*MockClient)(nil).CreateHeartBeat), ctx, req)
}

// CreateJobExecution mocks base method.
func (m *MockClient) CreateJobExecution(ctx context.Context, jobID string, req *models.ServiceCreateRunnerJobExecutionRequest) (*models.AppRunnerJobExecution, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateJobExecution", ctx, jobID, req)
	ret0, _ := ret[0].(*models.AppRunnerJobExecution)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateJobExecution indicates an expected call of CreateJobExecution.
func (mr *MockClientMockRecorder) CreateJobExecution(ctx, jobID, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateJobExecution", reflect.TypeOf((*MockClient)(nil).CreateJobExecution), ctx, jobID, req)
}

// CreateJobExecutionResult mocks base method.
func (m *MockClient) CreateJobExecutionResult(ctx context.Context, jobID, jobExecutionID string, req *models.ServiceCreateRunnerJobExecutionResultRequest) (*models.AppRunnerJobExecutionResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateJobExecutionResult", ctx, jobID, jobExecutionID, req)
	ret0, _ := ret[0].(*models.AppRunnerJobExecutionResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateJobExecutionResult indicates an expected call of CreateJobExecutionResult.
func (mr *MockClientMockRecorder) CreateJobExecutionResult(ctx, jobID, jobExecutionID, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateJobExecutionResult", reflect.TypeOf((*MockClient)(nil).CreateJobExecutionResult), ctx, jobID, jobExecutionID, req)
}

// GetJob mocks base method.
func (m *MockClient) GetJob(ctx context.Context, jobID string) (*models.AppRunnerJob, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetJob", ctx, jobID)
	ret0, _ := ret[0].(*models.AppRunnerJob)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetJob indicates an expected call of GetJob.
func (mr *MockClientMockRecorder) GetJob(ctx, jobID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetJob", reflect.TypeOf((*MockClient)(nil).GetJob), ctx, jobID)
}

// GetJobExecutions mocks base method.
func (m *MockClient) GetJobExecutions(ctx context.Context, jobID string) ([]*models.AppRunnerJobExecution, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetJobExecutions", ctx, jobID)
	ret0, _ := ret[0].([]*models.AppRunnerJobExecution)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetJobExecutions indicates an expected call of GetJobExecutions.
func (mr *MockClientMockRecorder) GetJobExecutions(ctx, jobID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetJobExecutions", reflect.TypeOf((*MockClient)(nil).GetJobExecutions), ctx, jobID)
}

// GetJobPlan mocks base method.
func (m *MockClient) GetJobPlan(ctx context.Context, jobID string) (*models.Planv1Plan, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetJobPlan", ctx, jobID)
	ret0, _ := ret[0].(*models.Planv1Plan)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetJobPlan indicates an expected call of GetJobPlan.
func (mr *MockClientMockRecorder) GetJobPlan(ctx, jobID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetJobPlan", reflect.TypeOf((*MockClient)(nil).GetJobPlan), ctx, jobID)
}

// GetJobs mocks base method.
func (m *MockClient) GetJobs(ctx context.Context, grp models.AppRunnerJobGroup, status models.AppRunnerJobStatus, limit *int64) ([]*models.AppRunnerJob, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetJobs", ctx, grp, status, limit)
	ret0, _ := ret[0].([]*models.AppRunnerJob)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetJobs indicates an expected call of GetJobs.
func (mr *MockClientMockRecorder) GetJobs(ctx, grp, status, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetJobs", reflect.TypeOf((*MockClient)(nil).GetJobs), ctx, grp, status, limit)
}

// GetSettings mocks base method.
func (m *MockClient) GetSettings(ctx context.Context) (*models.AppRunnerGroupSettings, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSettings", ctx)
	ret0, _ := ret[0].(*models.AppRunnerGroupSettings)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSettings indicates an expected call of GetSettings.
func (mr *MockClientMockRecorder) GetSettings(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSettings", reflect.TypeOf((*MockClient)(nil).GetSettings), ctx)
}

// SetRunnerID mocks base method.
func (m *MockClient) SetRunnerID(runnerID string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetRunnerID", runnerID)
}

// SetRunnerID indicates an expected call of SetRunnerID.
func (mr *MockClientMockRecorder) SetRunnerID(runnerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetRunnerID", reflect.TypeOf((*MockClient)(nil).SetRunnerID), runnerID)
}

// UpdateJobExecution mocks base method.
func (m *MockClient) UpdateJobExecution(ctx context.Context, jobID, jobExecutionID string, req *models.ServiceUpdateRunnerJobExecutionRequest) (*models.AppRunnerJobExecution, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateJobExecution", ctx, jobID, jobExecutionID, req)
	ret0, _ := ret[0].(*models.AppRunnerJobExecution)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateJobExecution indicates an expected call of UpdateJobExecution.
func (mr *MockClientMockRecorder) UpdateJobExecution(ctx, jobID, jobExecutionID, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateJobExecution", reflect.TypeOf((*MockClient)(nil).UpdateJobExecution), ctx, jobID, jobExecutionID, req)
}

// WriteOTELLogs mocks base method.
func (m *MockClient) WriteOTELLogs(ctx context.Context, req interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteOTELLogs", ctx, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteOTELLogs indicates an expected call of WriteOTELLogs.
func (mr *MockClientMockRecorder) WriteOTELLogs(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteOTELLogs", reflect.TypeOf((*MockClient)(nil).WriteOTELLogs), ctx, req)
}

// WriteOTELMetrics mocks base method.
func (m *MockClient) WriteOTELMetrics(ctx context.Context, req interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteOTELMetrics", ctx, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteOTELMetrics indicates an expected call of WriteOTELMetrics.
func (mr *MockClientMockRecorder) WriteOTELMetrics(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteOTELMetrics", reflect.TypeOf((*MockClient)(nil).WriteOTELMetrics), ctx, req)
}

// WriteOTELTraces mocks base method.
func (m *MockClient) WriteOTELTraces(ctx context.Context, req interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteOTELTraces", ctx, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteOTELTraces indicates an expected call of WriteOTELTraces.
func (mr *MockClientMockRecorder) WriteOTELTraces(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteOTELTraces", reflect.TypeOf((*MockClient)(nil).WriteOTELTraces), ctx, req)
}
