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

package describe

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func Load(file string) (Described, error) {
	desc := described{root: make(map[string]interface{}), pl: []string{}}
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(data, &desc.root)
	if err != nil {
		return nil, err
	}

	err = desc.Decode("pipeline", &desc.pl)
	if err != nil {
		return nil, err
	}

	return &desc, nil
}

func NewDescribe(params map[string]interface{}) Described {
	var desc = described{pl: []string{}, root: params}
	for k := range params {
		desc.pl = append(desc.pl, k)
	}
	return &desc
}
