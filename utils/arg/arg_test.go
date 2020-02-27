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

package arg

import "testing"

func TestParseArgs(t *testing.T) {
	result := ParseArgs([]string{
		"-c", "config.yaml",
		"-p", "test", "-arg1", "-arg2", "aaa", "-arg3=gg",
		"-p=with-space", "-cmd", "[", "echo", "1", "2", "]",
		"-p=bash", "-cmd=2333",
	})

	if len(result) != 3+1 {
		t.Error("Wrong pipe number!")
		return
	}

	// First Pipe
	if result[0]["module"] != "test" {
		t.Error()
	}
	if result[0]["arg1"] != true {
		t.Error()
	}
	if result[0]["arg2"] != "aaa" {
		t.Error()
	}
	if result[0]["arg3"] != "gg" {
		t.Error()
	}

	// Second Pipe
	if result[1]["module"] != "with-space" {
		t.Error()
	}
	if result[1]["cmd"] != "echo 1 2" {
		t.Error()
	}

	// Last Pipe
	if result[2]["module"] != "bash" {
		t.Error()
	}
	if result[2]["cmd"] != "2333" {
		t.Error()
	}

	// Error Test 1
	result = ParseArgs([]string{
		"p", "test",
	})
	if result != nil {
		t.Error()
	}

	// Error Test 2
	result = ParseArgs([]string{
		"-p", "test", "gg",
	})
	if result != nil {
		t.Error()
	}

	// Error Test 3
	result = ParseArgs(nil)
	if result != nil {
		t.Error()
	}

	// Error Test 4
	result = ParseArgs([]string{})
	if result != nil {
		t.Error()
	}
}
