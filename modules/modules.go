package modules

import (
	"github.com/Yesterday17/pug/modules/shell"
)

type NewFunc func(args map[string]interface{}) interface{}

var Modules = map[string]NewFunc{
	"shell": shell.NewShell,
}
