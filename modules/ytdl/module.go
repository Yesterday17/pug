package ytdl

import (
	"github.com/Yesterday17/pug/api"
	"github.com/Yesterday17/pug/utils/log"
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
	proxy := args["proxy"]
	if args["proxy"] == nil || args["proxy"].(string) == "" {
		log.Fatalf("[Bash] No Command provided!\n")
		proxy = ""
	}

	return &Module{
		BasePipe: api.BasePipe{
			PStatus: api.PipeWaiting,
		},
		Proxy: proxy.(string),
	}
}
