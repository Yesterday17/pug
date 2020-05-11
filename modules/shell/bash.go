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

package shell

import (
	"os"
	"os/exec"
	"reflect"

	"github.com/Yesterday17/pug/api"
	"github.com/Yesterday17/pug/modules/base"
	"github.com/Yesterday17/pug/utils/log"
	"github.com/Yesterday17/pug/utils/types"
)

var bashPipeBuilder = base.NewPipeBuilder(map[string]reflect.Kind{"cmd": reflect.String}, nil, newBashPipe)

func newBashPipe(m map[string]interface{}) (api.Pipe, api.PipeBuildError) {
	cmd, err := types.String(m, "cmd")
	if err != api.PipeNoError {
		return nil, err
	}

	if cmd == "" {
		return nil, api.PipeArgumentInvalid
	}
	return &bashPipe{command: cmd}, api.PipeNoError
}

type bashPipe struct {
	command string
}

func (b *bashPipe) Validate() map[string]reflect.Kind {
	return nil
}

func (b *bashPipe) Execute(work api.State) error {
	cmd := exec.Command("bash", "-c", b.command)

	// TODO: Use logger defined in API
	cmd.Stdout = log.DefaultLogger.WrappedLogWriter
	cmd.Stderr = log.DefaultLogger.WrappedErrorWriter

	cmd.Env = append(os.Environ(),
		"PUG_API_VERSION="+api.VERSION,
	)

	err := cmd.Start()
	if err != nil {
		return err
	}

	err = cmd.Wait()
	if err != nil {
		return err
	}

	return nil
}

func (b *bashPipe) Clone() api.Pipe {
	return &bashPipe{command: b.command}
}
