package nuonrunner

import (
	"context"

	"github.com/nuonco/nuon-go/models"
)

func (c *client) SetRunnerID(runnerID string) {
	c.RunnerID = runnerID
}

func (c *client) GetSettings(ctx context.Context) (*models.AppRunnerGroupSettings, error) {
	return nil, nil
}
