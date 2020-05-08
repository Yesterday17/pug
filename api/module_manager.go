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

type ModuleManager interface {
	// Exist validates whether a module named mid exists
	Exist(mid string) bool

	// ExistPipe validates whether a module named mid contains pipe named pid
	// It both checks mid and pid
	ExistPipe(mid, pid string) bool

	// Add a module to the manager
	// Then the module can be queried by the manager
	// If the name of this pipe is used, it will return an error
	Add(m Module) error

	// Module gets module named mid
	// If the module does not exist, it returns nil
	Module(mid string) Module

	// Modules get all mid of modules
	Modules() []string

	// Pipe gets constructor of pipe named pid in module named mid
	// If the module or pipe does not exist, it returns nil
	Pipe(mid, pid string) PipeConstructor

	// Preprocessors among all the modules loaded
	Preprocessors() []Preprocessor
}
