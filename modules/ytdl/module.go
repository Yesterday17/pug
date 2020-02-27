package ytdl

import (
	"github.com/Yesterday17/pug/api"
	"github.com/Yesterday17/pug/utils/describe"
	"github.com/Yesterday17/pug/utils/log"
)

type Module struct {
	api.BasePipe
	ModuleData
}

type ModuleData struct {
	Proxy string `mapstructure:"proxy"`
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
	var module Module
	err := describe.NewDescribe(args).Decode("", &module.ModuleData)
	if err != nil {
		log.Fatalf("[youtube-dl] Failed to parse arguments!")
		module.BasePipe.PStatus = api.PipeError
	}
	return &module
}
