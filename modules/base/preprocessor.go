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

	"github.com/Yesterday17/pug/api"
)

type preprocessor struct {
	regex *regexp.Regexp

	execute func(env map[string]interface{}, input string) (api.State, error)
}

/**
 * base.Preprocessor is a regular expression based preprocessor
 * It uses a regex string and an execute function as input
 *
 * The function will not throw an error unless the regexp is invalid
 */
func Preprocessor(regex string, exec func(map[string]interface{}, string) (api.State, error)) api.Preprocessor {
	r, err := regexp.Compile(regex)
	if err != nil {
		panic(err)
	}

	return &preprocessor{
		regex:   r,
		execute: exec,
	}
}

func (p *preprocessor) Match(input string) bool {
	return p.regex.MatchString(input)
}

func (p *preprocessor) Execute(env map[string]interface{}, input string) (api.State, error) {
	return p.execute(env, input)
}
