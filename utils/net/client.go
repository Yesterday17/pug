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
	"io"
	"net/http"
	"net/http/cookiejar"

	"github.com/tidwall/gjson"
)

const DefaultUserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) PUG/utils/net-latest Safari/537.36"

type Client struct {
	cookies   http.CookieJar
	UserAgent string
}

func NewClient() *Client {
	jar, _ := cookiejar.New(&cookiejar.Options{})
	return &Client{
		cookies:   jar,
		UserAgent: DefaultUserAgent,
	}
}

func (c *Client) Get(url string, headers Headers) (Response, error) {
	return c.Request("GET", url, headers, nil)
}

func (c *Client) GetBody(url string, headers Headers) ([]byte, error) {
	return c.RequestBody("GET", url, headers, nil)
}

func (c *Client) GetJSON(url string, headers Headers) (gjson.Result, error) {
	return c.RequestJSON("GET", url, headers, nil)
}

func (c *Client) Post(url string, headers Headers, body io.Reader) (Response, error) {
	return c.Request("POST", url, headers, body)
}

func (c *Client) PostBody(url string, headers Headers, body io.Reader) ([]byte, error) {
	return c.RequestBody("POST", url, headers, body)
}

func (c *Client) PostJSON(url string, headers Headers, body io.Reader) (gjson.Result, error) {
	return c.RequestJSON("POST", url, headers, body)
}

func (c *Client) Put(url string, headers Headers, body io.Reader) (Response, error) {
	return c.Request("PUT", url, headers, body)
}

func (c *Client) PutBody(url string, headers Headers, body io.Reader) ([]byte, error) {
	return c.RequestBody("PUT", url, headers, body)
}

func (c *Client) PutJSON(url string, headers Headers, body io.Reader) (gjson.Result, error) {
	return c.RequestJSON("PUT", url, headers, body)
}
