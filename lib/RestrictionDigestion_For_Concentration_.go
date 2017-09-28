package lib

import (
	"context"
	"fmt"
	"github.com/antha-lang/antha/antha/AnthaStandardLibrary/Packages/text"
	"github.com/antha-lang/antha/antha/anthalib/mixer"
	"github.com/antha-lang/antha/antha/anthalib/wtype"
	"github.com/antha-lang/antha/antha/anthalib/wunit"
	"github.com/antha-lang/antha/component"
	"github.com/antha-lang/antha/execute"
	"github.com/antha-lang/antha/inject"
)

// Input parameters for this protocol (data)

// Physical Inputs to this protocol with types

// Physical outputs from this protocol with types

// Data which is returned from this protocol, and data types

func _RestrictionDigestion_For_ConcentrationRequirements() {}

// Conditions to run on startup
func _RestrictionDigestion_For_ConcentrationSetup(_ctx context.Context, _input *RestrictionDigestion_For_ConcentrationInput) {
}

// The core process for this protocol, with the steps to be performed
// for every input
func _RestrictionDigestion_For_ConcentrationSteps(_ctx context.Context, _input *RestrictionDigestion_For_ConcentrationInput, _output *RestrictionDigestion_For_ConcentrationOutput) {

	statii := make([]string, 0)

	samples := make([]*wtype.LHComponent, 0)
	waterSample := mixer.SampleForTotalVolume(_input.Water, _input.ReactionVolume)
	samples = append(samples, waterSample)

	// workout volume of buffer to add in SI units
	BufferVol := wunit.NewVolume(float64(_input.ReactionVolume.SIValue()/float64(_input.BufferConcX)), "l")
	statii = append(statii, fmt.Sprintln("buffer volume conversion:", _input.ReactionVolume.SIValue(), _input.BufferConcX, float64(_input.ReactionVolume.SIValue()/float64(_input.BufferConcX)), " Buffervol = ", BufferVol.SIValue()))
	bufferSample := mixer.Sample(_input.Buffer, BufferVol)
	samples = append(samples, bufferSample)

	if _input.BSAvol.Mvalue != 0 {
		bsaSample := mixer.Sample(_input.BSAoptional, _input.BSAvol)
		samples = append(samples, bsaSample)
	}

	_input.DNASolution.CName = _input.DNAName

	// work out necessary volume to add
	DNAVol, err := wunit.VolumeForTargetMass(_input.DNAMassperReaction, _input.DNAConc) //NewVolume(float64((DNAMassperReaction.SIValue()/DNAConc.SIValue())),"l")

	if err != nil {
		execute.Errorf(_ctx, err.Error())
	}

	statii = append(statii, fmt.Sprintln("DNA MAss to Volume conversion:", _input.DNAMassperReaction.SIValue(), _input.DNAConc.SIValue(), float64((_input.DNAMassperReaction.SIValue()/_input.DNAConc.SIValue())), "DNAVol =", DNAVol.SIValue()))
	statii = append(statii, fmt.Sprintln("DNAVOL", DNAVol.ToString()))
	dnaSample := mixer.Sample(_input.DNASolution, DNAVol)
	samples = append(samples, dnaSample)

	for k, enzyme := range _input.EnzSolutions {

		stockconcinUperul := _input.StockReConcinUperml[k] / 1000
		enzvoltoaddinul := _input.DesiredConcinUperml[k] / stockconcinUperul

		var enzvoltoadd wunit.Volume

		if float64(enzvoltoaddinul) < 0.5 {
			enzvoltoadd = wunit.NewVolume(float64(0.5), "ul")
		} else {
			enzvoltoadd = wunit.NewVolume(float64(enzvoltoaddinul), "ul")
		}
		enzyme.CName = _input.EnzymeNames[k]
		text.Print("adding enzyme"+_input.EnzymeNames[k], "to"+_input.DNAName)
		enzSample := mixer.Sample(enzyme, enzvoltoadd)
		enzSample.CName = _input.EnzymeNames[k]
		samples = append(samples, enzSample)
	}

	// incubate the reaction mixture
	r1 := execute.Incubate(_ctx, execute.MixInto(_ctx, _input.OutPlate, "", samples...), execute.IncubateOpt{
		Temp: _input.ReactionTemp,
		Time: _input.ReactionTime,
	})
	// inactivate
	_output.Reaction = execute.Incubate(_ctx, r1, execute.IncubateOpt{
		Temp: _input.InactivationTemp,
		Time: _input.InactivationTime,
	})
}

// Run after controls and a steps block are completed to
// post process any data and provide downstream results
func _RestrictionDigestion_For_ConcentrationAnalysis(_ctx context.Context, _input *RestrictionDigestion_For_ConcentrationInput, _output *RestrictionDigestion_For_ConcentrationOutput) {
}

// A block of tests to perform to validate that the sample was processed correctly
// Optionally, destructive tests can be performed to validate results on a
// dipstick basis
func _RestrictionDigestion_For_ConcentrationValidation(_ctx context.Context, _input *RestrictionDigestion_For_ConcentrationInput, _output *RestrictionDigestion_For_ConcentrationOutput) {
}
func _RestrictionDigestion_For_ConcentrationRun(_ctx context.Context, input *RestrictionDigestion_For_ConcentrationInput) *RestrictionDigestion_For_ConcentrationOutput {
	output := &RestrictionDigestion_For_ConcentrationOutput{}
	_RestrictionDigestion_For_ConcentrationSetup(_ctx, input)
	_RestrictionDigestion_For_ConcentrationSteps(_ctx, input, output)
	_RestrictionDigestion_For_ConcentrationAnalysis(_ctx, input, output)
	_RestrictionDigestion_For_ConcentrationValidation(_ctx, input, output)
	return output
}

func RestrictionDigestion_For_ConcentrationRunSteps(_ctx context.Context, input *RestrictionDigestion_For_ConcentrationInput) *RestrictionDigestion_For_ConcentrationSOutput {
	soutput := &RestrictionDigestion_For_ConcentrationSOutput{}
	output := _RestrictionDigestion_For_ConcentrationRun(_ctx, input)
	if err := inject.AssignSome(output, &soutput.Data); err != nil {
		panic(err)
	}
	if err := inject.AssignSome(output, &soutput.Outputs); err != nil {
		panic(err)
	}
	return soutput
}

func RestrictionDigestion_For_ConcentrationNew() interface{} {
	return &RestrictionDigestion_For_ConcentrationElement{
		inject.CheckedRunner{
			RunFunc: func(_ctx context.Context, value inject.Value) (inject.Value, error) {
				input := &RestrictionDigestion_For_ConcentrationInput{}
				if err := inject.Assign(value, input); err != nil {
					return nil, err
				}
				output := _RestrictionDigestion_For_ConcentrationRun(_ctx, input)
				return inject.MakeValue(output), nil
			},
			In:  &RestrictionDigestion_For_ConcentrationInput{},
			Out: &RestrictionDigestion_For_ConcentrationOutput{},
		},
	}
}

var (
	_ = execute.MixInto
	_ = wtype.FALSE
	_ = wunit.Make_units
)

type RestrictionDigestion_For_ConcentrationElement struct {
	inject.CheckedRunner
}

type RestrictionDigestion_For_ConcentrationInput struct {
	BSAoptional         *wtype.LHComponent
	BSAvol              wunit.Volume
	Buffer              *wtype.LHComponent
	BufferConcX         int
	DNAConc             wunit.Concentration
	DNAMassperReaction  wunit.Mass
	DNAName             string
	DNASolution         *wtype.LHComponent
	DesiredConcinUperml []int
	EnzSolutions        []*wtype.LHComponent
	EnzymeNames         []string
	InactivationTemp    wunit.Temperature
	InactivationTime    wunit.Time
	OutPlate            *wtype.LHPlate
	ReactionTemp        wunit.Temperature
	ReactionTime        wunit.Time
	ReactionVolume      wunit.Volume
	StockReConcinUperml []int
	Water               *wtype.LHComponent
}

type RestrictionDigestion_For_ConcentrationOutput struct {
	Reaction *wtype.LHComponent
	Status   string
}

type RestrictionDigestion_For_ConcentrationSOutput struct {
	Data struct {
		Status string
	}
	Outputs struct {
		Reaction *wtype.LHComponent
	}
}

func init() {
	if err := addComponent(component.Component{Name: "RestrictionDigestion_For_Concentration",
		Constructor: RestrictionDigestion_For_ConcentrationNew,
		Desc: component.ComponentDesc{
			Desc: "",
			Path: "src/github.com/antha-lang/elements/an/Restriction_Digestion_For_Concentration/element.an",
			Params: []component.ParamDesc{
				{Name: "BSAoptional", Desc: "", Kind: "Inputs"},
				{Name: "BSAvol", Desc: "", Kind: "Parameters"},
				{Name: "Buffer", Desc: "", Kind: "Inputs"},
				{Name: "BufferConcX", Desc: "", Kind: "Parameters"},
				{Name: "DNAConc", Desc: "", Kind: "Parameters"},
				{Name: "DNAMassperReaction", Desc: "", Kind: "Parameters"},
				{Name: "DNAName", Desc: "", Kind: "Parameters"},
				{Name: "DNASolution", Desc: "", Kind: "Inputs"},
				{Name: "DesiredConcinUperml", Desc: "", Kind: "Parameters"},
				{Name: "EnzSolutions", Desc: "", Kind: "Inputs"},
				{Name: "EnzymeNames", Desc: "", Kind: "Parameters"},
				{Name: "InactivationTemp", Desc: "", Kind: "Parameters"},
				{Name: "InactivationTime", Desc: "", Kind: "Parameters"},
				{Name: "OutPlate", Desc: "", Kind: "Inputs"},
				{Name: "ReactionTemp", Desc: "", Kind: "Parameters"},
				{Name: "ReactionTime", Desc: "", Kind: "Parameters"},
				{Name: "ReactionVolume", Desc: "", Kind: "Parameters"},
				{Name: "StockReConcinUperml", Desc: "", Kind: "Parameters"},
				{Name: "Water", Desc: "", Kind: "Inputs"},
				{Name: "Reaction", Desc: "", Kind: "Outputs"},
				{Name: "Status", Desc: "", Kind: "Data"},
			},
		},
	}); err != nil {
		panic(err)
	}
}
