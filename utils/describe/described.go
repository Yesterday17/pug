package describe

import (
	"errors"
	"reflect"

	"github.com/mitchellh/mapstructure"
)

type Described interface {
	Exists(key string) bool
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
		"key": d.root[key],
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
