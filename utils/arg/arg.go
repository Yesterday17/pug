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

/**
pug [url] -p [Module_Name] -arg1ofModule1 [val] -arg2ofModule1=[val] -arg3booleanOfModule1
          -p [Module_Name] -arg1ofModule2 [val] -arg2ofModule2=[val] -arg3booleanOfModule2
*/

package arg

import "strings"

// ParseArgs parse commandline of pug
// assume the first element of args is -p
// or len(args) == 0
func ParseArgs(args []string) []map[string]interface{} {
	if args == nil || len(args) == 0 {
		return nil
	}

	ret := make([]map[string]interface{}, 0)

	last := map[string]interface{}{}
	var key string
	var val string
	for index, arg := range args {
		if key != "" {
			if arg == "[" || val != "" {
				val += arg + " "
				if arg == "]" {
					last[key] = val[2 : len(val)-3]
					key = ""
					val = ""
				}
				continue
			}

			k := key
			key = ""
			if arg[0] != byte('-') {
				// Common Value
				if k == "p" {
					k = "module"
				}
				last[k] = arg
				continue
			} else {
				if k == "p" {
					return nil
				}
				last[k] = true
			}
		}

		if arg[0] != byte('-') {
			return nil
		}

		entry := strings.Split(arg, "=")
		if entry[0] == "-p" && index != 0 {
			ret = append(ret, last)
			last = map[string]interface{}{}
		}

		key = entry[0][1:]
		if len(entry) > 1 {
			if key == "p" {
				key = "module"
			}
			last[key] = strings.Join(entry[1:], "=")
			key = ""
		}
	}
	ret = append(ret, last)
	return ret
}
