package nuonrunner

import (
	"context"

	"github.com/nuonco/nuon-runner-go/client/operations"
)

func (c *client) LockTerraformWorkspace(ctx context.Context, workspaceID string, jobID *string, reqBody any) (any, error) {
	resp, err := c.genClient.Operations.LockTerraformWorkspace(&operations.LockTerraformWorkspaceParams{
		Body:        reqBody,
		Context:     ctx,
		WorkspaceID: workspaceID,
		JobID:       jobID,
	}, c.getAuthInfo())
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}
