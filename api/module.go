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

type Module interface {
	// Name should be the UNIQUE name of a module
	// It is used to identify different modules
	// It MUST NOT contain ':' or '/' character
	Name() string

	// Description is used to describe what the whole module does
	Description() string

	// Author shows the author(s) of a module
	Author() []string

	// Usage tells users the basic usage of the whole module
	// And usage of all pipes in the module
	Usage() string

	// Preprocessor is used to preprocess input data and generate initial state
	Preprocessor() Preprocessor

	// Pipe is used to get PipeConstructors by Pipe Id
	// Pipe Id only need to be unique in module, and should not contain strange characters
	// If the constructor is not found, it returns nil PipeConstructor and false
	Pipe(pid string) (PipeConstructor, bool)
}
