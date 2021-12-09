package nebula

import "github.com/vesoft-inc/nebula-http-gateway/ccore/nebula/types"

type (
	GraphClient interface {
		Open() error
		Execute(stmt []byte) (ExecutionResponse, error)
		ExecuteJson(stmt []byte) ([]byte, error)
		Close() error
	}

	defaultGraphClient defaultClient
)

func NewGraphClient(endpoints []string, username, password string, opts ...Option) (GraphClient, error) {
	o := defaultOptions()
	for _, opt := range opts {
		opt(&o)
	}
	o.complete()
	if err := o.validate(); err != nil {
		return nil, err
	}

	driver, err := types.GetDriver(o.version)
	if err != nil {
		return nil, err
	}

	return &defaultGraphClient{
		o:      o,
		driver: driver,
		graph:  newDriverGraph(endpoints, username, password, &o.graph),
	}, nil
}

func (c *defaultGraphClient) Open() error {
	return c.graph.open(c.driver)
}

func (c *defaultGraphClient) Execute(stmt []byte) (ExecutionResponse, error) {
	return c.graph.Execute(c.graph.sessionId, stmt)
}

func (c *defaultGraphClient) ExecuteJson(stmt []byte) ([]byte, error) {
	if err := c.graph.open(c.driver); err != nil {
		return nil, err
	}

	return c.graph.ExecuteJson(c.graph.sessionId, stmt)
}

func (c *defaultGraphClient) Close() error {
	return c.graph.open(c.driver)
}
