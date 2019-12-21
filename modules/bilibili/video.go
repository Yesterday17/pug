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

package bilibili

import (
	"os"
	"strconv"
)

type bigInt int64

func (i bigInt) String() string {
	return strconv.FormatInt(int64(i), 10)
}

type Video struct {
	File   *os.File
	Chunks []chunk

	UposUri   string
	Auth      string
	BizID     bigInt
	ChunkSize bigInt
	Threads   bigInt
	EndPoint  string

	UploadID string
	Key      string
}

func NewVideo(file string) (*Video, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	return &Video{File: f}, nil
}
