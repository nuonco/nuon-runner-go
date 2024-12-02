package nuonrunner

import (
	"context"

	"github.com/nuonco/nuon-runner-go/client/operations"
	"github.com/nuonco/nuon-runner-go/models"
)

func (c *client) PublishMetrics(ctx context.Context, req []*models.ServicePublishMetricInput) error {
	_, err := c.genClient.Operations.PublishMetrics(&operations.PublishMetricsParams{
		Req:     req,
		Context: ctx,
	}, c.getAuthInfo())
	if err != nil {
		return err
	}

	return nil
}
