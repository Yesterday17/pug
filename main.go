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

	"github.com/Yesterday17/pug/modules"
	"github.com/Yesterday17/pug/utils/describe"
	"github.com/Yesterday17/pug/utils/log"
	"github.com/Yesterday17/pug/workflow"
)

func main() {
	var config, url string

	flag.StringVar(&config, "config", "", "")
	flag.StringVar(&url, "url", "", "")
	flag.Parse()

	desc, err := describe.Load(config)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	worker, err := workflow.NewWorker(desc, modules.Manager)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	err = worker.Start(url)
	if err != nil {
		log.Fatal(err.Error())
	}
}
