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

package modules

import (
	"errors"

	"github.com/Yesterday17/pug/api"
	"github.com/Yesterday17/pug/modules/shell"
	"github.com/Yesterday17/pug/modules/ytdl"
)

var (
	ModuleNameDuplicated = errors.New("module name duplicated")
	modules              = map[string]api.Module{}
)

func init() {
	_ = AddModule(shell.Module)
	_ = AddModule(ytdl.Module)
}

func AddModule(module api.Module) error {
	_, ok := modules[module.Name()]
	if ok {
		return ModuleNameDuplicated
	}
	return nil
}
