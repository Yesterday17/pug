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

type biliWrapper struct {
	Code    int           `json:"code"`
	Data    webArchivePre `json:"data"`
	Message string        `json:"message"`
	Ttl     int           `json:"ttl"`
}

type webArchivePre struct {
	// Ignored some useless entries
	// Activities interface{}    `json:"activities"`
	// Fav        interface{}    `json:"fav"`
	// MyInfo     interface{}    `json:"myinfo"`
	TypeList  []typeListItem `json:"typelist"`
	VideoJam  videoJam       `json:"videojam"`
	Watermark watermark      `json:"watermark"`
}

type typeListItem struct {
	Id            int            `json:"id"`
	Parent        int            `json:"parent"`
	Name          string         `json:"name"`
	Description   string         `json:"description"`
	CopyRight     int            `json:"copy_right"`
	IntroCopy     string         `json:"intro_copy"`
	IntroOriginal string         `json:"intro_original"`
	Notice        string         `json:"notice"`
	Show          bool           `json:"show"`
	Children      []typeListItem `json:"children"`
}

type videoJam struct {
	Level   int    `json:"level"`
	State   string `json:"state"`
	Comment string `json:"comment"`
}

type watermark struct {
	Ctime    string `json:"ctime"`
	Id       int    `json:"id"`
	Info     string `json:"info"`
	Md5      string `json:"md5"`
	Mid      int    `json:"mid"`
	Mtime    string `json:"mtime"`
	Position int    `json:"position"`
	State    int    `json:"state"`
	Tip      string `json:"tip"`
	Type     int    `json:"type"`
	Uname    string `json:"uname"`
	Url      string `json:"url"`
}

// https://member.bilibili.com/x/web/archive/pre
