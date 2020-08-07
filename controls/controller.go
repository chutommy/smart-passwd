package controls

import (
	mrand "math/rand"

	data "github.com/chutified/smart-passwd/data"
	random "github.com/chutified/smart-passwd/random"
	swaps "github.com/chutified/smart-passwd/swaps"
	"github.com/pkg/errors"
)

// Controller controls the password generation.
type Controller struct {
	rng *mrand.Rand
	se  *swaps.SwapEngine
	ds  *data.Service
}

// New constructs the Controller.
func New() *Controller {
	return &Controller{
		rng: random.GetRNG(),
		se:  swaps.New(),
		ds:  data.New(),
	}
}

// Init initializes all services.
func (c *Controller) Init() error {

	// Init data service.
	err := c.ds.Init(data.ConfigDB{
		Host:     "localhost",
		Port:     5432,
		User:     "tommychu",
		Password: "ferdajekamarad",
		DBname:   "word_list",
	})
	if err != nil {
		return errors.Wrap(err, "initializing data service")
	}

	return nil
}

// Stop stops all connections.
func (c *Controller) Stop() error {
	err := c.ds.Stop()
	if err != nil {
		return errors.Wrap(err, "closing data service's connection")
	}

	return nil
}
