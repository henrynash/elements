// demo of how to create units from raw values and unit strings
package lib

import (
	"context"
	"github.com/antha-lang/antha/antha/anthalib/wtype"
	"github.com/antha-lang/antha/antha/anthalib/wunit"
	"github.com/antha-lang/antha/component"
	"github.com/antha-lang/antha/execute"
	"github.com/antha-lang/antha/inject"
)

// Input parameters for this protocol

// Data which is returned from this protocol

// Physical inputs to this protocol

// Physical outputs from this protocol

func _Lesson4_Units_NewRequirements() {

}

// Actions to perform before protocol Lessonitself
func _Lesson4_Units_NewSetup(_ctx context.Context, _input *Lesson4_Units_NewInput) {

}

// Core process of the protocol: steps to be performed for each input
func _Lesson4_Units_NewSteps(_ctx context.Context, _input *Lesson4_Units_NewInput, _output *Lesson4_Units_NewOutput) {
	// this is importing the NewVolume function from the wunit package
	_output.MyVolume = wunit.NewVolume(_input.MyValue, _input.MyUnit)

	// Exercise: adjust the protocol Lessonto make a concentration instead
}

// Actions to perform after steps block to analyze data
func _Lesson4_Units_NewAnalysis(_ctx context.Context, _input *Lesson4_Units_NewInput, _output *Lesson4_Units_NewOutput) {

}

func _Lesson4_Units_NewValidation(_ctx context.Context, _input *Lesson4_Units_NewInput, _output *Lesson4_Units_NewOutput) {

}
func _Lesson4_Units_NewRun(_ctx context.Context, input *Lesson4_Units_NewInput) *Lesson4_Units_NewOutput {
	output := &Lesson4_Units_NewOutput{}
	_Lesson4_Units_NewSetup(_ctx, input)
	_Lesson4_Units_NewSteps(_ctx, input, output)
	_Lesson4_Units_NewAnalysis(_ctx, input, output)
	_Lesson4_Units_NewValidation(_ctx, input, output)
	return output
}

func Lesson4_Units_NewRunSteps(_ctx context.Context, input *Lesson4_Units_NewInput) *Lesson4_Units_NewSOutput {
	soutput := &Lesson4_Units_NewSOutput{}
	output := _Lesson4_Units_NewRun(_ctx, input)
	if err := inject.AssignSome(output, &soutput.Data); err != nil {
		panic(err)
	}
	if err := inject.AssignSome(output, &soutput.Outputs); err != nil {
		panic(err)
	}
	return soutput
}

func Lesson4_Units_NewNew() interface{} {
	return &Lesson4_Units_NewElement{
		inject.CheckedRunner{
			RunFunc: func(_ctx context.Context, value inject.Value) (inject.Value, error) {
				input := &Lesson4_Units_NewInput{}
				if err := inject.Assign(value, input); err != nil {
					return nil, err
				}
				output := _Lesson4_Units_NewRun(_ctx, input)
				return inject.MakeValue(output), nil
			},
			In:  &Lesson4_Units_NewInput{},
			Out: &Lesson4_Units_NewOutput{},
		},
	}
}

var (
	_ = execute.MixInto
	_ = wtype.FALSE
	_ = wunit.Make_units
)

type Lesson4_Units_NewElement struct {
	inject.CheckedRunner
}

type Lesson4_Units_NewInput struct {
	MyUnit  string
	MyValue float64
}

type Lesson4_Units_NewOutput struct {
	MyVolume wunit.Volume
}

type Lesson4_Units_NewSOutput struct {
	Data struct {
		MyVolume wunit.Volume
	}
	Outputs struct {
	}
}

func init() {
	if err := addComponent(component.Component{Name: "Lesson4_Units_New",
		Constructor: Lesson4_Units_NewNew,
		Desc: component.ComponentDesc{
			Desc: "demo of how to create units from raw values and unit strings\n",
			Path: "src/github.com/antha-lang/elements/AnthaAcademy/Lesson4_Units/A_units_new.an",
			Params: []component.ParamDesc{
				{Name: "MyUnit", Desc: "", Kind: "Parameters"},
				{Name: "MyValue", Desc: "", Kind: "Parameters"},
				{Name: "MyVolume", Desc: "", Kind: "Data"},
			},
		},
	}); err != nil {
		panic(err)
	}
}
