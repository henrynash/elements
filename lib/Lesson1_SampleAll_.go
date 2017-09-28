// example protocol Lessondemonstrating the use of the SampleAll function
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

// the bool type is a "boolean": which essentially means true or false

// Data which is returned from this protocol, and data types

// Physical Inputs to this protocol Lessonwith types

// Physical outputs from this protocol Lessonwith types

func _Lesson1_SampleAllRequirements() {

}

// Conditions to run on startup
func _Lesson1_SampleAllSetup(_ctx context.Context, _input *Lesson1_SampleAllInput) {

}

// The core process for this protocol, with the steps to be performed
// for every input
func _Lesson1_SampleAllSteps(_ctx context.Context, _input *Lesson1_SampleAllInput, _output *Lesson1_SampleAllOutput) {

	_output.Status = "Not sampled anything"

	// the SampleAll function samples the entire contents of the LHComponent
	// so there's no need to specify the volume
	// this if statement specifies that the SampleAll action will only be performed if SampleAll is set to true
	if _input.Sampleall == true {
		_output.Sample = mixer.SampleAll(_input.Solution)
		_output.Status = "Sampled everything"
	}

	// now move on to C_SampleForTotalVolume.an

}

// Run after controls and a steps block are completed to
// post process any data and provide downstream results
func _Lesson1_SampleAllAnalysis(_ctx context.Context, _input *Lesson1_SampleAllInput, _output *Lesson1_SampleAllOutput) {
}

// A block of tests to perform to validate that the sample was processed
//correctly. Optionally, destructive tests can be performed to validate
//results on a dipstick basis
func _Lesson1_SampleAllValidation(_ctx context.Context, _input *Lesson1_SampleAllInput, _output *Lesson1_SampleAllOutput) {

}
func _Lesson1_SampleAllRun(_ctx context.Context, input *Lesson1_SampleAllInput) *Lesson1_SampleAllOutput {
	output := &Lesson1_SampleAllOutput{}
	_Lesson1_SampleAllSetup(_ctx, input)
	_Lesson1_SampleAllSteps(_ctx, input, output)
	_Lesson1_SampleAllAnalysis(_ctx, input, output)
	_Lesson1_SampleAllValidation(_ctx, input, output)
	return output
}

func Lesson1_SampleAllRunSteps(_ctx context.Context, input *Lesson1_SampleAllInput) *Lesson1_SampleAllSOutput {
	soutput := &Lesson1_SampleAllSOutput{}
	output := _Lesson1_SampleAllRun(_ctx, input)
	if err := inject.AssignSome(output, &soutput.Data); err != nil {
		panic(err)
	}
	if err := inject.AssignSome(output, &soutput.Outputs); err != nil {
		panic(err)
	}
	return soutput
}

func Lesson1_SampleAllNew() interface{} {
	return &Lesson1_SampleAllElement{
		inject.CheckedRunner{
			RunFunc: func(_ctx context.Context, value inject.Value) (inject.Value, error) {
				input := &Lesson1_SampleAllInput{}
				if err := inject.Assign(value, input); err != nil {
					return nil, err
				}
				output := _Lesson1_SampleAllRun(_ctx, input)
				return inject.MakeValue(output), nil
			},
			In:  &Lesson1_SampleAllInput{},
			Out: &Lesson1_SampleAllOutput{},
		},
	}
}

var (
	_ = execute.MixInto
	_ = wtype.FALSE
	_ = wunit.Make_units
)

type Lesson1_SampleAllElement struct {
	inject.CheckedRunner
}

type Lesson1_SampleAllInput struct {
	Sampleall bool
	Solution  *wtype.LHComponent
}

type Lesson1_SampleAllOutput struct {
	Sample *wtype.LHComponent
	Status string
}

type Lesson1_SampleAllSOutput struct {
	Data struct {
		Status string
	}
	Outputs struct {
		Sample *wtype.LHComponent
	}
}

func init() {
	if err := addComponent(component.Component{Name: "Lesson1_SampleAll",
		Constructor: Lesson1_SampleAllNew,
		Desc: component.ComponentDesc{
			Desc: "example protocol Lessondemonstrating the use of the SampleAll function\n",
			Path: "src/github.com/antha-lang/elements/AnthaAcademy/Lesson1_RunningWorkflows/B_SampleAll.an",
			Params: []component.ParamDesc{
				{Name: "Sampleall", Desc: "the bool type is a \"boolean\": which essentially means true or false\n", Kind: "Parameters"},
				{Name: "Solution", Desc: "", Kind: "Inputs"},
				{Name: "Sample", Desc: "", Kind: "Outputs"},
				{Name: "Status", Desc: "", Kind: "Data"},
			},
		},
	}); err != nil {
		panic(err)
	}
}
