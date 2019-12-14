/**
PUG
Copyright (C) 2019  Yesterday17

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
	"github.com/Yesterday17/pug/api"
	"github.com/Yesterday17/pug/utils/log"
	"os"
	"os/exec"
)

func (s *shell) Meta() api.Metadata {
	if s.prev == nil {
		return nil
	}
	return s.prev.Meta()
}

func (s *shell) Media() api.Media {
	if s.prev == nil {
		return nil
	}
	return s.prev.Media()
}

func (s *shell) Do(prev api.Pipe, pl api.Pipeline) {
	s.prev = prev
	s.PStatus = api.PipeWorking

	// Environmental Variables
	PUGPrevMedia := "PUG_PREV_MEDIA="
	PUGPrevMeta := "PUG_PREV_META="
	PUGOutputMedia := pl.TempDir().NewFile(".conf")
	PUGOutputMeta := pl.TempDir().NewFile(".conf")

	cmd := exec.Command("bash", "-c", s.command)
	cmd.Env = append(os.Environ(),
		"PUG_VERSION="+api.VERSION,
		PUGPrevMedia,
		PUGPrevMeta,
		"PUG_OUTPUT_MEDIA="+PUGOutputMedia,
		"PUG_OUTPUT_META="+PUGOutputMeta,
	)
	output, err := cmd.Output()
	if err != nil {
		log.Error(err.Error())
		log.Error(string(output))
		s.PStatus = api.PipeError
	} else {
		log.Info(string(output))
		s.PStatus = api.PipeSuccess
	}
}
