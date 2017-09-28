// Make a general mastermix comprising of a list of components, list of volumes
// and specifying the number of reactions required
package lib

import (
	"context"
	"fmt"
	"github.com/antha-lang/antha/antha/anthalib/mixer"
	"github.com/antha-lang/antha/antha/anthalib/wtype"
	"github.com/antha-lang/antha/antha/anthalib/wunit"
	"github.com/antha-lang/antha/component"
	"github.com/antha-lang/antha/execute"
	"github.com/antha-lang/antha/inject"
)

// Input parameters for this protocol Lesson(data)

// Data which is returned from this protocol, and data types

// Physical Inputs to this protocol Lessonwith types

//TopUpBuffer *LHComponent // optional if nil this is ignored

// Physical outputs from this protocol Lessonwith types

func _Lesson0_MasterMixRequirements() {
}

// Conditions to run on startup
func _Lesson0_MasterMixSetup(_ctx context.Context, _input *Lesson0_MasterMixInput) {
}

// The core process for this protocol, with the steps to be performed
// for every input
func _Lesson0_MasterMixSteps(_ctx context.Context, _input *Lesson0_MasterMixInput, _output *Lesson0_MasterMixOutput) {

	var mastermix *wtype.LHComponent

	if len(_input.Components) != len(_input.ComponentVolumesperReaction) {
		panic("len(Components) != len(OtherComponentVolumes)")
	}

	eachmastermix := make([]*wtype.LHComponent, 0)

	for k, component := range _input.Components {
		if k == len(_input.Components) {
			component.Type = wtype.LTNeedToMix //"NeedToMix"
		}

		// multiply volume of each component by number of reactions per mastermix
		adjustedvol := wunit.NewVolume(float64(_input.Reactionspermastermix)*_input.ComponentVolumesperReaction[k].SIValue()*1000000, "ul")

		componentSample := mixer.Sample(component, adjustedvol)
		component.CName = "component" + fmt.Sprint(k+1)
		eachmastermix = append(eachmastermix, componentSample)

	}
	mastermix = execute.MixInto(_ctx, _input.OutPlate, "", eachmastermix...)

	_output.Mastermix = mastermix

	_output.Status = "Mastermix Made"

}

// Run after controls and a steps block are completed to
// post process any data and provide downstream results
func _Lesson0_MasterMixAnalysis(_ctx context.Context, _input *Lesson0_MasterMixInput, _output *Lesson0_MasterMixOutput) {
}

// A block of tests to perform to validate that the sample was processed correctly
// Optionally, destructive tests can be performed to validate results on a
// dipstick basis
func _Lesson0_MasterMixValidation(_ctx context.Context, _input *Lesson0_MasterMixInput, _output *Lesson0_MasterMixOutput) {
}
func _Lesson0_MasterMixRun(_ctx context.Context, input *Lesson0_MasterMixInput) *Lesson0_MasterMixOutput {
	output := &Lesson0_MasterMixOutput{}
	_Lesson0_MasterMixSetup(_ctx, input)
	_Lesson0_MasterMixSteps(_ctx, input, output)
	_Lesson0_MasterMixAnalysis(_ctx, input, output)
	_Lesson0_MasterMixValidation(_ctx, input, output)
	return output
}

func Lesson0_MasterMixRunSteps(_ctx context.Context, input *Lesson0_MasterMixInput) *Lesson0_MasterMixSOutput {
	soutput := &Lesson0_MasterMixSOutput{}
	output := _Lesson0_MasterMixRun(_ctx, input)
	if err := inject.AssignSome(output, &soutput.Data); err != nil {
		panic(err)
	}
	if err := inject.AssignSome(output, &soutput.Outputs); err != nil {
		panic(err)
	}
	return soutput
}

func Lesson0_MasterMixNew() interface{} {
	return &Lesson0_MasterMixElement{
		inject.CheckedRunner{
			RunFunc: func(_ctx context.Context, value inject.Value) (inject.Value, error) {
				input := &Lesson0_MasterMixInput{}
				if err := inject.Assign(value, input); err != nil {
					return nil, err
				}
				output := _Lesson0_MasterMixRun(_ctx, input)
				return inject.MakeValue(output), nil
			},
			In:  &Lesson0_MasterMixInput{},
			Out: &Lesson0_MasterMixOutput{},
		},
	}
}

var (
	_ = execute.MixInto
	_ = wtype.FALSE
	_ = wunit.Make_units
)

type Lesson0_MasterMixElement struct {
	inject.CheckedRunner
}

type Lesson0_MasterMixInput struct {
	ComponentVolumesperReaction []wunit.Volume
	Components                  []*wtype.LHComponent
	OutPlate                    *wtype.LHPlate
	Reactionspermastermix       int
}

type Lesson0_MasterMixOutput struct {
	Mastermix *wtype.LHComponent
	Status    string
}

type Lesson0_MasterMixSOutput struct {
	Data struct {
		Status string
	}
	Outputs struct {
		Mastermix *wtype.LHComponent
	}
}

func init() {
	if err := addComponent(component.Component{Name: "Lesson0_MasterMix",
		Constructor: Lesson0_MasterMixNew,
		Desc: component.ComponentDesc{
			Desc: "Make a general mastermix comprising of a list of components, list of volumes\nand specifying the number of reactions required\n",
			Path: "src/github.com/antha-lang/elements/AnthaAcademy/Lesson0_Examples/MakeMasterMix_PCR/Mastermix_one.an",
			Params: []component.ParamDesc{
				{Name: "ComponentVolumesperReaction", Desc: "", Kind: "Parameters"},
				{Name: "Components", Desc: "TopUpBuffer *LHComponent // optional if nil this is ignored\n", Kind: "Inputs"},
				{Name: "OutPlate", Desc: "", Kind: "Inputs"},
				{Name: "Reactionspermastermix", Desc: "", Kind: "Parameters"},
				{Name: "Mastermix", Desc: "", Kind: "Outputs"},
				{Name: "Status", Desc: "", Kind: "Data"},
			},
		},
	}); err != nil {
		panic(err)
	}
}
