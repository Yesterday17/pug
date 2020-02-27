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
	if len(args) == 0 {
		return nil
	}

	ret := make([]map[string]interface{}, 0)
	config := map[string]interface{}{
		"module": "config",
	}

	last := map[string]interface{}{}
	lastConfigModified := false
	var key, value string
	for index, arg := range args {
		if key != "" { // Parse value
			if (len(arg) > 0 && arg[0] == '[') || value != "" {
				// built-in argument with space
				// [arg with space]
				value += arg + " "
				if arg == "]" || arg[len(arg)-1] == ']' {
					last[key] = value[2 : len(value)-3]
					key = ""
					value = ""
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
				if k == "c" {
					config["file"] = arg
					lastConfigModified = true
				} else {
					last[k] = arg
				}
				continue
			} else {
				if k == "p" || k == "c" {
					// -p [REQUIRED name]
					// -c [REQUIRED config]
					return nil
				}
				last[k] = true
			}
		}

		// Parse Key
		if arg[0] != byte('-') {
			// key[0] is not '-'
			return nil
		}

		entry := strings.Split(arg, "=")
		if (entry[0] == "-p" || entry[0] == "-c") && index != 0 {
			if !lastConfigModified {
				ret = append(ret, last)
				last = map[string]interface{}{}
			}
			lastConfigModified = false
		}

		key = entry[0][1:]
		if len(entry) > 1 {
			// key=value
			if key == "p" {
				key = "module"
			}
			value := strings.Join(entry[1:], "=")

			if key == "c" {
				config["file"] = value
				lastConfigModified = true
			} else {
				last[key] = value
			}
			key = ""
		}
	}
	ret = append(ret, last, config)
	return ret
}
