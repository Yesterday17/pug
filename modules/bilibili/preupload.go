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
	"github.com/Yesterday17/pug/utils/net"
	"github.com/tidwall/gjson"
	"os"
	"strconv"
	"strings"
)

func (b *bilibili) PreUpload(file os.File) (gjson.Result, error) {
	stat, err := file.Stat()
	if err != nil {
		return gjson.Result{}, err
	}
	json, err := net.GetJSON(net.BuildUrl("member.bilibili.com", true, "preupload", map[string]string{
		"name":    file.Name(),
		"size":    strconv.FormatInt(stat.Size(), 10),
		"r":       b.Route.os,
		"profile": b.Route.profile,
		"ssl":     "0",
		"version": "2.7.0",
		"build":   "2070000",
	})+b.Route.query, nil)
	if err != nil {
		return gjson.Result{}, err
	}
	return json, nil
}

func (b *bilibili) UploadsPost(beforeUpload gjson.Result) (gjson.Result, error) {
	var endPoint string
	if beforeUpload.Get("endpoints.#").Int() > 0 {
		endPoint = beforeUpload.Get("endpoints.0").String()
	} else {
		endPoint = beforeUpload.Get("endpoint").String()
	}

	url := strings.ReplaceAll(beforeUpload.Get("upos_uri").String(), "upos:\\/\\/", endPoint)[4:] // 2: here to remove \/\/
	uploadsPost, err := net.PostJSON(net.BuildUrl(url, true, "", map[string]string{
		"uploads": "true",
		"output":  "json",
	}), net.Headers{
		"X-Upos-Auth": beforeUpload.Get("auth").String(),
	}, nil)
	if err != nil {
		return gjson.Result{}, err
	}
	return uploadsPost, nil
}
