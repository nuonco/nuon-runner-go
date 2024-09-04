package nuonrunner

import (
	"context"

	"github.com/nuonco/nuon-runner-go/client/operations"
	"github.com/nuonco/nuon-runner-go/models"
)

func (c *client) GetJobExecutions(ctx context.Context, jobID string) ([]*models.AppRunnerJobExecution, error) {
	resp, err := c.genClient.Operations.GetRunnerJobExecutions(&operations.GetRunnerJobExecutionsParams{
		RunnerJobID: jobID,
		Context:     ctx,
	}, c.getAuthInfo())
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) CreateJobExecution(ctx context.Context, jobID string, req *models.ServiceCreateRunnerJobExecutionRequest) (*models.AppRunnerJobExecution, error) {
	resp, err := c.genClient.Operations.CreateRunnerJobExecution(&operations.CreateRunnerJobExecutionParams{
		Req:         req,
		RunnerJobID: jobID,
		Context:     ctx,
	}, c.getAuthInfo())
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) UpdateJobExecution(ctx context.Context, jobID, jobExecutionID string, req *models.ServiceUpdateRunnerJobExecutionRequest) (*models.AppRunnerJobExecution, error) {
	resp, err := c.genClient.Operations.UpdateRunnerJobExecution(&operations.UpdateRunnerJobExecutionParams{
		Req:                  req,
		RunnerJobExecutionID: jobExecutionID,
		RunnerJobID:          jobID,
		Context:              ctx,
	}, c.getAuthInfo())
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) CreateJobExecutionResult(ctx context.Context, jobID, jobExecutionID string, req *models.ServiceCreateRunnerJobExecutionResultRequest) (*models.AppRunnerJobExecutionResult, error) {
	resp, err := c.genClient.Operations.CreateRunnerJobExecutionResult(&operations.CreateRunnerJobExecutionResultParams{
		Req:                  req,
		RunnerJobExecutionID: jobExecutionID,
		RunnerJobID:          jobID,
		Context:              ctx,
	}, c.getAuthInfo())
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}
