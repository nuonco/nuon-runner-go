package nuonrunner

import (
	"context"

	"github.com/nuonco/nuon-runner-go/client/operations"
)

func (c *client) UnlockTerraformWorkspace(ctx context.Context, workspaceID string, reqBody any) (any, error) {
	resp, err := c.genClient.Operations.UnlockTerraformWorkspace(&operations.UnlockTerraformWorkspaceParams{
		Body:        reqBody,
		Context:     ctx,
		WorkspaceID: workspaceID,
	}, c.getAuthInfo())
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}
