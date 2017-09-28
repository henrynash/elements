// example protocol Lessondemonstrating the use of the SampleForTotalVolume function
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

// Input parameters for this protocol Lesson(data)

// e.g. 2ul
// e.g. 20ul

// Data which is returned from this protocol, and data types

// Physical Inputs to this protocol Lessonwith types

// Physical outputs from this protocol Lessonwith types

func _Lesson1_SampleForTotalVolumeRequirements() {

}

// Conditions to run on startup
func _Lesson1_SampleForTotalVolumeSetup(_ctx context.Context, _input *Lesson1_SampleForTotalVolumeInput) {

}

// The core process for this protocol, with the steps to be performed
// for every input
func _Lesson1_SampleForTotalVolumeSteps(_ctx context.Context, _input *Lesson1_SampleForTotalVolumeInput, _output *Lesson1_SampleForTotalVolumeOutput) {

	// make empty slice of LHComponents (i.e. of length 0) ready to sequentially add all samples to
	// See golangbook chapter 6 for more details on slices and arrays
	allsamples := make([]*wtype.LHComponent, 0)

	// SampleForTotalVolume will "top up" solution to the TotalVolume with Diluent.
	// In this case it will still add diluent first but calculates the volume to add by substracting the volumes of subsequent components
	diluentsample := mixer.SampleForTotalVolume(_input.Diluent, _input.TotalVolume) // i.e. if TotalVolume == 20ul and SolutionVolume == 2ul then 18ul of Diluent will be sampled here

	// append will add the diluent sample to the allsamples slice
	allsamples = append(allsamples, diluentsample)

	solutionsample := mixer.Sample(_input.Solution, _input.SolutionVolume)

	allsamples = append(allsamples, solutionsample)

	// The Sample functions will not generate liquid handling instructions on their own
	// We need to tell Antha what to do with samples
	// For this we need to use one of the Mix functions
	// therefore finally we use Mix to combine samples into a new component
	_output.DilutedSample = execute.Mix(_ctx, allsamples...)

	// Now we have an antha element which will generate liquid handling instructions
	// let's see how to actually run the protocol
	// open the terminal and
	// work your way through the lessons there showing how to specify parameters and different types of workflow

}

// Run after controls and a steps block are completed to
// post process any data and provide downstream results
func _Lesson1_SampleForTotalVolumeAnalysis(_ctx context.Context, _input *Lesson1_SampleForTotalVolumeInput, _output *Lesson1_SampleForTotalVolumeOutput) {
}

// A block of tests to perform to validate that the sample was processed
//correctly. Optionally, destructive tests can be performed to validate
//results on a dipstick basis
func _Lesson1_SampleForTotalVolumeValidation(_ctx context.Context, _input *Lesson1_SampleForTotalVolumeInput, _output *Lesson1_SampleForTotalVolumeOutput) {

}
func _Lesson1_SampleForTotalVolumeRun(_ctx context.Context, input *Lesson1_SampleForTotalVolumeInput) *Lesson1_SampleForTotalVolumeOutput {
	output := &Lesson1_SampleForTotalVolumeOutput{}
	_Lesson1_SampleForTotalVolumeSetup(_ctx, input)
	_Lesson1_SampleForTotalVolumeSteps(_ctx, input, output)
	_Lesson1_SampleForTotalVolumeAnalysis(_ctx, input, output)
	_Lesson1_SampleForTotalVolumeValidation(_ctx, input, output)
	return output
}

func Lesson1_SampleForTotalVolumeRunSteps(_ctx context.Context, input *Lesson1_SampleForTotalVolumeInput) *Lesson1_SampleForTotalVolumeSOutput {
	soutput := &Lesson1_SampleForTotalVolumeSOutput{}
	output := _Lesson1_SampleForTotalVolumeRun(_ctx, input)
	if err := inject.AssignSome(output, &soutput.Data); err != nil {
		panic(err)
	}
	if err := inject.AssignSome(output, &soutput.Outputs); err != nil {
		panic(err)
	}
	return soutput
}

func Lesson1_SampleForTotalVolumeNew() interface{} {
	return &Lesson1_SampleForTotalVolumeElement{
		inject.CheckedRunner{
			RunFunc: func(_ctx context.Context, value inject.Value) (inject.Value, error) {
				input := &Lesson1_SampleForTotalVolumeInput{}
				if err := inject.Assign(value, input); err != nil {
					return nil, err
				}
				output := _Lesson1_SampleForTotalVolumeRun(_ctx, input)
				return inject.MakeValue(output), nil
			},
			In:  &Lesson1_SampleForTotalVolumeInput{},
			Out: &Lesson1_SampleForTotalVolumeOutput{},
		},
	}
}

var (
	_ = execute.MixInto
	_ = wtype.FALSE
	_ = wunit.Make_units
)

type Lesson1_SampleForTotalVolumeElement struct {
	inject.CheckedRunner
}

type Lesson1_SampleForTotalVolumeInput struct {
	Diluent        *wtype.LHComponent
	Solution       *wtype.LHComponent
	SolutionVolume wunit.Volume
	TotalVolume    wunit.Volume
}

type Lesson1_SampleForTotalVolumeOutput struct {
	DilutedSample *wtype.LHComponent
}

type Lesson1_SampleForTotalVolumeSOutput struct {
	Data struct {
	}
	Outputs struct {
		DilutedSample *wtype.LHComponent
	}
}

func init() {
	if err := addComponent(component.Component{Name: "Lesson1_SampleForTotalVolume",
		Constructor: Lesson1_SampleForTotalVolumeNew,
		Desc: component.ComponentDesc{
			Desc: "example protocol Lessondemonstrating the use of the SampleForTotalVolume function\n",
			Path: "src/github.com/antha-lang/elements/AnthaAcademy/Lesson1_RunningWorkflows/C_SampleForTotalVolume.an",
			Params: []component.ParamDesc{
				{Name: "Diluent", Desc: "", Kind: "Inputs"},
				{Name: "Solution", Desc: "", Kind: "Inputs"},
				{Name: "SolutionVolume", Desc: "e.g. 2ul\n", Kind: "Parameters"},
				{Name: "TotalVolume", Desc: "e.g. 20ul\n", Kind: "Parameters"},
				{Name: "DilutedSample", Desc: "", Kind: "Outputs"},
			},
		},
	}); err != nil {
		panic(err)
	}
}
