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

package describe

import (
	"errors"
	"reflect"

	"github.com/mitchellh/mapstructure"
)

type Described interface {
	Exists(key string) bool
	Sub(key string) Described
	Extract(key string) Described
	ExtractStrict(key string) (Described, error)
	Decode(key string, dst interface{}) error
	Range(func(key string))

	Root() map[string]interface{}
}

type described struct {
	root map[string]interface{}
	pl   []string
}

func (d *described) Sub(key string) Described {
	var desc = described{
		root: map[string]interface{}{},
		pl:   []string{},
	}
	ex := d.Extract(key)
	switch ex.Root()[key].(type) {
	case map[interface{}]interface{}:
		m := ex.Root()[key].(map[interface{}]interface{})
		for k, v := range m {
			switch k.(type) {
			case string:
				desc.root[k.(string)] = v
				desc.pl = append(desc.pl, k.(string))
			}
		}
	}
	return &desc
}

func (d *described) Extract(key string) Described {
	desc, err := d.ExtractStrict(key)
	if err != nil {
		return &described{
			root: map[string]interface{}{
				key: map[string]interface{}{},
			},
			pl: d.pl, // FIXME: copy d.pl
		}
	}
	return desc
}

func (d *described) ExtractStrict(key string) (Described, error) {
	var extracted described
	if !d.Exists(key) {
		return nil, errors.New("not exist: " + key)
	}

	reflect.Copy(reflect.ValueOf(extracted.pl), reflect.ValueOf(d.pl))
	extracted.root = map[string]interface{}{
		key: d.root[key],
	}
	return &extracted, nil
}

func (d *described) Exists(key string) bool {
	_, ok := d.root[key]
	return ok
}

func (d *described) Decode(key string, dst interface{}) error {
	var to interface{} = d.root
	if key != "" {
		if !d.Exists(key) {
			return errors.New("not found: " + key)
		}
		to = d.root[key]
	}
	return mapstructure.Decode(to, dst)
}

func (d *described) Range(callback func(key string)) {
	for _, p := range d.pl {
		callback(p)
	}
}

func (d *described) Root() map[string]interface{} {
	return d.root
}
