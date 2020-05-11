package ytdl

import (
	"os"
	"os/exec"
	"path/filepath"
	"reflect"

	"github.com/Yesterday17/pug/api"
	"github.com/Yesterday17/pug/modules/base"
	"github.com/Yesterday17/pug/utils/log"
)

type ytdlPipe struct {
}

var ytdlPipeBuilder = base.NewPipeBuilder(map[string]reflect.Kind{"url": reflect.String}, map[string]reflect.Kind{"proxy": reflect.String}, newYtdlPipe)

func newYtdlPipe(m map[string]interface{}) (api.Pipe, api.PipeBuildError) {
	return &ytdlPipe{}, api.PipeNoError
}

func (y *ytdlPipe) Validate() map[string]reflect.Kind {
	return map[string]reflect.Kind{
		"!url":  reflect.String,
		"+file": reflect.String,
	}
}

func (y *ytdlPipe) Execute(work api.State) (err error) {
	var proxy, url string

	if work.Has("proxy") {
		proxy, err = work.GetString(proxy)
		if err != nil {
			return
		}
	}

	url, _ = work.GetString("url")

	args := []string{
		url,
		"-o", "%(id)s.%(ext)s",
		"-f", "bestvideo+bestaudio",
		"--merge-output-format", "mkv",
		"--newline",
	}

	if proxy != "" {
		args = append(args, "--proxy", proxy)
	}

	cmd := exec.Command("youtube-dl", args...)
	cmd.Stdout = log.DefaultLogger.WrappedLogWriter
	cmd.Stderr = log.DefaultLogger.WrappedErrorWriter

	err = cmd.Start()
	if err != nil {
		return
	}

	err = cmd.Wait()
	if err != nil {
		return
	}

	// Media
	path, _ := os.Getwd()
	file := filepath.Join(path, "result.mkv")
	work.Set("file", file)

	return nil
}

func (y *ytdlPipe) Clone() api.Pipe {
	return &ytdlPipe{}
}
