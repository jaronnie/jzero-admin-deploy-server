package custom

import "server/server/errcodes"

type Custom struct {
}

func New() *Custom {
	return &Custom{}
}

func (c *Custom) Start() {
	errcodes.Register()
}

func (c *Custom) Stop() {}
