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

import (
	"container/list"
	"github.com/Yesterday17/pug/utils/log"
	"github.com/Yesterday17/pug/utils/temp"
)

type linearPipeline struct {
	line *list.List
	tmp  temp.Dir
}

func NewLinearPipeline() (Pipeline, error) {
	tmp, err := temp.NewDir()
	if err != nil {
		return &linearPipeline{}, err
	}

	return &linearPipeline{
		line: list.New(),
		tmp:  tmp,
	}, nil
}

func (l *linearPipeline) Append(p ...Pipe) {
	for _, pipe := range p {
		l.line.PushBack(pipe)
	}
}

func (l *linearPipeline) Run(p Pipe) {
	var prev = p
	for l.line.Len() > 0 {
		if prev.Status() == PipeError {
			log.Fatalf("Pipeline has met an error in pipe %s, program terminated.\n", prev.(Module).Name())
			break
		}

		current := l.line.Front()
		l.line.Remove(current)

		p := (current.Value).(Pipe)
		if p.Type() == EndpointPipe {
			p := p.(EndPointPipe)
			if err := p.PipeOut(prev, l); err != nil {
				log.Fatalf("Module %s has met an error: %s\n", p.(Module).Name(), err.Error())
			}
			break
		} else {
			p.Do(prev, l)
			prev = p
		}
	}
	if prev.Status() == PipeError {
		log.Fatalf("Pipeline has met an error in the last pipe %s, procedure may be unsuccessful.\n", prev.(interface{}).(Module).Name())
	}
}

func (l *linearPipeline) RunWith(start string) {
	l.Run(&BasePipe{
		PStatus:   PipeSuccess,
		Metadata:  Metadata{Link: start},
		MediaData: Media{},
	})
}

func (l *linearPipeline) TempDir() temp.Dir {
	return l.tmp
}
