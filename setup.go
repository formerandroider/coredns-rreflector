package rreflector

import (
	"github.com/caddyserver/caddy"
	"github.com/coredns/coredns/core/dnsserver"
	"github.com/coredns/coredns/plugin"
)

func init() {
	plugin.Register("rreflector", setup)
}

func setup(c *caddy.Controller) error {
	handler, err := rreflectorParse(c)
	if err != nil {
		return plugin.Error("rreflector", err)
	}

	dnsserver.GetConfig(c).AddPlugin(func(next plugin.Handler) plugin.Handler {
		return handler
	})

	return nil
}

func rreflectorParse(c *caddy.Controller) (*rreflectorHandler, error) {
	handler := newRReflectorHandler()

	i := 0
	for c.Next() {
		if i > 0 {
			return nil, plugin.ErrOnce
		}
		i++

		args := c.RemainingArgs()
		switch len(args) {
		case 0:
		case 1:
			handler.host = args[0]
		default:
			return nil, c.ArgErr()
		}
	}

	return handler, nil
}
