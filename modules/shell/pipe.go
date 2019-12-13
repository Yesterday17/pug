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

package shell

import (
	"github.com/Yesterday17/pug/api"
	"github.com/Yesterday17/pug/utils/log"
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

func (s *shell) Do(prev api.Pipe) {
	s.prev = prev
	s.PStatus = api.PipeWorking

	cmd := exec.Command(s.command, s.args...)
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
