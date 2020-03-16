package api

// TODO
type Worker interface {
	Pause() error
	Cancel() error
}
