package nuonrunner

import (
	"fmt"
	"net/http"
	"net/url"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-playground/validator/v10"

	genclient "github.com/nuonco/nuon-go/client"
)

//
//go:generate -command swagger go run github.com/go-swagger/go-swagger/cmd/swagger
//go:generate swagger generate client --skip-tag-packages -f ./generate.switch
//go:generate -command mockgen go run github.com/golang/mock/mockgen
//go:generate mockgen -destination=mock.go -source=client.go -package=nuonrunner
type Client interface {
	// SetRunnerID(runnerID string)
}

var _ Client = (*client)(nil)

type client struct {
	v *validator.Validate

	APIURL   string `validate:"required"`
	APIToken string
	RunnerID string

	genClient    *genclient.Nuon
	appTransport *appTransport
}

type clientOption func(*client) error

func New(opts ...clientOption) (*client, error) {
	c := &client{}
	for _, opt := range opts {
		if err := opt(c); err != nil {
			return nil, err
		}
	}

	if c.v == nil {
		c.v = validator.New()
	}

	if err := c.v.Struct(c); err != nil {
		return nil, err
	}

	apiURL, err := url.Parse(c.APIURL)
	if err != nil {
		return nil, fmt.Errorf("unable to parse api url: %w", err)
	}

	transport := httptransport.New(apiURL.Host, "", []string{apiURL.Scheme})
	appTransport := &appTransport{
		authToken: c.APIToken,
		transport: http.DefaultTransport,
	}
	c.appTransport = appTransport
	transport.Transport = appTransport
	genClient := genclient.New(transport, nil)
	c.genClient = genClient

	return c, nil
}

// WithAuthToken specifies the auth token to use
func WithAuthToken(token string) clientOption {
	return func(c *client) error {
		c.APIToken = token
		return nil
	}
}

// WithURL specifies the url to use
func WithURL(url string) clientOption {
	return func(c *client) error {
		c.APIURL = url
		return nil
	}
}

// WithRunnerID specifies the runner id to use
func WithRunnerID(runnerID string) clientOption {
	return func(c *client) error {
		c.RunnerID = runnerID
		return nil
	}
}

// WithValidator specifies a validator to use
func WithValidator(v *validator.Validate) clientOption {
	return func(c *client) error {
		c.v = v
		return nil
	}
}
