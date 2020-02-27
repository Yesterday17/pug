/**
MIT License

Copyright (c) 2019 Yesterday17

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

package conf

import (
	"fmt"
	"testing"

	"github.com/Yesterday17/pug/api"
)

func TestSerialize(t *testing.T) {
	result := Serialize(api.Metadata{
		Title:       "123",
		Author:      "456",
		Description: "78910",
		Cover:       "111213",
		Link:        "123",
		Short:       "456",
		From:        "78910",
		ReleaseTime: "111213",
	})

	if result != "title=123\nauthor=456\ndesc=78910\ncover=111213\nlink=123\nshort=456\nFrom=78910\nrelease=111213\n" {
		t.Error(result)
	}
}

func TestDeserialize(t *testing.T) {
	var s api.Metadata
	err := Deserialize("", s)
	if err == nil {
		t.Error("Passing non-pointer value should throw an error!")
	}

	err = Deserialize("title=123\nauthor=456\ndesc=78910\ncover=111213\nlink=123\nshort=456\nFrom=78910\nrelease=111213\n", &s)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(s)
}
