package api

import "github.com/Yesterday17/pug/utils/state"

// Preprocessor is executed before worker works
// It modifies user input, and create the initial State of a work
// If it meets an error, the worker will not work
type Preprocessor interface {
	// Match determines whether current input meets rule now
	// If true is returned, the Worker will select it as Preprocessor, ignoring other choices
	Match(input string) bool

	// Execute tries to build up the initial state of a work
	Execute(input string) (state.State, error)
}
