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
	"os"

	"github.com/Yesterday17/pug/api"
	"github.com/Yesterday17/pug/modules"
	"github.com/Yesterday17/pug/pugd"
	"github.com/Yesterday17/pug/utils/arg"
	"github.com/Yesterday17/pug/utils/log"
)

func appendToPipeline(pl api.Pipeline, name string, params map[string]interface{}) error {
	module, err := modules.NewModule(name, params)
	if err != nil {
		return err
	}
	pl.Append(module)
	return nil
}

func main() {
	pl, err := api.NewLinearPipeline()
	if err != nil {
		panic(err)
	}
	defer pl.TempDir().Clear()

	if len(os.Args) < 2 {
		// TODO: Show Usage
		return
	}

	start := os.Args[1]
	if start == "-d" || start == "daemon" || start == "--daemon" {
		log.Info("Launching pugd...")
		pugd.Main(map[string]interface{}{})
		return
	}

	// no module specified
	if len(os.Args) < 3 {
		// TODO: Show Usage
		return
	}

	ps := arg.ParseArgs(os.Args[2:])
	if ps == nil {
		// TODO: Show Usage
		return
	}

	file, ok := ps[len(ps)-1]["file"]
	if ok {
		// TODO: load from file
		_ = file
	} else {
		// @Deprecated
		for index := 0; index < len(ps)-1; index++ {
			params := ps[index]
			name := params["module"].(string)
			err := appendToPipeline(pl, name, params)
			if err != nil {
				log.Fatal(err.Error())
			}
		}
	}

	pl.RunWith(start)
}
