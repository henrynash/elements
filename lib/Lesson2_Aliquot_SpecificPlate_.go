// example protocol Lessonshowing The MixInto command which allows a specifc plate to be specified. i.e. plate with ID blahblahblah
// if no specific plate is specified a plate type can be
package lib

import (
	"context"
	"github.com/antha-lang/antha/antha/anthalib/mixer"
	"github.com/antha-lang/antha/antha/anthalib/wtype"
	"github.com/antha-lang/antha/antha/anthalib/wunit"
	"github.com/antha-lang/antha/antha/anthalib/wutil"
	"github.com/antha-lang/antha/component"
	"github.com/antha-lang/antha/execute"
	"github.com/antha-lang/antha/inject"
)

// Input parameters for this protocol Lesson(data)

// Data which is returned from this protocol, and data types

// Physical Inputs to this protocol Lessonwith types

// this time we're specifying what plate we're using

// Physical outputs from this protocol Lessonwith types

func _Lesson2_Aliquot_SpecificPlateRequirements() {

}

// Conditions to run on startup
func _Lesson2_Aliquot_SpecificPlateSetup(_ctx context.Context, _input *Lesson2_Aliquot_SpecificPlateInput) {

}

// The core process for this protocol, with the steps to be performed
// for every input
func _Lesson2_Aliquot_SpecificPlateSteps(_ctx context.Context, _input *Lesson2_Aliquot_SpecificPlateInput, _output *Lesson2_Aliquot_SpecificPlateOutput) {

	number := _input.SolutionVolume.SIValue() / _input.VolumePerAliquot.SIValue()
	possiblenumberofAliquots, _ := wutil.RoundDown(number)
	if possiblenumberofAliquots < _input.NumberofAliquots {
		execute.Errorf(_ctx, "Not enough solution for this many aliquots")
	}

	aliquots := make([]*wtype.LHComponent, 0)

	for i := 0; i < _input.NumberofAliquots; i++ {
		if _input.Solution.TypeName() == "dna" {
			_input.Solution.Type = wtype.LTDoNotMix
		}
		aliquotSample := mixer.Sample(_input.Solution, _input.VolumePerAliquot)

		// the MixInto command is used instead of Mix to specify the plate
		// MixInto allows you to specify the exact plate to MixInto (i.e. rather than just a plate type. e.g. barcode 123214234)
		// the three input fields to the MixInto command represent
		// 1. the plate
		// 2. well location as a  string e.g. "A1" (in this case leaving it blank "" will leave the well location up to the scheduler),
		// 3. the sample or array of samples to be mixed
		aliquot := execute.MixInto(_ctx, _input.OutPlate, "", aliquotSample)
		aliquots = append(aliquots, aliquot)
	}
	_output.Aliquots = aliquots
}

// Run after controls and a steps block are completed to
// post process any data and provide downstream results
func _Lesson2_Aliquot_SpecificPlateAnalysis(_ctx context.Context, _input *Lesson2_Aliquot_SpecificPlateInput, _output *Lesson2_Aliquot_SpecificPlateOutput) {
}

// A block of tests to perform to validate that the sample was processed
//correctly. Optionally, destructive tests can be performed to validate
//results on a dipstick basis
func _Lesson2_Aliquot_SpecificPlateValidation(_ctx context.Context, _input *Lesson2_Aliquot_SpecificPlateInput, _output *Lesson2_Aliquot_SpecificPlateOutput) {

}
func _Lesson2_Aliquot_SpecificPlateRun(_ctx context.Context, input *Lesson2_Aliquot_SpecificPlateInput) *Lesson2_Aliquot_SpecificPlateOutput {
	output := &Lesson2_Aliquot_SpecificPlateOutput{}
	_Lesson2_Aliquot_SpecificPlateSetup(_ctx, input)
	_Lesson2_Aliquot_SpecificPlateSteps(_ctx, input, output)
	_Lesson2_Aliquot_SpecificPlateAnalysis(_ctx, input, output)
	_Lesson2_Aliquot_SpecificPlateValidation(_ctx, input, output)
	return output
}

func Lesson2_Aliquot_SpecificPlateRunSteps(_ctx context.Context, input *Lesson2_Aliquot_SpecificPlateInput) *Lesson2_Aliquot_SpecificPlateSOutput {
	soutput := &Lesson2_Aliquot_SpecificPlateSOutput{}
	output := _Lesson2_Aliquot_SpecificPlateRun(_ctx, input)
	if err := inject.AssignSome(output, &soutput.Data); err != nil {
		panic(err)
	}
	if err := inject.AssignSome(output, &soutput.Outputs); err != nil {
		panic(err)
	}
	return soutput
}

func Lesson2_Aliquot_SpecificPlateNew() interface{} {
	return &Lesson2_Aliquot_SpecificPlateElement{
		inject.CheckedRunner{
			RunFunc: func(_ctx context.Context, value inject.Value) (inject.Value, error) {
				input := &Lesson2_Aliquot_SpecificPlateInput{}
				if err := inject.Assign(value, input); err != nil {
					return nil, err
				}
				output := _Lesson2_Aliquot_SpecificPlateRun(_ctx, input)
				return inject.MakeValue(output), nil
			},
			In:  &Lesson2_Aliquot_SpecificPlateInput{},
			Out: &Lesson2_Aliquot_SpecificPlateOutput{},
		},
	}
}

var (
	_ = execute.MixInto
	_ = wtype.FALSE
	_ = wunit.Make_units
)

type Lesson2_Aliquot_SpecificPlateElement struct {
	inject.CheckedRunner
}

type Lesson2_Aliquot_SpecificPlateInput struct {
	NumberofAliquots int
	OutPlate         *wtype.LHPlate
	Solution         *wtype.LHComponent
	SolutionVolume   wunit.Volume
	VolumePerAliquot wunit.Volume
}

type Lesson2_Aliquot_SpecificPlateOutput struct {
	Aliquots []*wtype.LHComponent
}

type Lesson2_Aliquot_SpecificPlateSOutput struct {
	Data struct {
	}
	Outputs struct {
		Aliquots []*wtype.LHComponent
	}
}

func init() {
	if err := addComponent(component.Component{Name: "Lesson2_Aliquot_SpecificPlate",
		Constructor: Lesson2_Aliquot_SpecificPlateNew,
		Desc: component.ComponentDesc{
			Desc: "example protocol Lessonshowing The MixInto command which allows a specifc plate to be specified. i.e. plate with ID blahblahblah\nif no specific plate is specified a plate type can be\n",
			Path: "src/github.com/antha-lang/elements/AnthaAcademy/Lesson2_mix/B_AliquotIntoSpecificPlate.an",
			Params: []component.ParamDesc{
				{Name: "NumberofAliquots", Desc: "", Kind: "Parameters"},
				{Name: "OutPlate", Desc: "this time we're specifying what plate we're using\n", Kind: "Inputs"},
				{Name: "Solution", Desc: "", Kind: "Inputs"},
				{Name: "SolutionVolume", Desc: "", Kind: "Parameters"},
				{Name: "VolumePerAliquot", Desc: "", Kind: "Parameters"},
				{Name: "Aliquots", Desc: "", Kind: "Outputs"},
			},
		},
	}); err != nil {
		panic(err)
	}
}
