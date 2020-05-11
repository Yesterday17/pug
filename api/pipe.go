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

package api

import (
	"reflect"
	"strconv"
)

type PipeBuildError int

const (
	PipeNoError PipeBuildError = iota
	PipeArgumentMissing
	PipeArgumentTypeMismatch
	PipeArgumentInvalid
)

func (pe PipeBuildError) Error() string {
	return strconv.Itoa(int(pe))
}

// PipeBuilder is designed to replace the old PipeConstructor
// It can be used to build pipe with given arguments
// Or determine what arguments are necessary
type PipeBuilder interface {
	// Build works as the original PipeConstructor
	// It is a function used to build up a new Pipe.
	// It returns a Pipe if arguments given are sufficient
	// Or it should returns nil Pipe and an error
	Build(map[string]interface{}) (Pipe, PipeBuildError)

	// Accept can be used to determine whether an argument is necessary
	Accept(key string, t reflect.Kind) bool

	// Must returns a MUST map of arguments
	Must() map[string]reflect.Kind

	// Optional returns an OPTIONAL map of arguments
	Optional() map[string]reflect.Kind
}

// Pipe is the minimal reuse unit in the project.
// Keep it simple, stupid
type Pipe interface {
	// Validate returns a map with string key and reflect.Kind Type
	// If a string begins with '+', it would be ADDED to work state
	// If a string begins with '-', it would be REMOVED from work state
	// If a string begins with '!', it is needed by this pipe
	// If a string begins with '?', it is optional to this pipe
	//
	// If a string begins with '+', then it **should** NOT exist in the previous state
	// THE 'SHOULD' MIGHT BE CHANGED TO MUST IN FURTHER VERSION
	//
	// If a string begins with '-' or '!', it MUST exist in the previous state
	//
	// If the function returns nil, it means type validation SHOULD be skipped
	Validate() map[string]reflect.Kind

	// Execute a pipe
	// Only pipes pass the validation can be executed
	// After execution, the State would be changed as Validate describes
	Execute(work State) error

	// Clone a pipe
	Clone() Pipe
}
