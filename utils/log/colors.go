/**
MIT License

Copyright (c) 2017 Ruslan Ianberdin

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

package log

type color string

const (
	// ANSI color escape codes
	Bold        color = "\033[1m"
	Yellow      color = "\033[33m"
	Cyan        color = "\033[36m"
	Gray        color = "\033[90m"
	Red         color = "\033[31m"
	Blue        color = "\033[34m"
	Pink        color = "\033[35m"
	Green       color = "\033[32m"
	LightRed    color = "\033[91m"
	LightGreen  color = "\033[92m"
	LightYellow color = "\033[93m"
	LightBlue   color = "\033[94m"
	LightPink   color = "\033[95m"
	LightCyan   color = "\033[96m"
	White       color = "\033[97m"
	Black       color = "\033[30m"
	endColor    color = "\033[39m" // "reset everything"
)
