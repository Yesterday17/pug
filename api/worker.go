package api

// Worker holds state itself, so it doesn't return anything other than error
type Worker interface {
	// Start accepts a string as input
	// It returns an error if no preprocessor is selected
	Start(input string) error

	// Pause pauses a work at current pipeline
	// Work before current pipe is kept
	Pause() error

	// Cancel destroys the current work and empties the worker
	Cancel() error
}
