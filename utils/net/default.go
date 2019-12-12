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

package net

import (
	"github.com/tidwall/gjson"
	"io"
)

var DefaultClient = NewClient()

func Get(url string, headers Headers) (Response, error) {
	return DefaultClient.Get(url, headers)
}

func GetBody(url string, headers Headers) ([]byte, error) {
	return DefaultClient.GetBody(url, headers)
}

func GetJSON(url string, headers Headers) (gjson.Result, error) {
	return DefaultClient.GetJSON(url, headers)
}

func Post(url string, headers Headers, body io.Reader) (Response, error) {
	return DefaultClient.Post(url, headers, body)
}

func PostBody(url string, headers Headers, body io.Reader) ([]byte, error) {
	return DefaultClient.PostBody(url, headers, body)
}

func PostJSON(url string, headers Headers, body io.Reader) (gjson.Result, error) {
	return DefaultClient.PostJSON(url, headers, body)
}

func Put(url string, headers Headers, body io.Reader) (Response, error) {
	return DefaultClient.Put(url, headers, body)
}

func PutBody(url string, headers Headers, body io.Reader) ([]byte, error) {
	return DefaultClient.PutBody(url, headers, body)
}

func PutJSON(url string, headers Headers, body io.Reader) (gjson.Result, error) {
	return DefaultClient.PutJSON(url, headers, body)
}
