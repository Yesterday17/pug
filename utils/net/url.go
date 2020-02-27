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

package net

import u "net/url"

func BuildUrl(domain string, ssl bool, path string, params map[string]string) string {
	// Protocol
	url := "http"
	if ssl {
		url += "s"
	}
	url += "://"
	// Domain
	url += domain + "/"
	// Path
	url += path
	// Params
	if len(params) > 0 {
		url += "?"
		for key, value := range params {
			url += u.QueryEscape(key)
			url += "="
			url += u.QueryEscape(value)
			url += "&"
		}
	}
	return url
}
