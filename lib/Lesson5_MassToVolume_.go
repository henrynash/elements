// example of how to convert a density and mass to a volume
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

func _Lesson5_MassToVolumeRequirements() {

}

// Actions to perform before protocol Lessonitself
func _Lesson5_MassToVolumeSetup(_ctx context.Context, _input *Lesson5_MassToVolumeInput) {

}

// Core process of the protocol: steps to be performed for each input
func _Lesson5_MassToVolumeSteps(_ctx context.Context, _input *Lesson5_MassToVolumeInput, _output *Lesson5_MassToVolumeOutput) {

	_output.Vol = wunit.MasstoVolume(_input.MyMass, _input.MyDensity)

	_output.BacktoMass = wunit.VolumetoMass(_output.Vol, _input.MyDensity)
}

// Actions to perform after steps block to analyze data
func _Lesson5_MassToVolumeAnalysis(_ctx context.Context, _input *Lesson5_MassToVolumeInput, _output *Lesson5_MassToVolumeOutput) {

}

func _Lesson5_MassToVolumeValidation(_ctx context.Context, _input *Lesson5_MassToVolumeInput, _output *Lesson5_MassToVolumeOutput) {

}
func _Lesson5_MassToVolumeRun(_ctx context.Context, input *Lesson5_MassToVolumeInput) *Lesson5_MassToVolumeOutput {
	output := &Lesson5_MassToVolumeOutput{}
	_Lesson5_MassToVolumeSetup(_ctx, input)
	_Lesson5_MassToVolumeSteps(_ctx, input, output)
	_Lesson5_MassToVolumeAnalysis(_ctx, input, output)
	_Lesson5_MassToVolumeValidation(_ctx, input, output)
	return output
}

func Lesson5_MassToVolumeRunSteps(_ctx context.Context, input *Lesson5_MassToVolumeInput) *Lesson5_MassToVolumeSOutput {
	soutput := &Lesson5_MassToVolumeSOutput{}
	output := _Lesson5_MassToVolumeRun(_ctx, input)
	if err := inject.AssignSome(output, &soutput.Data); err != nil {
		panic(err)
	}
	if err := inject.AssignSome(output, &soutput.Outputs); err != nil {
		panic(err)
	}
	return soutput
}

func Lesson5_MassToVolumeNew() interface{} {
	return &Lesson5_MassToVolumeElement{
		inject.CheckedRunner{
			RunFunc: func(_ctx context.Context, value inject.Value) (inject.Value, error) {
				input := &Lesson5_MassToVolumeInput{}
				if err := inject.Assign(value, input); err != nil {
					return nil, err
				}
				output := _Lesson5_MassToVolumeRun(_ctx, input)
				return inject.MakeValue(output), nil
			},
			In:  &Lesson5_MassToVolumeInput{},
			Out: &Lesson5_MassToVolumeOutput{},
		},
	}
}

var (
	_ = execute.MixInto
	_ = wtype.FALSE
	_ = wunit.Make_units
)

type Lesson5_MassToVolumeElement struct {
	inject.CheckedRunner
}

type Lesson5_MassToVolumeInput struct {
	MyDensity wunit.Density
	MyMass    wunit.Mass
}

type Lesson5_MassToVolumeOutput struct {
	BacktoMass wunit.Mass
	Vol        wunit.Volume
}

type Lesson5_MassToVolumeSOutput struct {
	Data struct {
		BacktoMass wunit.Mass
		Vol        wunit.Volume
	}
	Outputs struct {
	}
}

func init() {
	if err := addComponent(component.Component{Name: "Lesson5_MassToVolume",
		Constructor: Lesson5_MassToVolumeNew,
		Desc: component.ComponentDesc{
			Desc: "example of how to convert a density and mass to a volume\n",
			Path: "src/github.com/antha-lang/elements/AnthaAcademy/Lesson5_Units2/B_MasstoVolume.an",
			Params: []component.ParamDesc{
				{Name: "MyDensity", Desc: "", Kind: "Parameters"},
				{Name: "MyMass", Desc: "", Kind: "Parameters"},
				{Name: "BacktoMass", Desc: "", Kind: "Data"},
				{Name: "Vol", Desc: "", Kind: "Data"},
			},
		},
	}); err != nil {
		panic(err)
	}
}
