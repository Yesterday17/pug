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

	if m.Title != "" {
		m.Metadata.Title = m.Title
	}

	if m.AuthorInfo != "" {
		m.Metadata.Author = m.AuthorInfo
	}

	if m.DescriptionInfo != "" {
		m.Metadata.Description = m.DescriptionInfo
	}

	if m.Cover != "" {
		m.Metadata.Cover = m.Cover
	}

	if m.Link != "" {
		m.Metadata.Link = m.Link
	}

	if m.Short != "" {
		m.Metadata.Short = m.Short
	}

	if m.From != "" {
		m.Metadata.From = m.From
	}

	if m.ReleaseTime != "" {
		m.Metadata.ReleaseTime = m.ReleaseTime
	}
}
