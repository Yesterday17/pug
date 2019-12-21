package bilibili

import (
	"github.com/Yesterday17/pug/api"
	"github.com/Yesterday17/pug/utils/log"
	"os"
)

func (m *Module) Do(prev api.Pipe, pl api.Pipeline) {
	log.Warn("WARN Non-EndPointPipe utility of bilibili module has not been implemented!")
}

func (m *Module) PipeOut(prev api.Pipe, pl api.Pipeline) {
	file, err := os.Open(prev.Media().Path)
	if err != nil {
		//
	}

	m.SelectRoute()

	err = m.PreUpload(file)
	if err != nil {
		//
	}

	err = m.UploadsPost()
	if err != nil {
		//
	}

	m.SplitChunks(file)
	m.EmitUpload(file)
	m.AfterUpload(file)
}
