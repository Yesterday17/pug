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
	"io/ioutil"
	"net/http"

	"github.com/tidwall/gjson"
)

func (c *Client) Request(method, url string, headers Headers, body io.Reader) (Response, error) {
	cli := http.Client{
		Jar:     c.cookies,
		Timeout: 0,
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", c.UserAgent)

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	res, err := cli.Do(req)
	if err != nil {
		return nil, err
	}

	c.handleCookies(res)
	return res, nil
}

func (c *Client) RequestBody(method, url string, headers Headers, body io.Reader) ([]byte, error) {
	res, err := c.Request(method, url, headers, body)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (c *Client) RequestJSON(method, url string, headers Headers, body io.Reader) (gjson.Result, error) {
	data, err := c.RequestBody(method, url, headers, body)
	if err != nil {
		return gjson.Result{}, err
	}
	return gjson.Parse(string(data)), nil
}
