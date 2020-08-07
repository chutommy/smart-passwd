package handlers

import (
	config "github.com/chutified/smart-passwd/config"
	controls "github.com/chutified/smart-passwd/controls"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type PWDhandler struct {
	pwdCtrl controls.Controller
}

func NewPWD() *PWDhandler {
	return &PWDhandler{}
}

func (h *PWDhandler) Init(cfg config.DBConfig) error {

	// init the controller
	err := h.pwdCtrl.Init(cfg)
	if err != nil {
		return errors.Wrap(err, "initializing password controller")
	}

	return nil
}

func (h *PWDhandler) Close() error {
	//stop the controller
	err := h.pwdCtrl.Stop()
	if err != nil {
		return errors.Wrap(err, "stoping password controller")
	}

	return nil
}

func (h *PWDhandler) Gen(c *gin.Context) {}
