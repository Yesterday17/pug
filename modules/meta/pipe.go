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

package meta

import "github.com/Yesterday17/pug/api"

func (m *Module) Do(prev api.Pipe, pl api.Pipeline) {
	m.Metadata = prev.Meta()

	if m.Title != "" {
		m.Metadata.Title = m.ModuleData.Title
	}

	if m.ModuleData.Author != "" {
		m.Metadata.Author = m.ModuleData.Author
	}

	if m.ModuleData.Description != "" {
		m.Metadata.Description = m.ModuleData.Description
	}

	if m.ModuleData.Cover != "" {
		m.Metadata.Cover = m.ModuleData.Cover
	}

	if m.ModuleData.Link != "" {
		m.Metadata.Link = m.ModuleData.Link
	}

	if m.ModuleData.Short != "" {
		m.Metadata.Short = m.ModuleData.Short
	}

	if m.ModuleData.From != "" {
		m.Metadata.From = m.ModuleData.From
	}

	if m.ModuleData.ReleaseTime != "" {
		m.Metadata.ReleaseTime = m.ModuleData.ReleaseTime
	}
}
