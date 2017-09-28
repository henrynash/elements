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

func _Lesson4_Units_ConcentrationRequirements() {

}

// Actions to perform before protocol Lessonitself
func _Lesson4_Units_ConcentrationSetup(_ctx context.Context, _input *Lesson4_Units_ConcentrationInput) {

}

// Core process of the protocol: steps to be performed for each input
func _Lesson4_Units_ConcentrationSteps(_ctx context.Context, _input *Lesson4_Units_ConcentrationInput, _output *Lesson4_Units_ConcentrationOutput) {

	_output.ConcinMperL = _input.MyConc.MolPerL(_input.MolecularWeight)
	_output.ConcinGperL = _input.MyConc.GramPerL(_input.MolecularWeight)

}

// Actions to perform after steps block to analyze data
func _Lesson4_Units_ConcentrationAnalysis(_ctx context.Context, _input *Lesson4_Units_ConcentrationInput, _output *Lesson4_Units_ConcentrationOutput) {

}

func _Lesson4_Units_ConcentrationValidation(_ctx context.Context, _input *Lesson4_Units_ConcentrationInput, _output *Lesson4_Units_ConcentrationOutput) {

}
func _Lesson4_Units_ConcentrationRun(_ctx context.Context, input *Lesson4_Units_ConcentrationInput) *Lesson4_Units_ConcentrationOutput {
	output := &Lesson4_Units_ConcentrationOutput{}
	_Lesson4_Units_ConcentrationSetup(_ctx, input)
	_Lesson4_Units_ConcentrationSteps(_ctx, input, output)
	_Lesson4_Units_ConcentrationAnalysis(_ctx, input, output)
	_Lesson4_Units_ConcentrationValidation(_ctx, input, output)
	return output
}

func Lesson4_Units_ConcentrationRunSteps(_ctx context.Context, input *Lesson4_Units_ConcentrationInput) *Lesson4_Units_ConcentrationSOutput {
	soutput := &Lesson4_Units_ConcentrationSOutput{}
	output := _Lesson4_Units_ConcentrationRun(_ctx, input)
	if err := inject.AssignSome(output, &soutput.Data); err != nil {
		panic(err)
	}
	if err := inject.AssignSome(output, &soutput.Outputs); err != nil {
		panic(err)
	}
	return soutput
}

func Lesson4_Units_ConcentrationNew() interface{} {
	return &Lesson4_Units_ConcentrationElement{
		inject.CheckedRunner{
			RunFunc: func(_ctx context.Context, value inject.Value) (inject.Value, error) {
				input := &Lesson4_Units_ConcentrationInput{}
				if err := inject.Assign(value, input); err != nil {
					return nil, err
				}
				output := _Lesson4_Units_ConcentrationRun(_ctx, input)
				return inject.MakeValue(output), nil
			},
			In:  &Lesson4_Units_ConcentrationInput{},
			Out: &Lesson4_Units_ConcentrationOutput{},
		},
	}
}

var (
	_ = execute.MixInto
	_ = wtype.FALSE
	_ = wunit.Make_units
)

type Lesson4_Units_ConcentrationElement struct {
	inject.CheckedRunner
}

type Lesson4_Units_ConcentrationInput struct {
	MolecularWeight float64
	MyConc          wunit.Concentration
}

type Lesson4_Units_ConcentrationOutput struct {
	ConcinGperL wunit.Concentration
	ConcinMperL wunit.Concentration
}

type Lesson4_Units_ConcentrationSOutput struct {
	Data struct {
		ConcinGperL wunit.Concentration
		ConcinMperL wunit.Concentration
	}
	Outputs struct {
	}
}

func init() {
	if err := addComponent(component.Component{Name: "Lesson4_Units_Concentration",
		Constructor: Lesson4_Units_ConcentrationNew,
		Desc: component.ComponentDesc{
			Desc: "",
			Path: "src/github.com/antha-lang/elements/AnthaAcademy/Lesson4_Units/D_units_Conc.an",
			Params: []component.ParamDesc{
				{Name: "MolecularWeight", Desc: "", Kind: "Parameters"},
				{Name: "MyConc", Desc: "", Kind: "Parameters"},
				{Name: "ConcinGperL", Desc: "", Kind: "Data"},
				{Name: "ConcinMperL", Desc: "", Kind: "Data"},
			},
		},
	}); err != nil {
		panic(err)
	}
}
