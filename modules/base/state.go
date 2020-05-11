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

package base

import (
	"errors"
	"sync"

	"github.com/Yesterday17/pug/api"
)

var NotFoundError = errors.New("no content found")

func NewState() api.State {
	return &state{
		state: &sync.Map{},
	}
}

type state struct {
	state *sync.Map
}

func (s *state) Has(key string) bool {
	_, ok := s.state.Load(key)
	return ok
}

func (s *state) Get(key string) (interface{}, error) {
	val, ok := s.state.Load(key)
	if !ok {
		return nil, NotFoundError
	}
	return val, nil
}

func (s *state) GetInt(key string) (int, error) {
	val, err := s.Get(key)
	return val.(int), err
}

func (s *state) GetBool(key string) (bool, error) {
	val, err := s.Get(key)
	return val.(bool), err
}

func (s *state) GetString(key string) (string, error) {
	val, err := s.Get(key)
	return val.(string), err
}

func (s *state) GetFloat(key string) (float32, error) {
	val, err := s.Get(key)
	return val.(float32), err
}

func (s *state) Set(key string, value interface{}) {
	s.state.Store(key, value)
}

func (s *state) Delete(key string) {
	s.state.Delete(key)
}

func (s *state) Range(f func(key string, value interface{}) bool) {
	s.state.Range(func(key, value interface{}) bool {
		return f(key.(string), value)
	})
}
