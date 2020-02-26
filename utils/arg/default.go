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

package arg

func GetDefault(args map[string]interface{}, key string, d interface{}) interface{} {
	ret := args[key]
	if ret == nil {
		ret = d
	}
	return ret
}

func GetDefaultString(args map[string]interface{}, key string, d string) string {
	return GetDefault(args, key, d).(string)
}