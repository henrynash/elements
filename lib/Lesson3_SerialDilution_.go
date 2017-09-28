// protocol Lessonto Demonstrate how to perform sequential mixing using the example of
// making a serial dilution series from a solution and diluent
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

// e.g. 10 would take 1 part solution to 9 parts diluent for each dilution

// Data which is returned from this protocol, and data types

// Physical Inputs to this protocol Lessonwith types

// Physical outputs from this protocol Lessonwith types

func _Lesson3_SerialDilutionRequirements() {

}

// Conditions to run on startup
func _Lesson3_SerialDilutionSetup(_ctx context.Context, _input *Lesson3_SerialDilutionInput) {

}

// The core process for this protocol, with the steps to be performed
// for every input
func _Lesson3_SerialDilutionSteps(_ctx context.Context, _input *Lesson3_SerialDilutionInput, _output *Lesson3_SerialDilutionOutput) {

	dilutions := make([]*wtype.LHComponent, 0)

	var aliquot *wtype.LHComponent

	// calculate solution volume

	// create copy of TotalVolumeperDilution
	solutionVolume := (wunit.CopyVolume(_input.TotalVolumeperDilution))

	// use divideby method
	solutionVolume.DivideBy(float64(_input.DilutionFactor))

	// use same approach to work out diluent volume to add
	diluentVolume := (wunit.CopyVolume(_input.TotalVolumeperDilution))

	// this time using the substract method
	diluentVolume.Subtract(solutionVolume)

	// sample diluent
	diluentSample := mixer.Sample(_input.Diluent, diluentVolume)

	// Ensure liquid type set to Pre and Post Mix
	_input.Solution.Type = wtype.LTNeedToMix
	// check if the enzyme is specified and if not mix the

	// sample solution
	solutionSample := mixer.Sample(_input.Solution, solutionVolume)

	// mix both samples to OutPlate
	aliquot = execute.MixTo(_ctx, _input.OutPlate.Type, "", 1, diluentSample, solutionSample)

	// add to dilutions array
	dilutions = append(dilutions, aliquot)

	// loop through NumberOfDilutions until all serial dilutions are made
	for k := 1; k < _input.NumberOfDilutions; k++ {

		// take next sample of diluent
		nextdiluentSample := mixer.Sample(_input.Diluent, diluentVolume)

		// Ensure liquid type set to Pre and Post Mix
		aliquot.Type = wtype.LTNeedToMix

		// sample from previous dilution sample
		nextSample := mixer.Sample(aliquot, solutionVolume)

		// Mix sample into nextdiluent sample
		nextaliquot := execute.Mix(_ctx, nextdiluentSample, nextSample)

		// add to dilutions array
		dilutions = append(dilutions, nextaliquot)
		// reset aliquot
		aliquot = nextaliquot
	}

	// export as Output
	_output.Dilutions = dilutions

}

// Run after controls and a steps block are completed to
// post process any data and provide downstream results
func _Lesson3_SerialDilutionAnalysis(_ctx context.Context, _input *Lesson3_SerialDilutionInput, _output *Lesson3_SerialDilutionOutput) {
}

// A block of tests to perform to validate that the sample was processed correctly
// Optionally, destructive tests can be performed to validate results on a
// dipstick basis
func _Lesson3_SerialDilutionValidation(_ctx context.Context, _input *Lesson3_SerialDilutionInput, _output *Lesson3_SerialDilutionOutput) {

}
func _Lesson3_SerialDilutionRun(_ctx context.Context, input *Lesson3_SerialDilutionInput) *Lesson3_SerialDilutionOutput {
	output := &Lesson3_SerialDilutionOutput{}
	_Lesson3_SerialDilutionSetup(_ctx, input)
	_Lesson3_SerialDilutionSteps(_ctx, input, output)
	_Lesson3_SerialDilutionAnalysis(_ctx, input, output)
	_Lesson3_SerialDilutionValidation(_ctx, input, output)
	return output
}

func Lesson3_SerialDilutionRunSteps(_ctx context.Context, input *Lesson3_SerialDilutionInput) *Lesson3_SerialDilutionSOutput {
	soutput := &Lesson3_SerialDilutionSOutput{}
	output := _Lesson3_SerialDilutionRun(_ctx, input)
	if err := inject.AssignSome(output, &soutput.Data); err != nil {
		panic(err)
	}
	if err := inject.AssignSome(output, &soutput.Outputs); err != nil {
		panic(err)
	}
	return soutput
}

func Lesson3_SerialDilutionNew() interface{} {
	return &Lesson3_SerialDilutionElement{
		inject.CheckedRunner{
			RunFunc: func(_ctx context.Context, value inject.Value) (inject.Value, error) {
				input := &Lesson3_SerialDilutionInput{}
				if err := inject.Assign(value, input); err != nil {
					return nil, err
				}
				output := _Lesson3_SerialDilutionRun(_ctx, input)
				return inject.MakeValue(output), nil
			},
			In:  &Lesson3_SerialDilutionInput{},
			Out: &Lesson3_SerialDilutionOutput{},
		},
	}
}

var (
	_ = execute.MixInto
	_ = wtype.FALSE
	_ = wunit.Make_units
)

type Lesson3_SerialDilutionElement struct {
	inject.CheckedRunner
}

type Lesson3_SerialDilutionInput struct {
	Diluent                *wtype.LHComponent
	DilutionFactor         int
	NumberOfDilutions      int
	OutPlate               *wtype.LHPlate
	Solution               *wtype.LHComponent
	TotalVolumeperDilution wunit.Volume
}

type Lesson3_SerialDilutionOutput struct {
	Dilutions []*wtype.LHComponent
}

type Lesson3_SerialDilutionSOutput struct {
	Data struct {
	}
	Outputs struct {
		Dilutions []*wtype.LHComponent
	}
}

func init() {
	if err := addComponent(component.Component{Name: "Lesson3_SerialDilution",
		Constructor: Lesson3_SerialDilutionNew,
		Desc: component.ComponentDesc{
			Desc: "protocol Lessonto Demonstrate how to perform sequential mixing using the example of\nmaking a serial dilution series from a solution and diluent\n",
			Path: "src/github.com/antha-lang/elements/AnthaAcademy/Lesson3_MixPart2/C_SequentialMixing.an",
			Params: []component.ParamDesc{
				{Name: "Diluent", Desc: "", Kind: "Inputs"},
				{Name: "DilutionFactor", Desc: "e.g. 10 would take 1 part solution to 9 parts diluent for each dilution\n", Kind: "Parameters"},
				{Name: "NumberOfDilutions", Desc: "", Kind: "Parameters"},
				{Name: "OutPlate", Desc: "", Kind: "Inputs"},
				{Name: "Solution", Desc: "", Kind: "Inputs"},
				{Name: "TotalVolumeperDilution", Desc: "", Kind: "Parameters"},
				{Name: "Dilutions", Desc: "", Kind: "Outputs"},
			},
		},
	}); err != nil {
		panic(err)
	}
}
