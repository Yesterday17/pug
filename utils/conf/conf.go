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

package conf

import (
	"errors"
	"reflect"
	"strings"
)

func Serialize(s interface{}) (result string) {
	r := reflect.TypeOf(s)
	value := reflect.ValueOf(s)

	if r.Kind() == reflect.Ptr {
		r = r.Elem()
		value = value.Elem()
	}

	for i := 0; i < r.NumField(); i++ {
		field := r.Field(i)
		key := field.Name

		if conf, ok := field.Tag.Lookup("conf"); ok {
			if conf == "-" || strings.Contains(conf, "=") {
				continue
			}
			key = conf
		}

		switch value.Field(i).Kind() {
		case reflect.String:
			result += key + "=" + value.Field(i).String() + "\n"
		default:
			continue
		}
	}
	return
}

func Deserialize(s string, result interface{}) error {
	r := reflect.TypeOf(result)

	if r.Kind() != reflect.Ptr {
		return errors.New("result must be a pointer")
	}

	r = r.Elem()
	value := reflect.ValueOf(result).Elem()

	values := map[string]string{}
	for _, v := range strings.Split(s, "\n") {
		entry := strings.Split(v, "=")
		if len(entry) != 2 {
			continue
		}
		values[entry[0]] = entry[1]
	}

	for i := 0; i < r.NumField(); i++ {
		field := r.Field(i)

		if conf, ok := field.Tag.Lookup("conf"); ok {
			if val, ok := values[conf]; ok {
				value.Field(i).SetString(val)
			}
		}
	}
	return nil
}
