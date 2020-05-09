package ytdl

import (
	"github.com/Yesterday17/pug/api"
	"github.com/Yesterday17/pug/modules/base"
)

var Module = base.Module(
	"youtube-dl",
	"Download videos with tool youtube-dl",
	[]string{"Yesterday17"},
	"",
	preprocessor,
	map[string]api.PipeBuilder{
		"download": ytdlPipeBuilder,
	},
)
