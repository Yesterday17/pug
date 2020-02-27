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

package bilibili

import (
	"strings"

	"github.com/Yesterday17/pug/utils/net"
)

func (v *Video) PreUpload(m *Module) error {
	stat, err := v.File.Stat()
	if err != nil {
		return err
	}
	json, err := net.GetJSON(net.BuildUrl("member.bilibili.com", true, "preupload", map[string]string{
		"name":    v.File.Name(),
		"size":    bigInt(stat.Size()).String(),
		"r":       m.Route.os,
		"profile": m.Route.profile,
		"ssl":     "0",
		"version": "2.7.0",
		"build":   "2070000",
	})+m.Route.query, nil)
	if err != nil {
		return err
	}

	v.UposUri = json.Get("upos_uri").String()          // Upos Uri
	v.Auth = json.Get("auth").String()                 // Auth
	v.BizID = bigInt(json.Get("biz_id").Int())         // Biz ID
	v.ChunkSize = bigInt(json.Get("chunk_size").Int()) // Chunk Size
	v.Threads = bigInt(json.Get("threads").Int())      // Threads

	// EndPoint
	if json.Get("endpoints.#").Int() > 0 {
		v.EndPoint = json.Get("endpoints.0").String()
	} else {
		v.EndPoint = json.Get("endpoint").String()
	}
	return nil
}

func (v *Video) UploadsPost() error {
	url := strings.ReplaceAll(v.UposUri, "upos:\\/\\/", v.EndPoint)[2:] // 2: here to remove \/\/
	json, err := net.PostJSON(net.BuildUrl(url, true, "", map[string]string{
		"uploads": "true",
		"output":  "json",
	}), net.Headers{
		"X-Upos-Auth": v.Auth,
	}, nil)
	if err != nil {
		return err
	}

	v.UploadID = json.Get("upload_id").String()
	v.Key = json.Get("key").String()
	return nil
}
