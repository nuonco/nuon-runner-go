package nuonrunner

import (
	"context"

	"github.com/nuonco/nuon-runner-go/client/operations"
	"github.com/nuonco/nuon-runner-go/models"
)

func (c *client) GetInstallComponentPreviousConfig(ctx context.Context, installID, componentID string) (*models.ServiceGetInstallComponentPreviousConfigResponse, error) {
	resp, err := c.
		genClient.
		Operations.
		GetInstallComponentPreviousConfig(&operations.GetInstallComponentPreviousConfigParams{
			ComponentID: componentID,
			InstallID:   installID,
			Context:     ctx,
		}, c.getAuthInfo())
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}
