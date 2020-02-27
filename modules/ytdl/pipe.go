package ytdl

import (
	"bufio"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/Yesterday17/pug/api"
	"github.com/Yesterday17/pug/utils/log"
)

func (m *Module) Do(prev api.Pipe, pl api.Pipeline) {
	m.PStatus = api.PipeWorking
	m.Metadata = prev.Meta()
	m.MediaData = prev.Media()

	args := []string{
		prev.Meta().Link,
		"-o", "%(id)s.%(ext)s",
		"-f", "bestvideo+bestaudio",
		"--merge-output-format", "mkv",
	}

	if m.Proxy != "" {
		args = append(args, "--proxy", m.Proxy)
	}

	cmd := exec.Command("youtube-dl", args...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Errorf("%s\n", err.Error())
		m.PStatus = api.PipeError
		return
	}
	go func() {
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			log.Infof("%s\n", scanner.Text())
		}
	}()

	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Error(err.Error())
		m.PStatus = api.PipeError
		return
	}
	go func() {
		scanner := bufio.NewScanner(stderr)
		for scanner.Scan() {
			log.Errorf("%s\n", scanner.Text())
		}
	}()

	err = cmd.Start()
	if err != nil {
		log.Errorf("%s\n", err.Error())
		m.PStatus = api.PipeError
		return
	}

	err = cmd.Wait()
	if err != nil {
		log.Errorf("%s\n", err.Error())
		m.PStatus = api.PipeError
		return
	}

	// Metadata
	if strings.Contains(m.Metadata.Link, "youtube.com/watch?v=") {
		// Youtube
		m.Metadata.Short = strings.ReplaceAll(m.Metadata.Link, "https://www.youtube.com/watch?v=", "")
	}

	// Media
	path, _ := os.Getwd()
	m.MediaData.Path = filepath.Join(path, m.Metadata.Short+".mkv")
	log.Infof("%s\n", m.MediaData.Path)
}
