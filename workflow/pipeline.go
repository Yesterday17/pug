package workflow

import (
	"reflect"

	"github.com/Yesterday17/pug/api"
)

type pipeline struct {
	pipes    []api.Pipe
	validate map[string]reflect.Kind
}

func (ppl *pipeline) Validate() map[string]reflect.Kind {
	return ppl.validate
}

func (ppl *pipeline) Execute(work api.State) error {
	for _, p := range ppl.pipes {
		err := p.Execute(work)
		if err != nil {
			return err
		}
	}
	return nil
}

func (ppl *pipeline) Clone() api.Pipe {
	var ps []api.Pipe
	for _, p := range ppl.pipes {
		ps = append(ps, p.Clone())
	}
	return &pipeline{pipes: ps}
}

func (ppl *pipeline) Append(p ...api.Pipe) {
	ppl.pipes = append(ppl.pipes, p...)
	ppl.validate = generateValidation(ppl.validate, p...)
}

func NewPipeline(p ...api.Pipe) api.Pipeline {
	validate := map[string]reflect.Kind{}

	return &pipeline{pipes: p, validate: generateValidation(validate, p...)}
}

func generateValidation(validate map[string]reflect.Kind, pipes ...api.Pipe) map[string]reflect.Kind {
	for _, p := range pipes {
		v := p.Validate()
		if v == nil {
			continue
		}

		for k, t := range v {
			switch k[0] {
			case '-':
				if _, ok := validate["+"+k[1:]]; ok {
					delete(validate, "+"+k[1:])
					continue
				}
				fallthrough
			case '+', '?', '!':
				validate[k] = t
			}
		}
	}
	return validate
}
