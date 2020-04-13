package ytdl

import (
	"github.com/Yesterday17/pug/api"
	"github.com/Yesterday17/pug/modules/base"
)

var preprocessor = base.Preprocessor(
	"youtube\\.com/watch\\?v=[a-zA-Z0-9\\-]+",
	func(m map[string]interface{}, input string) (api.State, error) {
		state := base.NewState()
		state.Set("url", input)
		return state, nil
	},
)
