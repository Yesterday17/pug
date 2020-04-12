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

package types

import "github.com/Yesterday17/pug/api"

func Assert(m map[string]interface{}, key string) (interface{}, api.PipeConstructorError) {
	v, ok := m[key]
	if !ok {
		return nil, api.PipeArgumentMissing
	}
	return v, api.PipeNoError
}

func String(m map[string]interface{}, key string) (string, api.PipeConstructorError) {
	v, err := Assert(m, key)
	if err != api.PipeNoError {
		return "", err
	}

	switch v.(type) {
	case string:
		return v.(string), api.PipeNoError
	default:
		return "", api.PipeArgumentTypeMismatch
	}
}
