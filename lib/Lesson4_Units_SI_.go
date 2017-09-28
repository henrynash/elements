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

func _Lesson4_Units_SIRequirements() {

}

// Actions to perform before protocol Lessonitself
func _Lesson4_Units_SISetup(_ctx context.Context, _input *Lesson4_Units_SIInput) {

}

// Core process of the protocol: steps to be performed for each input
func _Lesson4_Units_SISteps(_ctx context.Context, _input *Lesson4_Units_SIInput, _output *Lesson4_Units_SIOutput) {

	_output.SIMass = _input.MyMass.SIValue()
	_output.SIMassUnit = _input.MyMass.Unit().BaseSISymbol()

	_output.RawMass = _input.MyMass.RawValue()
	_output.RawMassUnit = _input.MyMass.Unit().PrefixedSymbol()

}

// Actions to perform after steps block to analyze data
func _Lesson4_Units_SIAnalysis(_ctx context.Context, _input *Lesson4_Units_SIInput, _output *Lesson4_Units_SIOutput) {

}

func _Lesson4_Units_SIValidation(_ctx context.Context, _input *Lesson4_Units_SIInput, _output *Lesson4_Units_SIOutput) {

}
func _Lesson4_Units_SIRun(_ctx context.Context, input *Lesson4_Units_SIInput) *Lesson4_Units_SIOutput {
	output := &Lesson4_Units_SIOutput{}
	_Lesson4_Units_SISetup(_ctx, input)
	_Lesson4_Units_SISteps(_ctx, input, output)
	_Lesson4_Units_SIAnalysis(_ctx, input, output)
	_Lesson4_Units_SIValidation(_ctx, input, output)
	return output
}

func Lesson4_Units_SIRunSteps(_ctx context.Context, input *Lesson4_Units_SIInput) *Lesson4_Units_SISOutput {
	soutput := &Lesson4_Units_SISOutput{}
	output := _Lesson4_Units_SIRun(_ctx, input)
	if err := inject.AssignSome(output, &soutput.Data); err != nil {
		panic(err)
	}
	if err := inject.AssignSome(output, &soutput.Outputs); err != nil {
		panic(err)
	}
	return soutput
}

func Lesson4_Units_SINew() interface{} {
	return &Lesson4_Units_SIElement{
		inject.CheckedRunner{
			RunFunc: func(_ctx context.Context, value inject.Value) (inject.Value, error) {
				input := &Lesson4_Units_SIInput{}
				if err := inject.Assign(value, input); err != nil {
					return nil, err
				}
				output := _Lesson4_Units_SIRun(_ctx, input)
				return inject.MakeValue(output), nil
			},
			In:  &Lesson4_Units_SIInput{},
			Out: &Lesson4_Units_SIOutput{},
		},
	}
}

var (
	_ = execute.MixInto
	_ = wtype.FALSE
	_ = wunit.Make_units
)

type Lesson4_Units_SIElement struct {
	inject.CheckedRunner
}

type Lesson4_Units_SIInput struct {
	MyMass wunit.Mass
}

type Lesson4_Units_SIOutput struct {
	RawMass     float64
	RawMassUnit string
	SIMass      float64
	SIMassUnit  string
}

type Lesson4_Units_SISOutput struct {
	Data struct {
		RawMass     float64
		RawMassUnit string
		SIMass      float64
		SIMassUnit  string
	}
	Outputs struct {
	}
}

func init() {
	if err := addComponent(component.Component{Name: "Lesson4_Units_SI",
		Constructor: Lesson4_Units_SINew,
		Desc: component.ComponentDesc{
			Desc: "",
			Path: "src/github.com/antha-lang/elements/AnthaAcademy/Lesson4_Units/C_units_SI.an",
			Params: []component.ParamDesc{
				{Name: "MyMass", Desc: "", Kind: "Parameters"},
				{Name: "RawMass", Desc: "", Kind: "Data"},
				{Name: "RawMassUnit", Desc: "", Kind: "Data"},
				{Name: "SIMass", Desc: "", Kind: "Data"},
				{Name: "SIMassUnit", Desc: "", Kind: "Data"},
			},
		},
	}); err != nil {
		panic(err)
	}
}
