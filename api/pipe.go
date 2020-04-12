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

type PipeConstructorError int

const (
	PipeNoError PipeConstructorError = iota
	PipeArgumentMissing
	PipeArgumentTypeMismatch
	PipeArgumentInvalid
)

// PipeConstructor is a function to build up a new Pipe
// It returns a Pipe if arguments given are sufficient
// Or it should returns nil Pipe and an error
type PipeConstructor func(map[string]interface{}) (Pipe, PipeConstructorError)

// Pipe is the minimal reuse unit in the project.
// Keep it simple, stupid
type Pipe interface {
	// Validate returns a map with string key and interface{} value
	// If the string begins with '+', a value Named string[1:] would be ADDED to work state
	// If a string begins with '-', a value Named string[1:] would be REMOVED from work state
	// Else, the value named string in work state would be modified
	//
	// If a string begins with '+', then it **should** NOT exist in the previous state
	// THE 'SHOULD' MIGHT BE CHANGED TO MUST IN FURTHER VERSION
	//
	// If a string begins with '-' or neither those two, it MUST exist in the previous state
	//
	// If the function returns nil, it means type validation SHOULD be skipped
	Validate() map[string]interface{}

	// Execute a pipe
	// Only pipes pass the validation can be executed
	// After execution, the State would be changed as Validate describes
	Execute(work State) error
}
