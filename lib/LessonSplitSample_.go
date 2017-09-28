// protocol LessonSplitSample performs something.
package lib

import

// Place golang packages to import here
(
	"context"
	"github.com/antha-lang/antha/antha/anthalib/mixer"
	"github.com/antha-lang/antha/antha/anthalib/wtype"
	"github.com/antha-lang/antha/antha/anthalib/wunit"
	"github.com/antha-lang/antha/component"
	"github.com/antha-lang/antha/execute"
	"github.com/antha-lang/antha/inject"
)

// Parameters to this protocol

// Output data of this protocol

// Physical inputs to this protocol

// Physical outputs to this protocol

// Conditions to run on startup
func _LessonSplitSampleSetup(_ctx context.Context, _input *LessonSplitSampleInput) {

}

// The core process for this protocol. These steps are executed for each input.
func _LessonSplitSampleSteps(_ctx context.Context, _input *LessonSplitSampleInput, _output *LessonSplitSampleOutput) {
	sampleA := mixer.Sample(_input.InputSolution, _input.VolumeA)
	sampleB := mixer.Sample(_input.InputSolution, _input.VolumeB)
	_output.ComponentA = execute.MixNamed(_ctx, _input.Platetype, _input.WellA, _input.PlateNameA, sampleA)
	_output.ComponentB = execute.MixNamed(_ctx, _input.Platetype, _input.WellB, _input.PlateNameB, sampleB)
}

// Run after controls and a steps block are completed to post process any data
// and provide downstream results
func _LessonSplitSampleAnalysis(_ctx context.Context, _input *LessonSplitSampleInput, _output *LessonSplitSampleOutput) {

}

// A block of tests to perform to validate that the sample was processed
// correctly. Optionally, destructive tests can be performed to validate
// results on a dipstick basis
func _LessonSplitSampleValidation(_ctx context.Context, _input *LessonSplitSampleInput, _output *LessonSplitSampleOutput) {

}
func _LessonSplitSampleRun(_ctx context.Context, input *LessonSplitSampleInput) *LessonSplitSampleOutput {
	output := &LessonSplitSampleOutput{}
	_LessonSplitSampleSetup(_ctx, input)
	_LessonSplitSampleSteps(_ctx, input, output)
	_LessonSplitSampleAnalysis(_ctx, input, output)
	_LessonSplitSampleValidation(_ctx, input, output)
	return output
}

func LessonSplitSampleRunSteps(_ctx context.Context, input *LessonSplitSampleInput) *LessonSplitSampleSOutput {
	soutput := &LessonSplitSampleSOutput{}
	output := _LessonSplitSampleRun(_ctx, input)
	if err := inject.AssignSome(output, &soutput.Data); err != nil {
		panic(err)
	}
	if err := inject.AssignSome(output, &soutput.Outputs); err != nil {
		panic(err)
	}
	return soutput
}

func LessonSplitSampleNew() interface{} {
	return &LessonSplitSampleElement{
		inject.CheckedRunner{
			RunFunc: func(_ctx context.Context, value inject.Value) (inject.Value, error) {
				input := &LessonSplitSampleInput{}
				if err := inject.Assign(value, input); err != nil {
					return nil, err
				}
				output := _LessonSplitSampleRun(_ctx, input)
				return inject.MakeValue(output), nil
			},
			In:  &LessonSplitSampleInput{},
			Out: &LessonSplitSampleOutput{},
		},
	}
}

var (
	_ = execute.MixInto
	_ = wtype.FALSE
	_ = wunit.Make_units
)

type LessonSplitSampleElement struct {
	inject.CheckedRunner
}

type LessonSplitSampleInput struct {
	InputSolution *wtype.LHComponent
	PlateNameA    string
	PlateNameB    string
	Platetype     string
	VolumeA       wunit.Volume
	VolumeB       wunit.Volume
	WellA         string
	WellB         string
}

type LessonSplitSampleOutput struct {
	ComponentA *wtype.LHComponent
	ComponentB *wtype.LHComponent
}

type LessonSplitSampleSOutput struct {
	Data struct {
	}
	Outputs struct {
		ComponentA *wtype.LHComponent
		ComponentB *wtype.LHComponent
	}
}

func init() {
	if err := addComponent(component.Component{Name: "LessonSplitSample",
		Constructor: LessonSplitSampleNew,
		Desc: component.ComponentDesc{
			Desc: "protocol LessonSplitSample performs something.\n",
			Path: "src/github.com/antha-lang/elements/AnthaAcademy/Exercises/LiquidHandlingExercises/SplitSample/SplitSample.an",
			Params: []component.ParamDesc{
				{Name: "InputSolution", Desc: "", Kind: "Inputs"},
				{Name: "PlateNameA", Desc: "", Kind: "Parameters"},
				{Name: "PlateNameB", Desc: "", Kind: "Parameters"},
				{Name: "Platetype", Desc: "", Kind: "Parameters"},
				{Name: "VolumeA", Desc: "", Kind: "Parameters"},
				{Name: "VolumeB", Desc: "", Kind: "Parameters"},
				{Name: "WellA", Desc: "", Kind: "Parameters"},
				{Name: "WellB", Desc: "", Kind: "Parameters"},
				{Name: "ComponentA", Desc: "", Kind: "Outputs"},
				{Name: "ComponentB", Desc: "", Kind: "Outputs"},
			},
		},
	}); err != nil {
		panic(err)
	}
}
