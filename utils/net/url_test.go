/**
MIT License

Copyright (c) 2019-2020 Yesterday17

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package net

import "testing"

type testEntry struct {
	result string
	expect string
}

func TestBuildUrl(t *testing.T) {
	domain := "www.bilibili.com"
	cases := []testEntry{
		{
			BuildUrl(domain, false, "", map[string]string{}),
			"http://www.bilibili.com/",
		},
		{
			BuildUrl(domain, true, "", map[string]string{}),
			"https://www.bilibili.com/",
		},
		{
			BuildUrl(domain, true, "preupload", map[string]string{}),
			"https://www.bilibili.com/preupload",
		},
		{
			BuildUrl(domain, true, "preupload", map[string]string{
				"a": "1 2 3",
				"b": "+-*/",
			}),
			"https://www.bilibili.com/preupload?a=1+2+3&b=%2B-%2A%2F&",
		},
	}

	for i, c := range cases {
		if c.result != c.expect {
			t.Errorf("Expected %s, but got %s.", c.expect, c.result)
		} else {
			t.Logf("Test case #%d passed.", i+1)
		}
	}
}
