package workflow

import (
	"errors"
	"reflect"

	"github.com/Yesterday17/pug/api"
	"github.com/Yesterday17/pug/modules/base"
	"github.com/Yesterday17/pug/utils/describe"
)

type worker struct {
	desc    describe.Described
	modules api.ModuleManager

	flow []api.Pipe

	state api.State
}

func NewWorker(desc describe.Described, manager api.ModuleManager) (api.Worker, error) {
	w := &worker{desc: desc, modules: manager, flow: []api.Pipe{}}

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

		pipe, err := pc(root)
		if err != api.PipeNoError {
			return nil, errors.New("pipe construct error")
		}
		w.flow = append(w.flow, pipe)
	}

	if len(w.flow) == 0 {
		return nil, errors.New("no workflow constructed")
	}

	// Workflow type validation
	testState := base.NewState()
	for _, p := range w.flow {
		v := p.Validate()
		if v == nil {
			continue
		}

		for k, t := range v {
			// k: key, t: type
			if len(k) <= 0 {
				return nil, errors.New("empty validate string")
			}

			switch k[0] {
			case '+':
				testState.Set(k[1:], t)
				fallthrough
			case '?':
				continue
			case '-', '!':
				if !testState.Has(k[1:]) {
					return nil, errors.New("invalid state") // FIXME: description
				}

				inState, _ := testState.Get(k)
				if reflect.TypeOf(t).Kind() != reflect.TypeOf(inState).Kind() {
					return nil, errors.New("state type mismatch") // FIXME: description
				}

				if k[0] == '-' {
					testState.Delete(k[1:])
				}
			default:
				return nil, errors.New("invalid control character")
			}
		}
	}

	return w, nil
}

func (w *worker) Start(input string) error {
	var err error

	// Find matching preprocessor and create initial state
	for _, p := range w.modules.Preprocessors() {
		if p.Match(input) {
			w.state, err = p.Execute(w.desc.Env(), input)
			if err != nil {
				return err
			}
			break
		}
	}

	if w.state == nil {
		return errors.New("no matching preprocessor found")
	}

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
