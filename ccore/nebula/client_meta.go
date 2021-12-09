package nebula

import "github.com/vesoft-inc/nebula-http-gateway/ccore/nebula/types"

type (
	MetaClient interface {
		Open() error
		Close() error
	}

	defaultMetaClient defaultClient
)

func NewMetaClient(endpoints []string,  opts ...Option) (MetaClient, error) {
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

	return &defaultMetaClient{
		o:            o,
		driver:       driver,
		meta:         newDriverMeta(endpoints, &o.meta),
	}, nil
}

func (c *defaultMetaClient) Open() error {
	return c.meta.open(c.driver)
}

func (c *defaultMetaClient) Close() error {
	return c.meta.open(c.driver)
}
