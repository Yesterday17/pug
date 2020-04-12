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
	// Whether current key exists in this Described
	Exists(key string) bool

	// Sub tries to get a sub Described
	Sub(key string) Described

	// Extract generates a sub-description with strict rules
	Extract(key string) (Described, error)
	// ExtractWeak generates a sub description, or return an empty description after meeting an error
	ExtractWeak(key string) Described

	// Decode uses mapstructure to extract a description to a struct
	Decode(key string, dst interface{}) error

	Root() map[string]interface{}
	Env() map[string]interface{}
	Workflow() []string
}

type described struct {
	// root is the root node of current yaml file
	root map[string]interface{}
	// env describes pre-preprocessor states
	env map[string]interface{}
	// workflow is an array of pipeline strings
	workflow []string
}

func (d *described) Sub(key string) Described {
	var desc = described{
		root:     map[string]interface{}{},
		env:      d.env,
		workflow: []string{},
	}

	ex := d.ExtractWeak(key)
	switch ex.Root()[key].(type) {
	case map[interface{}]interface{}:
		m := ex.Root()[key].(map[interface{}]interface{})
		for k, v := range m {
			switch k.(type) {
			case string:
				desc.root[k.(string)] = v
				desc.workflow = append(desc.workflow, k.(string))
			}
		}
	}
	return &desc
}

func (d *described) ExtractWeak(key string) Described {
	desc, err := d.Extract(key)
	if err != nil {
		return &described{
			root:     map[string]interface{}{},
			env:      d.env,
			workflow: d.workflow,
		}
	}
	return desc
}

func (d *described) Extract(key string) (Described, error) {
	var extracted described
	if !d.Exists(key) {
		return nil, errors.New("not exist: " + key)
	}

	reflect.Copy(reflect.ValueOf(extracted.workflow), reflect.ValueOf(d.workflow))
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
	return mapstructure.WeakDecode(to, dst)
}

func (d *described) Root() map[string]interface{} {
	return d.root
}

func (d *described) Env() map[string]interface{} {
	return d.env
}

func (d *described) Workflow() []string {
	return d.workflow
}
