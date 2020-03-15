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

package main

import (
	"flag"

	"github.com/Yesterday17/pug/api"
	"github.com/Yesterday17/pug/modules"
	"github.com/Yesterday17/pug/utils/describe"
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
	var pl api.Pipeline
	var daemon bool
	var config, url string

	flag.StringVar(&config, "config", "", "")
	flag.StringVar(&url, "url", "", "")
	flag.BoolVar(&daemon, "daemon", false, "")

	desc, err := describe.Load(config)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	desc.Range(func(key string) {
		root := desc.Sub(key).Root()
		name, ok := root["module"]
		if !ok {
			log.Fatal("module not found")
			return
		}
		err := appendToPipeline(pl, name.(string), root)
		if err != nil {
			log.Fatal(err.Error())
			return
		}
	})

	pl.RunWith(url)
}
