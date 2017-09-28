// Example protocol Lessondemonstrating the use of the Sample function
package lib

import // this is the name of the protocol Lessonthat will be called in a workflow or other antha element

// we need to import the wtype package to use the LHComponent type
// the mixer package is required to use the Sample function
(
	"context"
	"github.com/antha-lang/antha/antha/anthalib/mixer"
	"github.com/antha-lang/antha/antha/anthalib/wtype"
	"github.com/antha-lang/antha/antha/anthalib/wunit"
	"github.com/antha-lang/antha/component"
	"github.com/antha-lang/antha/execute"
	"github.com/antha-lang/antha/inject"
)

// Input parameters for this protocol Lesson(data)

// antha, like golang is a strongly typed language in which the type of a variable must be declared.
// In this case we're creating a variable called SampleVolume which is of type Volume;
// the type system allows the antha compiler to catch many types of common errors before the programme is run
// the antha type system extends this to biological types such as volumes here.
// functions require inputs of particular types to be adhered to

// Data which is returned from this protocol, and data types

// Antha inherits all standard primitives valid in golang;
//for example the string type shown here used to return a textual message

// Physical Inputs to this protocol Lessonwith types

// the LHComponent is the principal liquidhandling type in antha
// the * signifies that this is a pointer to the component rather than the component itself
// most key antha functions such as Sample and Mix use *LHComponent rather than LHComponent
// since the type is imported from the wtype package we need to use  *LHComponent rather than simply *LHComponent

// Physical outputs from this protocol Lessonwith types

// An output LHComponent variable is created called Sample

func _Lesson1_SampleRequirements() {

}

// Conditions to run on startup
func _Lesson1_SampleSetup(_ctx context.Context, _input *Lesson1_SampleInput) {

}

// The core process for this protocol, with the steps to be performed
// for every input
func _Lesson1_SampleSteps(_ctx context.Context, _input *Lesson1_SampleInput, _output *Lesson1_SampleOutput) {

	// the Sample function is imported from the mixer library
	// in the mixer library the function signature can be found, here it is:
	// func Sample(l *LHComponent, v Volume) *LHComponent {
	// The function signature  shows that the function requires a *LHComponent and a Volume and returns an *LHComponent
	_output.Sample = mixer.Sample(_input.Solution, _input.SampleVolume)

	// The Sample function is not sufficient to generate liquid handling instructions alone,
	// We would need a Mix command to instruct where to put the sample

	// we can also create data outputs as a string like this
	_output.Status = _input.SampleVolume.ToString() + " of " + _input.Solution.CName + " sampled"

}

// Run after controls and a steps block are completed to
// post process any data and provide downstream results
func _Lesson1_SampleAnalysis(_ctx context.Context, _input *Lesson1_SampleInput, _output *Lesson1_SampleOutput) {
}

// A block of tests to perform to validate that the sample was processed
//correctly. Optionally, destructive tests can be performed to validate
//results on a dipstick basis
func _Lesson1_SampleValidation(_ctx context.Context, _input *Lesson1_SampleInput, _output *Lesson1_SampleOutput) {

}
func _Lesson1_SampleRun(_ctx context.Context, input *Lesson1_SampleInput) *Lesson1_SampleOutput {
	output := &Lesson1_SampleOutput{}
	_Lesson1_SampleSetup(_ctx, input)
	_Lesson1_SampleSteps(_ctx, input, output)
	_Lesson1_SampleAnalysis(_ctx, input, output)
	_Lesson1_SampleValidation(_ctx, input, output)
	return output
}

func Lesson1_SampleRunSteps(_ctx context.Context, input *Lesson1_SampleInput) *Lesson1_SampleSOutput {
	soutput := &Lesson1_SampleSOutput{}
	output := _Lesson1_SampleRun(_ctx, input)
	if err := inject.AssignSome(output, &soutput.Data); err != nil {
		panic(err)
	}
	if err := inject.AssignSome(output, &soutput.Outputs); err != nil {
		panic(err)
	}
	return soutput
}

func Lesson1_SampleNew() interface{} {
	return &Lesson1_SampleElement{
		inject.CheckedRunner{
			RunFunc: func(_ctx context.Context, value inject.Value) (inject.Value, error) {
				input := &Lesson1_SampleInput{}
				if err := inject.Assign(value, input); err != nil {
					return nil, err
				}
				output := _Lesson1_SampleRun(_ctx, input)
				return inject.MakeValue(output), nil
			},
			In:  &Lesson1_SampleInput{},
			Out: &Lesson1_SampleOutput{},
		},
	}
}

var (
	_ = execute.MixInto
	_ = wtype.FALSE
	_ = wunit.Make_units
)

type Lesson1_SampleElement struct {
	inject.CheckedRunner
}

type Lesson1_SampleInput struct {
	SampleVolume wunit.Volume
	Solution     *wtype.LHComponent
}

type Lesson1_SampleOutput struct {
	Sample *wtype.LHComponent
	Status string
}

type Lesson1_SampleSOutput struct {
	Data struct {
		Status string
	}
	Outputs struct {
		Sample *wtype.LHComponent
	}
}

func init() {
	if err := addComponent(component.Component{Name: "Lesson1_Sample",
		Constructor: Lesson1_SampleNew,
		Desc: component.ComponentDesc{
			Desc: "Example protocol Lessondemonstrating the use of the Sample function\n",
			Path: "src/github.com/antha-lang/elements/AnthaAcademy/Lesson1_RunningWorkflows/A_Sample.an",
			Params: []component.ParamDesc{
				{Name: "SampleVolume", Desc: "antha, like golang is a strongly typed language in which the type of a variable must be declared.\nIn this case we're creating a variable called SampleVolume which is of type Volume;\nthe type system allows the antha compiler to catch many types of common errors before the programme is run\nthe antha type system extends this to biological types such as volumes here.\nfunctions require inputs of particular types to be adhered to\n", Kind: "Parameters"},
				{Name: "Solution", Desc: "the LHComponent is the principal liquidhandling type in antha\nthe * signifies that this is a pointer to the component rather than the component itself\nmost key antha functions such as Sample and Mix use *LHComponent rather than LHComponent\nsince the type is imported from the wtype package we need to use  *LHComponent rather than simply *LHComponent\n", Kind: "Inputs"},
				{Name: "Sample", Desc: "An output LHComponent variable is created called Sample\n", Kind: "Outputs"},
				{Name: "Status", Desc: "Antha inherits all standard primitives valid in golang;\nfor example the string type shown here used to return a textual message\n", Kind: "Data"},
			},
		},
	}); err != nil {
		panic(err)
	}
}
