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

// CreateJobExecutionOutputs mocks base method.
func (m *MockClient) CreateJobExecutionOutputs(ctx context.Context, jobID, jobExecutionID string, req *models.ServiceCreateRunnerJobExecutionOutputsRequest) (*models.AppRunnerJobExecutionOutputs, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateJobExecutionOutputs", ctx, jobID, jobExecutionID, req)
	ret0, _ := ret[0].(*models.AppRunnerJobExecutionOutputs)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateJobExecutionOutputs indicates an expected call of CreateJobExecutionOutputs.
func (mr *MockClientMockRecorder) CreateJobExecutionOutputs(ctx, jobID, jobExecutionID, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateJobExecutionOutputs", reflect.TypeOf((*MockClient)(nil).CreateJobExecutionOutputs), ctx, jobID, jobExecutionID, req)
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

// GetActionWorkflowConfig mocks base method.
func (m *MockClient) GetActionWorkflowConfig(ctx context.Context, workflowConfigID string) (*models.AppActionWorkflowConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActionWorkflowConfig", ctx, workflowConfigID)
	ret0, _ := ret[0].(*models.AppActionWorkflowConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetActionWorkflowConfig indicates an expected call of GetActionWorkflowConfig.
func (mr *MockClientMockRecorder) GetActionWorkflowConfig(ctx, workflowConfigID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActionWorkflowConfig", reflect.TypeOf((*MockClient)(nil).GetActionWorkflowConfig), ctx, workflowConfigID)
}

// GetActionWorkflowLatestConfig mocks base method.
func (m *MockClient) GetActionWorkflowLatestConfig(ctx context.Context, workflowID string) (*models.AppActionWorkflowConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActionWorkflowLatestConfig", ctx, workflowID)
	ret0, _ := ret[0].(*models.AppActionWorkflowConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetActionWorkflowLatestConfig indicates an expected call of GetActionWorkflowLatestConfig.
func (mr *MockClientMockRecorder) GetActionWorkflowLatestConfig(ctx, workflowID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActionWorkflowLatestConfig", reflect.TypeOf((*MockClient)(nil).GetActionWorkflowLatestConfig), ctx, workflowID)
}

// GetAppConfig mocks base method.
func (m *MockClient) GetAppConfig(ctx context.Context, appID, appConfigID string) (*models.AppAppConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAppConfig", ctx, appID, appConfigID)
	ret0, _ := ret[0].(*models.AppAppConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAppConfig indicates an expected call of GetAppConfig.
func (mr *MockClientMockRecorder) GetAppConfig(ctx, appID, appConfigID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAppConfig", reflect.TypeOf((*MockClient)(nil).GetAppConfig), ctx, appID, appConfigID)
}

// GetInstallActionWorkflowRun mocks base method.
func (m *MockClient) GetInstallActionWorkflowRun(ctx context.Context, installID, runID string) (*models.AppInstallActionWorkflowRun, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInstallActionWorkflowRun", ctx, installID, runID)
	ret0, _ := ret[0].(*models.AppInstallActionWorkflowRun)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInstallActionWorkflowRun indicates an expected call of GetInstallActionWorkflowRun.
func (mr *MockClientMockRecorder) GetInstallActionWorkflowRun(ctx, installID, runID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInstallActionWorkflowRun", reflect.TypeOf((*MockClient)(nil).GetInstallActionWorkflowRun), ctx, installID, runID)
}

// GetInstallComponenetLastActivePlan mocks base method.
func (m *MockClient) GetInstallComponenetLastActivePlan(ctx context.Context, installId, componentId string) (*models.ServiceGetInstallComponenetLastActivePlanResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInstallComponenetLastActivePlan", ctx, installId, componentId)
	ret0, _ := ret[0].(*models.ServiceGetInstallComponenetLastActivePlanResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInstallComponenetLastActivePlan indicates an expected call of GetInstallComponenetLastActivePlan.
func (mr *MockClientMockRecorder) GetInstallComponenetLastActivePlan(ctx, installId, componentId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInstallComponenetLastActivePlan", reflect.TypeOf((*MockClient)(nil).GetInstallComponenetLastActivePlan), ctx, installId, componentId)
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

// GetJobPlanJSON mocks base method.
func (m *MockClient) GetJobPlanJSON(ctx context.Context, jobID string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetJobPlanJSON", ctx, jobID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetJobPlanJSON indicates an expected call of GetJobPlanJSON.
func (mr *MockClientMockRecorder) GetJobPlanJSON(ctx, jobID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetJobPlanJSON", reflect.TypeOf((*MockClient)(nil).GetJobPlanJSON), ctx, jobID)
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

// LockTerraformWorkspace mocks base method.
func (m *MockClient) LockTerraformWorkspace(ctx context.Context, workspaceID string, jobID *string, reqBody any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LockTerraformWorkspace", ctx, workspaceID, jobID, reqBody)
	ret0, _ := ret[0].(error)
	return ret0
}

// LockTerraformWorkspace indicates an expected call of LockTerraformWorkspace.
func (mr *MockClientMockRecorder) LockTerraformWorkspace(ctx, workspaceID, jobID, reqBody interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LockTerraformWorkspace", reflect.TypeOf((*MockClient)(nil).LockTerraformWorkspace), ctx, workspaceID, jobID, reqBody)
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

// UnlockTerraformWorkspace mocks base method.
func (m *MockClient) UnlockTerraformWorkspace(ctx context.Context, workspaceID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnlockTerraformWorkspace", ctx, workspaceID)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnlockTerraformWorkspace indicates an expected call of UnlockTerraformWorkspace.
func (mr *MockClientMockRecorder) UnlockTerraformWorkspace(ctx, workspaceID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnlockTerraformWorkspace", reflect.TypeOf((*MockClient)(nil).UnlockTerraformWorkspace), ctx, workspaceID)
}

// UpdateInstallActionWorkflowRunStep mocks base method.
func (m *MockClient) UpdateInstallActionWorkflowRunStep(ctx context.Context, installID, workflowID, runID string, req *models.ServiceUpdateInstallActionWorkflowRunStepRequest) (*models.AppInstallActionWorkflowRunStep, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateInstallActionWorkflowRunStep", ctx, installID, workflowID, runID, req)
	ret0, _ := ret[0].(*models.AppInstallActionWorkflowRunStep)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateInstallActionWorkflowRunStep indicates an expected call of UpdateInstallActionWorkflowRunStep.
func (mr *MockClientMockRecorder) UpdateInstallActionWorkflowRunStep(ctx, installID, workflowID, runID, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateInstallActionWorkflowRunStep", reflect.TypeOf((*MockClient)(nil).UpdateInstallActionWorkflowRunStep), ctx, installID, workflowID, runID, req)
}

// UpdateJob mocks base method.
func (m *MockClient) UpdateJob(ctx context.Context, jobID string, req *models.ServiceUpdateRunnerJobRequest) (*models.AppRunnerJob, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateJob", ctx, jobID, req)
	ret0, _ := ret[0].(*models.AppRunnerJob)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateJob indicates an expected call of UpdateJob.
func (mr *MockClientMockRecorder) UpdateJob(ctx, jobID, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateJob", reflect.TypeOf((*MockClient)(nil).UpdateJob), ctx, jobID, req)
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

// UpdateTerraformStateJSON mocks base method.
func (m *MockClient) UpdateTerraformStateJSON(ctx context.Context, workspaceID string, jobID *string, reqBody any) (any, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTerraformStateJSON", ctx, workspaceID, jobID, reqBody)
	ret0, _ := ret[0].(any)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTerraformStateJSON indicates an expected call of UpdateTerraformStateJSON.
func (mr *MockClientMockRecorder) UpdateTerraformStateJSON(ctx, workspaceID, jobID, reqBody interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTerraformStateJSON", reflect.TypeOf((*MockClient)(nil).UpdateTerraformStateJSON), ctx, workspaceID, jobID, reqBody)
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
