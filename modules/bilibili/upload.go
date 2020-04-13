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
	"errors"
	"fmt"
	"strings"

	"github.com/Yesterday17/pug/api"
	"github.com/Yesterday17/pug/utils/log"
	"github.com/Yesterday17/pug/utils/net"
)

type uploadPipe struct {
	route route
}

func newUploadPipe(m map[string]interface{}) (api.Pipe, api.PipeConstructorError) {
	return nil, 0
}

func (u *uploadPipe) Validate() map[string]interface{} {
	return map[string]interface{}{}
}

func (u *uploadPipe) Execute(work api.State) error {
	u.route = selectRoute()

	video, _ := work.GetString("video")
	v, err := NewVideo(video)
	if err != nil {
		return err
	}

	err = u.PreUpload(v)
	if err != nil {
		return err
	}

	err = u.UploadsPost(v)
	if err != nil {
		return err
	}

	v.SplitChunks()
	u.EmitUpload(v)
	u.AfterUpload(v)

	u.Submit([]*Video{v})
	return nil
}

type route struct {
	os      string
	query   string
	url     string
	profile string
}

var (
	upos = route{
		os:      "upos",
		query:   "os=upos&upcdn=ws",
		url:     "https://upos-sz-upcdnws.acgvideo.com/OK",
		profile: "ugcupos/bup", // if interactive it should be ugcupos/iv
	}
	kodo = route{
		os:      "kodo",
		query:   "os=kodo&bucket=bvcupcdnkodobm",
		url:     "https://up-na0.qbox.me/crossdomain.xml",
		profile: "ugcupos/bupfetch",
	}
)

func selectRoute() route {
	// TODO: Select route after ping
	return upos
}
func (u *uploadPipe) SendChunk(v *Video, c chunk) error {
	url := strings.ReplaceAll(v.UposUri, "upos:\\/\\/", v.EndPoint[2:]) +
		"/?" +
		"partNumber=" + (c.index + 1).String() +
		"&uploadId=" + v.UploadID +
		"&chunk=" + c.index.String() +
		"&chunks=" + c.total.String() +
		"&size=" + c.size.String() +
		"&start=" + c.start.String() +
		"&end=" + c.end.String() +
		"&total=" + c.totalSize.String()

	body, err := net.PutBody(url, net.Headers{
		"X-Upos-Auth": v.Auth,
	}, nil)
	if err != nil {
		return err
	} else if string(body) != "MULTIPART_PUT_SUCCESS" {
		return errors.New(string(body))
	} else {
		return nil
	}
}

func (u *uploadPipe) EmitUpload(v *Video) {
	for _, c := range v.Chunks {
		err := u.SendChunk(v, c)
		if err != nil {
			log.Error(err.Error())
		}
	}
}

func (u *uploadPipe) AfterUpload(v *Video) {
	url := net.BuildUrl(strings.ReplaceAll(v.UposUri, "upos:\\/\\/", v.EndPoint[2:]), true, "", map[string]string{
		"output":   "json",
		"name":     v.File.Name(),
		"profile":  u.route.profile,
		"uploadId": v.UploadID,
		"biz_id":   v.BizID.String(),
	})

	// Build Payload
	payload := `{"parts": [`
	for i := range v.Chunks {
		if i != 0 {
			payload += ","
		}
		payload += fmt.Sprintf(`{"partNumber":%d,"eTag":"etag"}`, i)
	}
	payload += "]}"

	json, err := net.PostJSON(url, map[string]string{"X-Upos-Auth": v.Auth}, strings.NewReader(payload))
	if err != nil {
		//
	}
	_ = json
}
