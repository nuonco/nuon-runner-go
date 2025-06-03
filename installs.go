package nuonrunner

import (
	"context"

	"github.com/nuonco/nuon-runner-go/client/operations"
	"github.com/nuonco/nuon-runner-go/models"
)

func (c *client) CreateInstallPlan(ctx context.Context, installID string, req *models.ServiceCreateInstallPlanRequest) (*models.AppInstallPlan, error) {
	resp, err := c.genClient.Operations.CreateInstallPlan(&operations.CreateInstallPlanParams{
		InstallID: installID,
		Req:       req,
		Context:   ctx,
	}, c.getAuthInfo())
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *client) GetInstallOwnerPlan(ctx context.Context, installID, OwnerID string) (*models.AppInstallPlan, error) {
	resp, err := c.genClient.Operations.GetInstallOwnerPlan(&operations.GetInstallOwnerPlanParams{
		InstallID: installID,
		OwnerID:   OwnerID,
		Context:   ctx,
	}, c.getAuthInfo())
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}
