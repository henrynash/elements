// protocol Lessonfor running pcr for one sample using a mastermix
package lib

import (
	"context"
	"fmt"
	"github.com/antha-lang/antha/antha/AnthaStandardLibrary/Packages/enzymes"
	"github.com/antha-lang/antha/antha/anthalib/mixer"
	"github.com/antha-lang/antha/antha/anthalib/wtype"
	"github.com/antha-lang/antha/antha/anthalib/wunit"
	"github.com/antha-lang/antha/component"
	"github.com/antha-lang/antha/execute"
	"github.com/antha-lang/antha/inject"
	"strings"
)

// Input parameters for this protocol Lesson(data)

// PCRprep parameters:

// Total volume for a single reaction; the reaction will be topped up with ReactionBuffer (usually water) to reach this volume

/*
	// let's be ambitious and try this as part of type polymerase Polymeraseconc Volume

	//Templatetype string  // e.g. colony, genomic, pure plasmid... will effect efficiency. We could get more sophisticated here later on...
	//FullTemplatesequence string // better to use Sid's type system here after proof of concept
	//FullTemplatelength int	// clearly could be calculated from the sequence... Sid will have a method to do this already so check!
	//TargetTemplatesequence string // better to use Sid's type system here after proof of concept
	//TargetTemplatelengthinBP int
*/
// Reaction parameters: (could be a entered as thermocycle parameters type possibly?)

//Denaturationtemp Temperature

// Should be calculated from primer and template binding
// should be calculated from template length and polymerase rate

// Data which is returned from this protocol, and data types

// Physical Inputs to this protocol Lessonwith types

// Buffer to use to top up the reaction to TotalReactionVolume. Typical buffer for this would be water.

// Physical outputs from this protocol Lessonwith types

func _Lesson0_PCR_MMXRequirements() {
}

// Conditions to run on startup
func _Lesson0_PCR_MMXSetup(_ctx context.Context, _input *Lesson0_PCR_MMXInput) {
}

// The core process for this protocol, with the steps to be performed
// for every input
func _Lesson0_PCR_MMXSteps(_ctx context.Context, _input *Lesson0_PCR_MMXInput, _output *Lesson0_PCR_MMXOutput) {

	// rename components

	_input.Template.CName = _input.TemplateName
	_input.FwdPrimer.CName = _input.FwdPrimerName
	_input.RevPrimer.CName = _input.RevPrimerName

	var allVolumes []wunit.Volume

	allVolumes = append(allVolumes, _input.MasterMixVolume, _input.Templatevolume)

	if !_input.PrimersalreadyAddedtoMasterMix {
		allVolumes = append(allVolumes, _input.FwdPrimerVol, _input.RevPrimerVol)
	}

	if !_input.PolymeraseAlreadyaddedtoMastermix {
		allVolumes = append(allVolumes, _input.PolymeraseVolume)
	}

	// calculate volume of water to add
	waterVol := wunit.SubtractVolumes(_input.TotalReactionVolume, allVolumes)

	var mastermix *wtype.LHComponent
	// Top up with reaction buffer if necessary.
	if waterVol.GreaterThan(wunit.NewVolume(0.5, "ul")) {
		waterSample := mixer.Sample(_input.ReactionBuffer, waterVol)
		mastermix = execute.MixInto(_ctx, _input.OutPlate, _input.WellPosition, waterSample)
	}

	// Make a mastermix

	mmxSample := mixer.Sample(_input.MasterMix, _input.MasterMixVolume)

	// pipette out to make mastermix
	if mastermix != nil {
		mastermix = execute.Mix(_ctx, mastermix, mmxSample)
	} else {
		mastermix = execute.MixInto(_ctx, _input.OutPlate, _input.WellPosition, mmxSample)
	}

	// rest samples to zero
	samples := make([]*wtype.LHComponent, 0)

	// if this is false do stuff inside {}

	// add primers

	if !_input.PrimersalreadyAddedtoMasterMix {
		FwdPrimerSample := mixer.Sample(_input.FwdPrimer, _input.FwdPrimerVol)
		samples = append(samples, FwdPrimerSample)
		RevPrimerSample := mixer.Sample(_input.RevPrimer, _input.RevPrimerVol)
		samples = append(samples, RevPrimerSample)
	}

	// add template
	templateSample := mixer.Sample(_input.Template, _input.Templatevolume)
	samples = append(samples, templateSample)

	for j := range samples {
		if !_input.PolymeraseAlreadyaddedtoMastermix && j == len(samples)-1 {
			samples[j].Type = wtype.LTPostMix
		}
		mastermix = execute.Mix(_ctx, mastermix, samples[j])
	}
	reaction := mastermix

	// this needs to go after an initial denaturation!
	if !_input.PolymeraseAlreadyaddedtoMastermix {

		polySample := mixer.Sample(_input.PCRPolymerase, _input.PolymeraseVolume)
		polySample.Type = wtype.LTPostMix

		reaction = execute.Mix(_ctx, reaction, polySample)
	}

	// thermocycle parameters called from enzyme lookup:

	polymerase := _input.PCRPolymerase.CName

	extensionTemp := enzymes.DNApolymerasetemps[polymerase]["extensiontemp"]
	meltingTemp := enzymes.DNApolymerasetemps[polymerase]["meltingtemp"]

	var pcrSteps []string

	initialDenat := fmt.Sprint("Initial Denaturation: ", meltingTemp.ToString(), " for ", _input.InitDenaturationtime.ToString())

	cycles := fmt.Sprint(_input.Numberofcycles, " cycles of : ")

	spacer := "***"

	denat := fmt.Sprint("Denature: ", meltingTemp.ToString(), " for ", _input.Denaturationtime.ToString())

	anneal := fmt.Sprint("Anneal: ", _input.AnnealingTemp.ToString(), " for ", _input.Annealingtime.ToString())

	extend := fmt.Sprint("Extension: ", extensionTemp.ToString(), " for ", _input.Extensiontime.ToString())

	spacer = "***"

	finalExtension := fmt.Sprint(" Then Final Extension: ", extensionTemp.ToString(), " for ", _input.Finalextensiontime.ToString())

	// all done
	_output.Reaction = reaction //r1

	_output.Reaction.CName = _input.ReactionName

	message := "Put Reactions in ThermoCycler with following cycle conditions. Return to deck once PCR has finished if running DNA_Gel"

	pcrSteps = append(pcrSteps, initialDenat, cycles, spacer, denat, anneal, extend, spacer, finalExtension, message)

	thermocycleMessage := strings.Join(pcrSteps, ";")

	_output.Reaction = execute.MixerPrompt(_ctx, _output.Reaction, thermocycleMessage)
}

// Run after controls and a steps block are completed to
// post process any data and provide downstream results
func _Lesson0_PCR_MMXAnalysis(_ctx context.Context, _input *Lesson0_PCR_MMXInput, _output *Lesson0_PCR_MMXOutput) {
}

// A block of tests to perform to validate that the sample was processed correctly
// Optionally, destructive tests can be performed to validate results on a
// dipstick basis
func _Lesson0_PCR_MMXValidation(_ctx context.Context, _input *Lesson0_PCR_MMXInput, _output *Lesson0_PCR_MMXOutput) {
}
func _Lesson0_PCR_MMXRun(_ctx context.Context, input *Lesson0_PCR_MMXInput) *Lesson0_PCR_MMXOutput {
	output := &Lesson0_PCR_MMXOutput{}
	_Lesson0_PCR_MMXSetup(_ctx, input)
	_Lesson0_PCR_MMXSteps(_ctx, input, output)
	_Lesson0_PCR_MMXAnalysis(_ctx, input, output)
	_Lesson0_PCR_MMXValidation(_ctx, input, output)
	return output
}

func Lesson0_PCR_MMXRunSteps(_ctx context.Context, input *Lesson0_PCR_MMXInput) *Lesson0_PCR_MMXSOutput {
	soutput := &Lesson0_PCR_MMXSOutput{}
	output := _Lesson0_PCR_MMXRun(_ctx, input)
	if err := inject.AssignSome(output, &soutput.Data); err != nil {
		panic(err)
	}
	if err := inject.AssignSome(output, &soutput.Outputs); err != nil {
		panic(err)
	}
	return soutput
}

func Lesson0_PCR_MMXNew() interface{} {
	return &Lesson0_PCR_MMXElement{
		inject.CheckedRunner{
			RunFunc: func(_ctx context.Context, value inject.Value) (inject.Value, error) {
				input := &Lesson0_PCR_MMXInput{}
				if err := inject.Assign(value, input); err != nil {
					return nil, err
				}
				output := _Lesson0_PCR_MMXRun(_ctx, input)
				return inject.MakeValue(output), nil
			},
			In:  &Lesson0_PCR_MMXInput{},
			Out: &Lesson0_PCR_MMXOutput{},
		},
	}
}

var (
	_ = execute.MixInto
	_ = wtype.FALSE
	_ = wunit.Make_units
)

type Lesson0_PCR_MMXElement struct {
	inject.CheckedRunner
}

type Lesson0_PCR_MMXInput struct {
	AnnealingTemp                     wunit.Temperature
	Annealingtime                     wunit.Time
	Denaturationtime                  wunit.Time
	Extensiontime                     wunit.Time
	Finalextensiontime                wunit.Time
	FwdPrimer                         *wtype.LHComponent
	FwdPrimerName                     string
	FwdPrimerSeq                      wtype.DNASequence
	FwdPrimerVol                      wunit.Volume
	InitDenaturationtime              wunit.Time
	MasterMix                         *wtype.LHComponent
	MasterMixVolume                   wunit.Volume
	Numberofcycles                    int
	OutPlate                          *wtype.LHPlate
	PCRPolymerase                     *wtype.LHComponent
	PolymeraseAlreadyaddedtoMastermix bool
	PolymeraseVolume                  wunit.Volume
	PrimersalreadyAddedtoMasterMix    bool
	ReactionBuffer                    *wtype.LHComponent
	ReactionName                      string
	RevPrimer                         *wtype.LHComponent
	RevPrimerName                     string
	RevPrimerSeq                      wtype.DNASequence
	RevPrimerVol                      wunit.Volume
	Targetsequence                    wtype.DNASequence
	Template                          *wtype.LHComponent
	TemplateName                      string
	Templatevolume                    wunit.Volume
	TotalReactionVolume               wunit.Volume
	WellPosition                      string
}

type Lesson0_PCR_MMXOutput struct {
	FWDPrimerBindingSiteinTemplate int
	Reaction                       *wtype.LHComponent
	RevPrimerBindingSiteinTemplate int
}

type Lesson0_PCR_MMXSOutput struct {
	Data struct {
		FWDPrimerBindingSiteinTemplate int
		RevPrimerBindingSiteinTemplate int
	}
	Outputs struct {
		Reaction *wtype.LHComponent
	}
}

func init() {
	if err := addComponent(component.Component{Name: "Lesson0_PCR_MMX",
		Constructor: Lesson0_PCR_MMXNew,
		Desc: component.ComponentDesc{
			Desc: "protocol Lessonfor running pcr for one sample using a mastermix\n",
			Path: "src/github.com/antha-lang/elements/AnthaAcademy/Lesson0_Examples/MakeMasterMix_PCR/PCR.an",
			Params: []component.ParamDesc{
				{Name: "AnnealingTemp", Desc: "Should be calculated from primer and template binding\n", Kind: "Parameters"},
				{Name: "Annealingtime", Desc: "Denaturationtemp Temperature\n", Kind: "Parameters"},
				{Name: "Denaturationtime", Desc: "", Kind: "Parameters"},
				{Name: "Extensiontime", Desc: "should be calculated from template length and polymerase rate\n", Kind: "Parameters"},
				{Name: "Finalextensiontime", Desc: "", Kind: "Parameters"},
				{Name: "FwdPrimer", Desc: "", Kind: "Inputs"},
				{Name: "FwdPrimerName", Desc: "", Kind: "Parameters"},
				{Name: "FwdPrimerSeq", Desc: "", Kind: "Parameters"},
				{Name: "FwdPrimerVol", Desc: "", Kind: "Parameters"},
				{Name: "InitDenaturationtime", Desc: "", Kind: "Parameters"},
				{Name: "MasterMix", Desc: "", Kind: "Inputs"},
				{Name: "MasterMixVolume", Desc: "", Kind: "Parameters"},
				{Name: "Numberofcycles", Desc: "\t\t// let's be ambitious and try this as part of type polymerase Polymeraseconc Volume\n\n\t\t//Templatetype string  // e.g. colony, genomic, pure plasmid... will effect efficiency. We could get more sophisticated here later on...\n\t\t//FullTemplatesequence string // better to use Sid's type system here after proof of concept\n\t\t//FullTemplatelength int\t// clearly could be calculated from the sequence... Sid will have a method to do this already so check!\n\t\t//TargetTemplatesequence string // better to use Sid's type system here after proof of concept\n\t\t//TargetTemplatelengthinBP int\n\nReaction parameters: (could be a entered as thermocycle parameters type possibly?)\n", Kind: "Parameters"},
				{Name: "OutPlate", Desc: "", Kind: "Inputs"},
				{Name: "PCRPolymerase", Desc: "", Kind: "Inputs"},
				{Name: "PolymeraseAlreadyaddedtoMastermix", Desc: "", Kind: "Parameters"},
				{Name: "PolymeraseVolume", Desc: "", Kind: "Parameters"},
				{Name: "PrimersalreadyAddedtoMasterMix", Desc: "", Kind: "Parameters"},
				{Name: "ReactionBuffer", Desc: "Buffer to use to top up the reaction to TotalReactionVolume. Typical buffer for this would be water.\n", Kind: "Inputs"},
				{Name: "ReactionName", Desc: "", Kind: "Parameters"},
				{Name: "RevPrimer", Desc: "", Kind: "Inputs"},
				{Name: "RevPrimerName", Desc: "", Kind: "Parameters"},
				{Name: "RevPrimerSeq", Desc: "", Kind: "Parameters"},
				{Name: "RevPrimerVol", Desc: "", Kind: "Parameters"},
				{Name: "Targetsequence", Desc: "", Kind: "Parameters"},
				{Name: "Template", Desc: "", Kind: "Inputs"},
				{Name: "TemplateName", Desc: "", Kind: "Parameters"},
				{Name: "Templatevolume", Desc: "", Kind: "Parameters"},
				{Name: "TotalReactionVolume", Desc: "Total volume for a single reaction; the reaction will be topped up with ReactionBuffer (usually water) to reach this volume\n", Kind: "Parameters"},
				{Name: "WellPosition", Desc: "", Kind: "Parameters"},
				{Name: "FWDPrimerBindingSiteinTemplate", Desc: "", Kind: "Data"},
				{Name: "Reaction", Desc: "", Kind: "Outputs"},
				{Name: "RevPrimerBindingSiteinTemplate", Desc: "", Kind: "Data"},
			},
		},
	}); err != nil {
		panic(err)
	}
}

/*type Polymerase struct {
	LHComponent
	Rate_BPpers float64
	Fidelity_errorrate float64 // could dictate how many colonies are checked in validation!
	Extensiontemp Temperature
	Hotstart bool
	StockConcentration Concentration // this is normally in U?
	TargetConcentration Concentration
	// this is also a glycerol solution rather than a watersolution!
}
*/
