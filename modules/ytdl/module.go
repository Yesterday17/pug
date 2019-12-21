package ytdl

import (
	"github.com/Yesterday17/pug/api"
	"github.com/Yesterday17/pug/utils/arg"
)

type Module struct {
	api.BasePipe

	Proxy string
}

func (m *Module) Name() string {
	return "youtube-dl"
}

func (m *Module) Description() string {
	return "Download videos with tool youtube-dl."
}

func (m *Module) Author() []string {
	return []string{
		"Yesterday17",
	}
}

func NewYtDl(args map[string]interface{}) interface{} {
	return &Module{
		BasePipe: api.BasePipe{
			PStatus: api.PipeWaiting,
		},
		Proxy: arg.GetDefaultString(args, "proxy", ""),
	}
}
