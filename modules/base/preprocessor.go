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
	"regexp"
)

type preprocessor struct {
	regex *regexp.Regexp

	std chan interface{}
}

/**
 * base.Preprocessor is a regular expression based preprocessor
 * It uses a regex string and a interface{} channel as input
 *
 * The interface{} chan will receive a input and needs to generate an output
 * The type of output **must** be either api.State or error
 *
 * The function will not throw an error unless the regexp is invalid
 */
func Preprocessor(regex string, std chan interface{}) (api.Preprocessor, error) {
	r, err := regexp.Compile(regex)
	if err != nil {
		return nil, err
	}

	return &preprocessor{
		regex: r,
		std:   std,
	}, nil
}

func (p *preprocessor) Match(input string) bool {
	return p.regex.MatchString(input)
}

func (p *preprocessor) Execute(input string) (api.State, error) {
	p.std <- input
	out := <-p.std
	switch out.(type) {
	case error:
		return nil, out.(error)
	default:
		return out.(api.State), nil
	}
}
