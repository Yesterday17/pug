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

package api

type PipeStatus int
type PipeType int

const (
	FilterPipe   PipeType = 0
	SourcePipe   PipeType = 1
	EndpointPipe PipeType = 2

	PipeWaiting PipeStatus = 0
	PipeWorking PipeStatus = 1
	PipeSuccess PipeStatus = 2
	PipeError   PipeStatus = 3
)

type Pipe interface {
	Type() PipeType
	Status() PipeStatus

	Meta() Metadata
	Media() Media

	Auth() error
	Do(prev Pipe, pl Pipeline)
}

type EndPointPipe interface {
	Pipe
	PipeOut(prev Pipe, pl Pipeline) error
}
