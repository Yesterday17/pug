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

package base

import "github.com/Yesterday17/pug/api"

type module struct {
	name        string
	description string
	author      []string
	usage       string

	preprocessor api.Preprocessor
	pipes        map[string]api.PipeBuilder
}

func Module(name, description string, author []string, usage string, preprocessor api.Preprocessor, pipes map[string]api.PipeBuilder) api.Module {
	return &module{
		name:         name,
		description:  description,
		author:       author,
		usage:        usage,
		preprocessor: preprocessor,
		pipes:        pipes,
	}
}

func (m *module) Name() string {
	return m.name
}

func (m *module) Description() string {
	return m.description
}

func (m *module) Author() []string {
	return m.author
}

func (m *module) Usage() string {
	return m.usage
}

func (m *module) Preprocessor() api.Preprocessor {
	return m.preprocessor
}

func (m *module) Pipe(pid string) (api.PipeBuilder, bool) {
	p, ok := m.pipes[pid]
	return p, ok
}

func (m *module) Pipes() []string {
	var ret []string
	for name := range m.pipes {
		ret = append(ret, name)
	}

	return ret
}
