// Aliquot a solution into a specified plate.
// optionally premix the solution before aliquoting
package lib

import (
	"context"
	"github.com/antha-lang/antha/antha/anthalib/mixer"
	"github.com/antha-lang/antha/antha/anthalib/wtype"
	"github.com/antha-lang/antha/antha/anthalib/wunit"
	"github.com/antha-lang/antha/component"
	"github.com/antha-lang/antha/execute"
	"github.com/antha-lang/antha/inject"
)

// Input parameters for this protocol (data)

// Data which is returned from this protocol, and data types

// Physical Inputs to this protocol with types

// Physical outputs from this protocol with types

func _SimpleAliquotMixRequirements() {

}

// Conditions to run on startup
func _SimpleAliquotMixSetup(_ctx context.Context, _input *SimpleAliquotMixInput) {

}

// The core process for this protocol, with the steps to be performed
// for every input
func _SimpleAliquotMixSteps(_ctx context.Context, _input *SimpleAliquotMixInput, _output *SimpleAliquotMixOutput) {

	for i := 0; i < _input.NumberOfAliquots; i++ {
		sampleA := mixer.Sample(_input.SampleName, _input.AliquotVolume)
		aliquot := execute.Mix(_ctx, sampleA)
		_output.Aliquots = append(_output.Aliquots, aliquot)
	}

}

// Run after controls and a steps block are completed to
// post process any data and provide downstream results
func _SimpleAliquotMixAnalysis(_ctx context.Context, _input *SimpleAliquotMixInput, _output *SimpleAliquotMixOutput) {
}

// A block of tests to perform to validate that the sample was processed
//correctly. Optionally, destructive tests can be performed to validate
//results on a dipstick basis
func _SimpleAliquotMixValidation(_ctx context.Context, _input *SimpleAliquotMixInput, _output *SimpleAliquotMixOutput) {

}
func _SimpleAliquotMixRun(_ctx context.Context, input *SimpleAliquotMixInput) *SimpleAliquotMixOutput {
	output := &SimpleAliquotMixOutput{}
	_SimpleAliquotMixSetup(_ctx, input)
	_SimpleAliquotMixSteps(_ctx, input, output)
	_SimpleAliquotMixAnalysis(_ctx, input, output)
	_SimpleAliquotMixValidation(_ctx, input, output)
	return output
}

func SimpleAliquotMixRunSteps(_ctx context.Context, input *SimpleAliquotMixInput) *SimpleAliquotMixSOutput {
	soutput := &SimpleAliquotMixSOutput{}
	output := _SimpleAliquotMixRun(_ctx, input)
	if err := inject.AssignSome(output, &soutput.Data); err != nil {
		panic(err)
	}
	if err := inject.AssignSome(output, &soutput.Outputs); err != nil {
		panic(err)
	}
	return soutput
}

func SimpleAliquotMixNew() interface{} {
	return &SimpleAliquotMixElement{
		inject.CheckedRunner{
			RunFunc: func(_ctx context.Context, value inject.Value) (inject.Value, error) {
				input := &SimpleAliquotMixInput{}
				if err := inject.Assign(value, input); err != nil {
					return nil, err
				}
				output := _SimpleAliquotMixRun(_ctx, input)
				return inject.MakeValue(output), nil
			},
			In:  &SimpleAliquotMixInput{},
			Out: &SimpleAliquotMixOutput{},
		},
	}
}

var (
	_ = execute.MixInto
	_ = wtype.FALSE
	_ = wunit.Make_units
)

type SimpleAliquotMixElement struct {
	inject.CheckedRunner
}

type SimpleAliquotMixInput struct {
	AliquotVolume    wunit.Volume
	NumberOfAliquots int
	SampleName       *wtype.LHComponent
}

type SimpleAliquotMixOutput struct {
	Aliquots []*wtype.LHComponent
}

type SimpleAliquotMixSOutput struct {
	Data struct {
	}
	Outputs struct {
		Aliquots []*wtype.LHComponent
	}
}

func init() {
	if err := addComponent(component.Component{Name: "SimpleAliquotMix",
		Constructor: SimpleAliquotMixNew,
		Desc: component.ComponentDesc{
			Desc: "Aliquot a solution into a specified plate.\noptionally premix the solution before aliquoting\n",
			Path: "src/github.com/antha-lang/elements/an/AnthaAcademy/AnthaLangAcademy/Lesson3_Mix_Loops/SimpleAliquot/SimpleAliquotMix/SimpleAliquotMix.an",
			Params: []component.ParamDesc{
				{Name: "AliquotVolume", Desc: "", Kind: "Parameters"},
				{Name: "NumberOfAliquots", Desc: "", Kind: "Parameters"},
				{Name: "SampleName", Desc: "", Kind: "Inputs"},
				{Name: "Aliquots", Desc: "", Kind: "Outputs"},
			},
		},
	}); err != nil {
		panic(err)
	}
}