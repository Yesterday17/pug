package base

import (
	"reflect"

	"github.com/Yesterday17/pug/api"
)

type pipeBuilder struct {
	must     map[string]reflect.Kind
	optional map[string]reflect.Kind
	builder  func(args map[string]interface{}) (api.Pipe, api.PipeBuildError)
}

func NewPipeBuilder(must, optional map[string]reflect.Kind, builder func(args map[string]interface{}) (api.Pipe, api.PipeBuildError)) api.PipeBuilder {
	return &pipeBuilder{
		must:     must,
		optional: optional,
		builder:  builder,
	}
}

func (p *pipeBuilder) Build(args map[string]interface{}) (api.Pipe, api.PipeBuildError) {
	return p.builder(args)
}

func (p *pipeBuilder) Accept(k string, t reflect.Kind) bool {
	if p.must != nil {
		ty, ok := p.must[k]
		if ok && ty == t {
			return true
		}
	}

	if p.optional != nil {
		ty, ok := p.optional[k]
		return ok && ty == t
	}
	return false
}

// TODO: return a clone
func (p *pipeBuilder) Must() map[string]reflect.Kind {
	if p.must == nil {
		return map[string]reflect.Kind{}
	}
	return p.must
}

// TODO: return a clone
func (p *pipeBuilder) Optional() map[string]reflect.Kind {
	if p.optional == nil {
		return map[string]reflect.Kind{}
	}
	return p.optional
}
