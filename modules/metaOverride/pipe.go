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

package metaOverride

import "github.com/Yesterday17/pug/api"

func (m *Module) Do(prev api.Pipe, pl api.Pipeline) {
	m.Metadata = prev.Meta()

	if m.title != "" {
		m.Metadata.Title = m.title
	}

	if m.author != "" {
		m.Metadata.Author = m.author
	}

	if m.description != "" {
		m.Metadata.Description = m.description
	}

	if m.cover != "" {
		m.Metadata.Cover = m.cover
	}

	if m.link != "" {
		m.Metadata.Link = m.link
	}

	if m.short != "" {
		m.Metadata.Short = m.short
	}

	if m.from != "" {
		m.Metadata.From = m.from
	}

	if m.releaseTime != "" {
		m.Metadata.ReleaseTime = m.releaseTime
	}
}
