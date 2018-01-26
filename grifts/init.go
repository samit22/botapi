package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/samitghimire/botapi/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
