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

package pugd

import (
	"github.com/Yesterday17/pug/api"
)

type Module struct {
	api.BasePipe
	args map[string]interface{}
}

func (m *Module) Name() string {
	return "PUG daemon"
}

func (m *Module) Description() string {
	return "Background daemon for PUG."
}

func (m *Module) Author() []string {
	return []string{
		"Yesterday17",
	}
}

func NewPUGd(args map[string]interface{}) interface{} {
	return &Module{
		BasePipe: api.BasePipe{
			PStatus: api.PipeWaiting,
		},
		args: args,
	}
}
