package nuonrunner

import (
	"context"

	"github.com/nuonco/nuon-runner-go/client/operations"
	"github.com/nuonco/nuon-runner-go/models"
)

func (c *client) GetInstallComponentPreviousConfig(ctx context.Context, installId, componentId string) (*models.ServiceGetInstallComponentPreviousConfigResponse, error) {
	resp, err := c.
		genClient.
		Operations.
		GetInstallComponentPreviousConfig(&operations.GetInstallComponentPreviousConfigParams{
			ComponentID: componentId,
			InstallID:   installId,
		}, c.getAuthInfo())
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}
