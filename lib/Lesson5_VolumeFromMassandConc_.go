// example of how to convert a concentration and mass to a volume
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

func _Lesson5_VolumeFromMassandConcRequirements() {

}

// Actions to perform before protocol Lessonitself
func _Lesson5_VolumeFromMassandConcSetup(_ctx context.Context, _input *Lesson5_VolumeFromMassandConcInput) {

}

// Core process of the protocol: steps to be performed for each input
func _Lesson5_VolumeFromMassandConcSteps(_ctx context.Context, _input *Lesson5_VolumeFromMassandConcInput, _output *Lesson5_VolumeFromMassandConcOutput) {
	_output.DNAVol, _output.Err = wunit.VolumeForTargetMass(_input.DNAMassperReaction, _input.DNAConc)
	if _output.Err != nil {
		execute.Errorf(_ctx, _output.Err.Error())
	}
}

// Actions to perform after steps block to analyze data
func _Lesson5_VolumeFromMassandConcAnalysis(_ctx context.Context, _input *Lesson5_VolumeFromMassandConcInput, _output *Lesson5_VolumeFromMassandConcOutput) {

}

func _Lesson5_VolumeFromMassandConcValidation(_ctx context.Context, _input *Lesson5_VolumeFromMassandConcInput, _output *Lesson5_VolumeFromMassandConcOutput) {

}
func _Lesson5_VolumeFromMassandConcRun(_ctx context.Context, input *Lesson5_VolumeFromMassandConcInput) *Lesson5_VolumeFromMassandConcOutput {
	output := &Lesson5_VolumeFromMassandConcOutput{}
	_Lesson5_VolumeFromMassandConcSetup(_ctx, input)
	_Lesson5_VolumeFromMassandConcSteps(_ctx, input, output)
	_Lesson5_VolumeFromMassandConcAnalysis(_ctx, input, output)
	_Lesson5_VolumeFromMassandConcValidation(_ctx, input, output)
	return output
}

func Lesson5_VolumeFromMassandConcRunSteps(_ctx context.Context, input *Lesson5_VolumeFromMassandConcInput) *Lesson5_VolumeFromMassandConcSOutput {
	soutput := &Lesson5_VolumeFromMassandConcSOutput{}
	output := _Lesson5_VolumeFromMassandConcRun(_ctx, input)
	if err := inject.AssignSome(output, &soutput.Data); err != nil {
		panic(err)
	}
	if err := inject.AssignSome(output, &soutput.Outputs); err != nil {
		panic(err)
	}
	return soutput
}

func Lesson5_VolumeFromMassandConcNew() interface{} {
	return &Lesson5_VolumeFromMassandConcElement{
		inject.CheckedRunner{
			RunFunc: func(_ctx context.Context, value inject.Value) (inject.Value, error) {
				input := &Lesson5_VolumeFromMassandConcInput{}
				if err := inject.Assign(value, input); err != nil {
					return nil, err
				}
				output := _Lesson5_VolumeFromMassandConcRun(_ctx, input)
				return inject.MakeValue(output), nil
			},
			In:  &Lesson5_VolumeFromMassandConcInput{},
			Out: &Lesson5_VolumeFromMassandConcOutput{},
		},
	}
}

var (
	_ = execute.MixInto
	_ = wtype.FALSE
	_ = wunit.Make_units
)

type Lesson5_VolumeFromMassandConcElement struct {
	inject.CheckedRunner
}

type Lesson5_VolumeFromMassandConcInput struct {
	DNAConc            wunit.Concentration
	DNAMassperReaction wunit.Mass
}

type Lesson5_VolumeFromMassandConcOutput struct {
	DNAVol wunit.Volume
	Err    error
}

type Lesson5_VolumeFromMassandConcSOutput struct {
	Data struct {
		DNAVol wunit.Volume
		Err    error
	}
	Outputs struct {
	}
}

func init() {
	if err := addComponent(component.Component{Name: "Lesson5_VolumeFromMassandConc",
		Constructor: Lesson5_VolumeFromMassandConcNew,
		Desc: component.ComponentDesc{
			Desc: "example of how to convert a concentration and mass to a volume\n",
			Path: "src/github.com/antha-lang/elements/AnthaAcademy/Lesson5_Units2/C_VolumefromMassandConc.an",
			Params: []component.ParamDesc{
				{Name: "DNAConc", Desc: "", Kind: "Parameters"},
				{Name: "DNAMassperReaction", Desc: "", Kind: "Parameters"},
				{Name: "DNAVol", Desc: "", Kind: "Data"},
				{Name: "Err", Desc: "", Kind: "Data"},
			},
		},
	}); err != nil {
		panic(err)
	}
}
