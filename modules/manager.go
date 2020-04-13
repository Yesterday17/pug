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

	Manager = NewManager()
)

func NewManager() api.ModuleManager {
	m := &moduleManager{modules: map[string]api.Module{}}

	// Basic modules
	_ = m.Add(shell.Module)
	_ = m.Add(ytdl.Module)

	return m
}

type moduleManager struct {
	modules       map[string]api.Module
	preprocessors []api.Preprocessor
}

func (m *moduleManager) Exist(mid string) bool {
	_, ok := m.modules[mid]
	return ok
}

func (m *moduleManager) ExistPipe(mid, pid string) bool {
	ok := m.Exist(mid)
	if !ok {
		return false
	}

	_, ok = m.modules[mid].Pipe(pid)
	return ok
}

func (m *moduleManager) Add(module api.Module) error {
	if m.Exist(module.Name()) {
		return ModuleNameDuplicated
	}
	m.modules[module.Name()] = module

	p := module.Preprocessor()
	if p != nil {
		m.preprocessors = append(m.preprocessors, p)
	}
	return nil
}

func (m *moduleManager) Module(mid string) api.Module {
	module, ok := m.modules[mid]
	if !ok {
		return nil
	}
	return module
}

func (m *moduleManager) Pipe(mid, pid string) api.PipeConstructor {
	module := m.Module(mid)
	if module == nil {
		return nil
	}

	pipe, ok := module.Pipe(pid)
	if !ok {
		return nil
	}
	return pipe
}

func (m *moduleManager) Preprocessors() []api.Preprocessor {
	return m.preprocessors
}
