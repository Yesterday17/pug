/**
PUG
Copyright (C) 2019-2020  Yesterday17

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package bash

import (
	"bufio"
	"os"
	"os/exec"

	"github.com/Yesterday17/pug/api"
	"github.com/Yesterday17/pug/utils/conf"
	"github.com/Yesterday17/pug/utils/log"
)

func (m *Module) Do(prev api.Pipe, pl api.Pipeline) {
	if m.Command == "" {
		return
	}

	m.PStatus = api.PipeWorking
	m.Metadata = prev.Meta()
	m.MediaData = prev.Media()

	// Environmental Variables
	PUGPrevMedia, err := pl.TempDir().NewContentFile(prev.Media().Serialize(), ".conf")
	// TODO: Embed error handle in api
	if err != nil {
		log.Errorf("%s\n", err.Error())
		m.PStatus = api.PipeError
		return
	}
	PUGPrevMeta, err := pl.TempDir().NewContentFile(prev.Meta().Serialize(), ".conf")
	// TODO: Embed error handle in api
	if err != nil {
		log.Errorf("%s\n", err.Error())
		m.PStatus = api.PipeError
		return
	}
	PUGOutputMedia := pl.TempDir().NewFile(".conf")
	PUGOutputMeta := pl.TempDir().NewFile(".conf")

	cmd := exec.Command("bash", "-c", m.Command)

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

	cmd.Env = append(os.Environ(),
		"PUG_VERSION="+api.VERSION,
		"PUG_PREV_MEDIA="+PUGPrevMedia,
		"PUG_PREV_META="+PUGPrevMeta,
		"PUG_OUTPUT_MEDIA="+PUGOutputMedia,
		"PUG_OUTPUT_META="+PUGOutputMeta,
	)

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

	// Load output file
	// MENTION: NO ERROR HANDLE HERE
	_ = conf.ReadAndDeserialize(PUGOutputMedia, &m.MediaData)
	_ = conf.ReadAndDeserialize(PUGOutputMeta, &m.Metadata)

	m.PStatus = api.PipeSuccess
	return
}
