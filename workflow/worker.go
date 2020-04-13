package workflow

import (
	"errors"

	"github.com/Yesterday17/pug/api"
	"github.com/Yesterday17/pug/utils/describe"
)

type worker struct {
	desc    describe.Described
	modules api.ModuleManager

	flow *workerFlowItem

	state api.State
}

type workerFlowItem struct {
	self api.Pipe
	next *workerFlowItem
}

func NewWorker(desc describe.Described, manager api.ModuleManager) (api.Worker, error) {
	w := &worker{desc: desc, modules: manager}

	flow := &workerFlowItem{}
	w.flow = flow
	for _, key := range desc.Workflow() {
		root := desc.Sub(key).Root()
		m, ok := root["module"]
		if !ok {
			return nil, errors.New("module not found")
		}

		p, ok := root["pipe"]
		if !ok {
			return nil, errors.New("pipe not found")
		}

		pc := manager.Pipe(m.(string), p.(string))
		if pc == nil {
			return nil, errors.New("module or pipe not found")
		}

		flow.self, _ = pc(root)
		flow.next = &workerFlowItem{}
		flow = flow.next
	}

	if w.flow.self == nil {
		return nil, errors.New("no workflow constructed")
	}

	return w, nil
}

func (w *worker) Start(input string) error {
	var state api.State
	var err error

	for _, p := range w.modules.Preprocessors() {
		if p.Match(input) {
			state, err = p.Execute(w.desc.Env(), input)
			if err != nil {
				return err
			}
			break
		}
	}

	if state == nil {
		return errors.New("no matching preprocessor found")
	}

	// TODO

	return nil
}

func (w *worker) Pause() error {
	// TODO
	return nil
}

func (w *worker) Cancel() error {
	// TODO
	return nil
}

func (w *worker) Clone() api.Worker {
	// TODO
	return nil
}

func (w *worker) Clean() {
	w.state = nil
}
