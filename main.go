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

package main

import (
	"github.com/Yesterday17/pug/api"
	"github.com/Yesterday17/pug/modules"
)

func main() {
	pl, err := api.NewLinearPipeline()
	if err != nil {
		panic(err)
	}
	defer pl.TempDir().Clear()

	// TODO: Remove test
	pl.Append(
		modules.Modules["bash"](map[string]interface{}{
			"cmd": "date",
		}).(api.Pipe),
		modules.Modules["bash"](map[string]interface{}{
			"cmd": "echo \"version: $PUG_VERSION\"\necho \"output media: $PUG_OUTPUT_MEDIA\"",
		}).(api.Pipe),
	)
	pl.Run()
}
