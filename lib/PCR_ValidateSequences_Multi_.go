// Perform multiple PCR reactions with common default parameters using a mastermix
package lib

import (
	"context"
	"github.com/antha-lang/antha/antha/anthalib/wtype"
	"github.com/antha-lang/antha/antha/anthalib/wunit"
	"github.com/antha-lang/antha/component"
	"github.com/antha-lang/antha/execute"
	"github.com/antha-lang/antha/inject"
)

// Input parameters for this protocol (data)

// map of which reaction uses which template e.g. ["left homology arm"]:"templatename".
// A "default" may also be specified which will be used for any reaction which has no entry specified.

// map of which reaction uses which primer pair e.g. ["left homology arm"]:"fwdprimer","revprimer".
// A "default" may also be specified which will be used for any reaction which has no entry specified.

// Default behaviour will randomise the output order of pcr reactions.
// To specify an order add the list of PCRReactions here in order here.

// Volume of template in each reaction

// Volume of each primer to add. Will only be used if PrimersalreadyAddedtoMasterMix is not selected.

// Volume of polymerase enzyme to add per reaction. Will only be used if PolymeraseAlreadyaddedtoMastermix is not selected.

// Volume of mastermix to add to the reaction.

// Total volume for a single reaction; the reaction will be topped up with TopUpSolution (usually water) to reach this volume.

// Select this if the primers have already been added to the mastermix.
// If this is selected no primers will be added to any reactions.
// Should only be used if all reactions share the same primers.

// Select this if the polymerase has already been added to the mastermix.

// Data which is returned from this protocol, and data types

// a reaction by reaction description of all sets of conditions suggested for each reaction.
// amplicon expected for each reaction

// Physical Inputs to this protocol with types

// Actual FWD primer component type to use. e.g. dna_part. Will only be used if PrimersalreadyAddedtoMasterMix is not selected.

// Actual REV primer component type to use. e.g. dna_part. Will only be used if PrimersalreadyAddedtoMasterMix is not selected.

// Actual Template type to use. e.g. dna_part, culture

// Valid options are Q5Polymerase and Taq. To make a custom Polymerase use the NewLHComponent element and wire in here.
// This input will only be used if PolymeraseAlreadyaddedtoMastermix is not selected.

// Use the MasterMixMaker element to make the mastermix and wire it in here.

// Buffer to use to top up the reaction to TotalReactionVolume. Typical solution for this would be water.

// Type of plate to use for the reaction.
// Recommended plates: 96well plate (pcrplate_skirted) (Bio-Rad, 96 well hard shell skirted plate, Cat No #HSP9901)
// 96 well semi-skirted pcr plate (pcrplate) (Bio-Rad)

// Physical outputs from this protocol with types

// The PCR reaction products as a slice.

// The PCR reaction products in the form of a map of components.

func _PCR_ValidateSequences_MultiRequirements() {
}

// Conditions to run on startup
func _PCR_ValidateSequences_MultiSetup(_ctx context.Context, _input *PCR_ValidateSequences_MultiInput) {
}

// The core process for this protocol, with the steps to be performed
// for every input
func _PCR_ValidateSequences_MultiSteps(_ctx context.Context, _input *PCR_ValidateSequences_MultiInput, _output *PCR_ValidateSequences_MultiOutput) {

	// set up a counter to use as an index for increasing well position
	var counter int

	// set up some empty slices to fill as we iterate through the reactions
	_output.Reactions = make([]*wtype.LHComponent, 0)
	volumes := make([]wunit.Volume, 0)
	welllocations := make([]string, 0)
	// initialise map
	_output.ReactionMap = make(map[string]*wtype.LHComponent)

	// To allow the use of specifying defaults in either the template map or the primer map,
	// we first evaluate which map has more entries and make a slice of reactions to iterate through based on that.
	var reactions []string

	if len(_input.SpecifyReactionOrder) == 0 {
		if len(_input.Reactiontotemplate) >= len(_input.Reactiontoprimerpair) {
			for reactionname := range _input.Reactiontotemplate {
				reactions = append(reactions, reactionname)
			}
		} else {
			for reactionname := range _input.Reactiontoprimerpair {
				reactions = append(reactions, reactionname)
			}
		}
	} else {
		if len(_input.Reactiontotemplate) >= len(_input.Reactiontoprimerpair) && len(_input.SpecifyReactionOrder) == len(_input.Reactiontotemplate) {
			for _, reaction := range _input.SpecifyReactionOrder {
				if _, found := _input.Reactiontotemplate[reaction]; !found {
					execute.Errorf(_ctx, "Reaction %s specified in SpecifyReactionOrder but not present in the Reactiontotemplate which sets the reaction list.", reaction)
				}
			}
			reactions = _input.SpecifyReactionOrder
		} else if len(_input.Reactiontotemplate) <= len(_input.Reactiontoprimerpair) && len(_input.SpecifyReactionOrder) == len(_input.Reactiontoprimerpair) {
			for _, reaction := range _input.SpecifyReactionOrder {
				if _, found := _input.Reactiontoprimerpair[reaction]; !found {
					execute.Errorf(_ctx, "Reaction %s specified in SpecifyReactionOrder but not present in Reactiontoprimerpair which sets the reaction list.", reaction)
				}
			}
			reactions = _input.SpecifyReactionOrder
		} else {
			execute.Errorf(_ctx, "Reaction order specified but the number of PCRReactions (%d) in this list differs to the length of both Reactiontotemplate map (%d) and Reactiontoprimerpair map (%d). Please ensure the list is equal to one of these (the other may contain defaults) or remove the SpecifyReactionOrder list. ", len(_input.SpecifyReactionOrder), len(_input.Reactiontotemplate), len(_input.Reactiontoprimerpair))
		}
	}

	for _, reactionname := range reactions {

		// look up template from map
		var template wtype.DNASequence

		if templateSeq, found := _input.Reactiontotemplate[reactionname]; found {
			template = templateSeq
		} else if templateSeq, found := _input.Reactiontotemplate["default"]; found {
			template = templateSeq
		} else {
			execute.Errorf(_ctx, `No template set for %s and no "default" primers set`, reactionname)
		}

		// look up primers from map
		var fwdPrimer wtype.DNASequence
		var revPrimer wtype.DNASequence

		if primers, found := _input.Reactiontoprimerpair[reactionname]; found {
			fwdPrimer, revPrimer = primers[0], primers[1]
		} else if primers, found := _input.Reactiontoprimerpair["default"]; found {
			fwdPrimer, revPrimer = primers[0], primers[1]
		} else {
			execute.Errorf(_ctx, `No primers set for %s and no "default" primers set`, reactionname)
		}

		// use counter to find next available well position in plate

		var allwellpositionsforplate []string

		allwellpositionsforplate = _input.Plate.AllWellPositions(wtype.BYCOLUMN)

		wellposition := allwellpositionsforplate[counter]

		// Run PCR_vol element
		result := PCR_mmx_ValidateSequencesRunSteps(_ctx, &PCR_mmx_ValidateSequencesInput{MasterMixVolume: _input.DefaultMasterMixVolume,
			PrimersalreadyAddedtoMasterMix:    _input.PrimersalreadyAddedtoMasterMix,
			PolymeraseAlreadyaddedtoMastermix: _input.PolymeraseAlreadyaddedtoMastermix,
			FwdPrimerSeq:                      fwdPrimer,
			RevPrimerSeq:                      revPrimer,
			TemplateSequence:                  template,
			ReactionName:                      reactionname,
			FwdPrimerVol:                      _input.DefaultPrimerVolume,
			RevPrimerVol:                      _input.DefaultPrimerVolume,
			PolymeraseVolume:                  _input.DefaultPolymeraseVolume,
			Templatevolume:                    _input.DefaultTemplateVol,
			NumberOfCycles:                    30,
			InitDenaturationTime:              wunit.NewTime(30, "s"),
			DenaturationTime:                  wunit.NewTime(5, "s"),
			AnnealingTime:                     wunit.NewTime(10, "s"),
			FinalExtensionTime:                wunit.NewTime(180, "s"),
			OptionalWellPosition:              wellposition,
			TotalReactionVolume:               _input.TotalReactionVolume,

			FwdPrimer:      _input.FwdPrimertype,
			RevPrimer:      _input.RevPrimertype,
			PCRPolymerase:  _input.DefaultPolymerase,
			MasterMix:      _input.MasterMix,
			Template:       _input.Templatetype,
			ReactionBuffer: _input.TopUpSolution,
			OutPlate:       _input.Plate},
		)

		// add result to reactions slice
		_output.Reactions = append(_output.Reactions, result.Outputs.Reaction)
		volumes = append(volumes, result.Outputs.Reaction.Volume())
		welllocations = append(welllocations, wellposition)
		_output.ReactionMap[reactionname] = result.Outputs.Reaction
		_output.Amplicons[reactionname] = result.Data.Amplicon
		_output.ThermoCycleConditionsUsed[reactionname] = result.Data.ThermoCycleConditionsUsed
		// increase counter by 1 ready for next iteration of loop
		counter++

	}

	//MixerPrompt(Reactions[0], "Put Reactions in ThermoCylcer and return to deck once PCR has finished if running DNA_Gel")

	// once all values of loop have been completed, export the plate contents as a csv file
	wtype.ExportPlateCSV(_input.Projectname+".csv", _input.Plate, _input.Projectname+"outputPlate", welllocations, _output.Reactions, volumes)

}

// Run after controls and a steps block are completed to
// post process any data and provide downstream results
func _PCR_ValidateSequences_MultiAnalysis(_ctx context.Context, _input *PCR_ValidateSequences_MultiInput, _output *PCR_ValidateSequences_MultiOutput) {
}

// A block of tests to perform to validate that the sample was processed correctly
// Optionally, destructive tests can be performed to validate results on a
// dipstick basis
func _PCR_ValidateSequences_MultiValidation(_ctx context.Context, _input *PCR_ValidateSequences_MultiInput, _output *PCR_ValidateSequences_MultiOutput) {
}
func _PCR_ValidateSequences_MultiRun(_ctx context.Context, input *PCR_ValidateSequences_MultiInput) *PCR_ValidateSequences_MultiOutput {
	output := &PCR_ValidateSequences_MultiOutput{}
	_PCR_ValidateSequences_MultiSetup(_ctx, input)
	_PCR_ValidateSequences_MultiSteps(_ctx, input, output)
	_PCR_ValidateSequences_MultiAnalysis(_ctx, input, output)
	_PCR_ValidateSequences_MultiValidation(_ctx, input, output)
	return output
}

func PCR_ValidateSequences_MultiRunSteps(_ctx context.Context, input *PCR_ValidateSequences_MultiInput) *PCR_ValidateSequences_MultiSOutput {
	soutput := &PCR_ValidateSequences_MultiSOutput{}
	output := _PCR_ValidateSequences_MultiRun(_ctx, input)
	if err := inject.AssignSome(output, &soutput.Data); err != nil {
		panic(err)
	}
	if err := inject.AssignSome(output, &soutput.Outputs); err != nil {
		panic(err)
	}
	return soutput
}

func PCR_ValidateSequences_MultiNew() interface{} {
	return &PCR_ValidateSequences_MultiElement{
		inject.CheckedRunner{
			RunFunc: func(_ctx context.Context, value inject.Value) (inject.Value, error) {
				input := &PCR_ValidateSequences_MultiInput{}
				if err := inject.Assign(value, input); err != nil {
					return nil, err
				}
				output := _PCR_ValidateSequences_MultiRun(_ctx, input)
				return inject.MakeValue(output), nil
			},
			In:  &PCR_ValidateSequences_MultiInput{},
			Out: &PCR_ValidateSequences_MultiOutput{},
		},
	}
}

var (
	_ = execute.MixInto
	_ = wtype.FALSE
	_ = wunit.Make_units
)

type PCR_ValidateSequences_MultiElement struct {
	inject.CheckedRunner
}

type PCR_ValidateSequences_MultiInput struct {
	DefaultMasterMixVolume            wunit.Volume
	DefaultPolymerase                 *wtype.LHComponent
	DefaultPolymeraseVolume           wunit.Volume
	DefaultPrimerVolume               wunit.Volume
	DefaultTemplateVol                wunit.Volume
	FwdPrimertype                     *wtype.LHComponent
	MasterMix                         *wtype.LHComponent
	Plate                             *wtype.LHPlate
	PolymeraseAlreadyaddedtoMastermix bool
	PrimersalreadyAddedtoMasterMix    bool
	Projectname                       string
	Reactiontoprimerpair              map[string][2]wtype.DNASequence
	Reactiontotemplate                map[string]wtype.DNASequence
	RevPrimertype                     *wtype.LHComponent
	SpecifyReactionOrder              []string
	Templatetype                      *wtype.LHComponent
	TopUpSolution                     *wtype.LHComponent
	TotalReactionVolume               wunit.Volume
}

type PCR_ValidateSequences_MultiOutput struct {
	Amplicons                 map[string]wtype.DNASequence
	ReactionMap               map[string]*wtype.LHComponent
	Reactions                 []*wtype.LHComponent
	ThermoCycleConditionsUsed map[string]string
}

type PCR_ValidateSequences_MultiSOutput struct {
	Data struct {
		Amplicons                 map[string]wtype.DNASequence
		ThermoCycleConditionsUsed map[string]string
	}
	Outputs struct {
		ReactionMap map[string]*wtype.LHComponent
		Reactions   []*wtype.LHComponent
	}
}

func init() {
	if err := addComponent(component.Component{Name: "PCR_ValidateSequences_Multi",
		Constructor: PCR_ValidateSequences_MultiNew,
		Desc: component.ComponentDesc{
			Desc: "Perform multiple PCR reactions with common default parameters using a mastermix\n",
			Path: "src/github.com/antha-lang/elements/an/PCR_ValidateSequences_Multi/element.an",
			Params: []component.ParamDesc{
				{Name: "DefaultMasterMixVolume", Desc: "Volume of mastermix to add to the reaction.\n", Kind: "Parameters"},
				{Name: "DefaultPolymerase", Desc: "Valid options are Q5Polymerase and Taq. To make a custom Polymerase use the NewLHComponent element and wire in here.\nThis input will only be used if PolymeraseAlreadyaddedtoMastermix is not selected.\n", Kind: "Inputs"},
				{Name: "DefaultPolymeraseVolume", Desc: "Volume of polymerase enzyme to add per reaction. Will only be used if PolymeraseAlreadyaddedtoMastermix is not selected.\n", Kind: "Parameters"},
				{Name: "DefaultPrimerVolume", Desc: "Volume of each primer to add. Will only be used if PrimersalreadyAddedtoMasterMix is not selected.\n", Kind: "Parameters"},
				{Name: "DefaultTemplateVol", Desc: "Volume of template in each reaction\n", Kind: "Parameters"},
				{Name: "FwdPrimertype", Desc: "Actual FWD primer component type to use. e.g. dna_part. Will only be used if PrimersalreadyAddedtoMasterMix is not selected.\n", Kind: "Inputs"},
				{Name: "MasterMix", Desc: "Use the MasterMixMaker element to make the mastermix and wire it in here.\n", Kind: "Inputs"},
				{Name: "Plate", Desc: "Type of plate to use for the reaction.\nRecommended plates: 96well plate (pcrplate_skirted) (Bio-Rad, 96 well hard shell skirted plate, Cat No #HSP9901)\n96 well semi-skirted pcr plate (pcrplate) (Bio-Rad)\n", Kind: "Inputs"},
				{Name: "PolymeraseAlreadyaddedtoMastermix", Desc: "Select this if the polymerase has already been added to the mastermix.\n", Kind: "Parameters"},
				{Name: "PrimersalreadyAddedtoMasterMix", Desc: "Select this if the primers have already been added to the mastermix.\nIf this is selected no primers will be added to any reactions.\nShould only be used if all reactions share the same primers.\n", Kind: "Parameters"},
				{Name: "Projectname", Desc: "", Kind: "Parameters"},
				{Name: "Reactiontoprimerpair", Desc: "map of which reaction uses which primer pair e.g. [\"left homology arm\"]:\"fwdprimer\",\"revprimer\".\nA \"default\" may also be specified which will be used for any reaction which has no entry specified.\n", Kind: "Parameters"},
				{Name: "Reactiontotemplate", Desc: "map of which reaction uses which template e.g. [\"left homology arm\"]:\"templatename\".\nA \"default\" may also be specified which will be used for any reaction which has no entry specified.\n", Kind: "Parameters"},
				{Name: "RevPrimertype", Desc: "Actual REV primer component type to use. e.g. dna_part. Will only be used if PrimersalreadyAddedtoMasterMix is not selected.\n", Kind: "Inputs"},
				{Name: "SpecifyReactionOrder", Desc: "Default behaviour will randomise the output order of pcr reactions.\nTo specify an order add the list of PCRReactions here in order here.\n", Kind: "Parameters"},
				{Name: "Templatetype", Desc: "Actual Template type to use. e.g. dna_part, culture\n", Kind: "Inputs"},
				{Name: "TopUpSolution", Desc: "Buffer to use to top up the reaction to TotalReactionVolume. Typical solution for this would be water.\n", Kind: "Inputs"},
				{Name: "TotalReactionVolume", Desc: "Total volume for a single reaction; the reaction will be topped up with TopUpSolution (usually water) to reach this volume.\n", Kind: "Parameters"},
				{Name: "Amplicons", Desc: "amplicon expected for each reaction\n", Kind: "Data"},
				{Name: "ReactionMap", Desc: "The PCR reaction products in the form of a map of components.\n", Kind: "Outputs"},
				{Name: "Reactions", Desc: "The PCR reaction products as a slice.\n", Kind: "Outputs"},
				{Name: "ThermoCycleConditionsUsed", Desc: "a reaction by reaction description of all sets of conditions suggested for each reaction.\n", Kind: "Data"},
			},
		},
	}); err != nil {
		panic(err)
	}
}
