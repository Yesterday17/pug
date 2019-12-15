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
)

type shell struct {
	api.BasePipe

	command string
}

func (s *shell) Name() string {
	return "Shell"
}

func (s *shell) Description() string {
	return "It does nothing to pipe contents, but just runs a command."
}

func (s *shell) Author() []string {
	return []string{
		"Yesterday17",
	}
}

func NewBash(args map[string]interface{}) interface{} {
	cmd := args["cmd"]
	if args["cmd"] == nil || args["cmd"].(string) == "" {
		log.Fatalf("[Bash] No command provided!\n")
		cmd = ""
	}
	return &shell{
		BasePipe: api.BasePipe{
			PStatus: api.PipeError,
		},
		command: cmd.(string),
	}
}
