package nuonrunner

import (
	"context"

	"github.com/nuonco/nuon-runner-go/models"
)

func (c *client) CreateJobExecution(ctx context.Context, jobID string, req *models.ServiceCreateRunnerJobExecutionRequest) (*models.AppRunnerJobExecution, error) {
	return nil, nil
}

func (c *client) CreateJobExecutionHeartBeat(ctx context.Context, jobExecutionID string, req *models.ServiceCreateRunnerJobExecutionHeartBeatRequest) (*models.AppRunnerJobExecutionHeartBeat, error) {
	return nil, nil
}

func (c *client) UpdateJobExecution(ctx context.Context, jobExecutionID string, req *models.ServiceUpdateRunnerJobExecutionRequest) (*models.AppRunnerJobExecution, error) {
	return nil, nil
}

func (c *client) GetJobs(ctx context.Context, status string, typ string) ([]*models.AppRunnerJob, error) {
	return nil, nil
}

func (c *client) GetJob(ctx context.Context, jobID string) (*models.AppRunnerJob, error) {
	return nil, nil
}

func (c *client) GetJobPlan(ctx context.Context, jobID string) (interface{}, error) {
	return nil, nil
}
