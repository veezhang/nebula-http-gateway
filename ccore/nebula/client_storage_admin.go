package nebula

import "github.com/vesoft-inc/nebula-http-gateway/ccore/nebula/types"

type (
	StorageAdminClient interface {
		Open() error
		Close() error
	}

	defaultStorageAdminClient defaultClient
)

func NewStorageAdminClient(endpoints []string, opts ...Option) (StorageAdminClient, error) {
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

	return &defaultStorageAdminClient{
		o:            o,
		driver:       driver,
		storageAdmin: newDriverStorageAdmin(endpoints, &o.storageAdmin),
	}, nil
}

func (c *defaultStorageAdminClient) Open() error {
	return c.storageAdmin.open(c.driver)
}

func (c *defaultStorageAdminClient) Close() error {
	return c.storageAdmin.open(c.driver)
}
