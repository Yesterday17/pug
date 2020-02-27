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

package temp

import (
	"encoding/hex"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
)

type Dir string

func NewDir(folder string) (Dir, error) {
	dir, err := ioutil.TempDir(folder, "pug")
	if err != nil {
		return "", err
	}
	return Dir(dir), nil
}

func (d Dir) Clear() {
	_ = os.RemoveAll(string(d))
}

func (d Dir) NewFile(ext string) string {
	randBytes := make([]byte, 16)
	rand.Read(randBytes)
	return filepath.Join(string(d), hex.EncodeToString(randBytes)+ext)
}

func (d Dir) NewContentFile(content string, ext string) (string, error) {
	file := d.NewFile(ext)
	if err := ioutil.WriteFile(file, []byte(content), 0644); err != nil {
		return "", err
	}
	return file, nil
}
