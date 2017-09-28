// Perform multiple PCR reactions with common default parameters
package lib

import (
	"context"
	"fmt"
	"github.com/antha-lang/antha/antha/anthalib/wtype"
	"github.com/antha-lang/antha/antha/anthalib/wunit"
	"github.com/antha-lang/antha/component"
	"github.com/antha-lang/antha/execute"
	"github.com/antha-lang/antha/inject"
	"github.com/antha-lang/antha/inventory"
)

// Input parameters for this protocol Lesson(data)

// PCRprep parameters

// map of which reaction uses which template e.g. ["left homology arm"]:"templatename"
// map of which reaction uses which primer pair e.g. ["left homology arm"]:"fwdprimer","revprimer"
// Volume of template in each reaction
// e.g. for  10X Q5 buffer this would be 10
// Volume for each reaction

// look up table of additives to volumes of each additive; e.g. ["DMSO"]:"3ul"

// Data which is returned from this protocol, and data types

// return an error message if an error is encountered

// Physical Inputs to this protocol Lessonwith types

// Physical outputs from this protocol Lessonwith types

func _Lesson0_PCR_MultiRequirements() {
}

// Conditions to run on startup
func _Lesson0_PCR_MultiSetup(_ctx context.Context, _input *Lesson0_PCR_MultiInput) {
}

// The core process for this protocol, with the steps to be performed
// for every input
func _Lesson0_PCR_MultiSteps(_ctx context.Context, _input *Lesson0_PCR_MultiInput, _output *Lesson0_PCR_MultiOutput) {

	// set up a counter to use as an index for increasing well position
	var counter int

	// set up some empty slices to fill as we iterate through the reactions
	_output.Reactions = make([]*wtype.LHComponent, 0)
	volumes := make([]wunit.Volume, 0)
	welllocations := make([]string, 0)
	// initialise map
	_output.ReactionMap = make(map[string]*wtype.LHComponent)

	additives := make([]*wtype.LHComponent, 0)
	additivevolumes := make([]wunit.Volume, 0)

	// get additive info
	for additive, volume := range _input.AdditiveToAdditiveVolume {
		additivecomponent, err := inventory.NewComponent(_ctx, additive)

		if err == inventory.ErrUnknownType {
			// if not found in factory use dmso as the base liquid
			// handling type and change name to additivename
			// specified
			additivecomponent, err = inventory.NewComponent(_ctx, "DMSO")
			if err == nil {
				additivecomponent.CName = additive
			}
		}
		if err != nil {
			execute.Errorf(_ctx, "cannot make component: %s", err)
		}
		additives = append(additives, additivecomponent)
		additivevolumes = append(additivevolumes, volume)
	}

	// range through the Reaction to template map
	for reactionname, templatename := range _input.Reactiontotemplate {

		// use counter to find next available well position in plate

		var allwellpositionsforplate []string

		allwellpositionsforplate = _input.Plate.AllWellPositions(wtype.BYCOLUMN)

		wellposition := allwellpositionsforplate[counter]

		// Run PCR_vol element
		result := Lesson0_PCRRunSteps(_ctx, &Lesson0_PCRInput{WaterVolume: _input.DefaultWaterVolume,
			ReactionVolume:        _input.DefaultReactionVolume,
			BufferConcinX:         _input.DefaultBufferConcinX,
			FwdPrimerName:         _input.Reactiontoprimerpair[reactionname][0],
			RevPrimerName:         _input.Reactiontoprimerpair[reactionname][1],
			TemplateName:          templatename,
			ReactionName:          reactionname,
			FwdPrimerVol:          _input.DefaultPrimerVolume,
			RevPrimerVol:          _input.DefaultPrimerVolume,
			AdditiveVols:          additivevolumes,
			Templatevolume:        _input.DefaultTemplateVol,
			PolymeraseVolume:      _input.DefaultPolymeraseVolume,
			DNTPVol:               _input.DefaultDNTPVol,
			Numberofcycles:        1,
			InitDenaturationtime:  wunit.NewTime(30, "s"),
			Denaturationtime:      wunit.NewTime(5, "s"),
			Annealingtime:         wunit.NewTime(10, "s"),
			AnnealingTemp:         wunit.NewTemperature(72, "C"), // Should be calculated from primer and template binding
			Extensiontime:         wunit.NewTime(60, "s"),        // should be calculated from template length and polymerase rate
			Finalextensiontime:    wunit.NewTime(180, "s"),
			Hotstart:              false,
			AddPrimerstoMasterMix: false,
			WellPosition:          wellposition,

			FwdPrimer:     _input.FwdPrimertype,
			RevPrimer:     _input.RevPrimertype,
			DNTPS:         _input.DefaultDNTPS,
			PCRPolymerase: _input.DefaultPolymerase,
			Buffer:        _input.DefaultBuffer,
			Water:         _input.DefaultWater,
			Template:      _input.Templatetype,
			Additives:     additives,
			OutPlate:      _input.Plate},
		)

		// add result to reactions slice
		_output.Reactions = append(_output.Reactions, result.Outputs.Reaction)
		volumes = append(volumes, result.Outputs.Reaction.Volume())
		welllocations = append(welllocations, wellposition)
		_output.ReactionMap[reactionname] = result.Outputs.Reaction

		if result.Data.Status != "Success" {

			errormessage := "Reaction failure: " + reactionname

			_output.Errors = append(_output.Errors, fmt.Errorf(errormessage))

			execute.Errorf(_ctx, "Oops", errormessage)
		}

		// increase counter by 1 ready for next iteration of loop
		counter++

	}
	_, err := wtype.ExportPlateCSV(_input.Projectname+".csv", _input.Plate, _input.Projectname+"outputPlate", welllocations, _output.Reactions, volumes)
	// once all values of loop have been completed, export the plate contents as a csv file

	if err != nil {
		_output.Errors = append(_output.Errors, err)
	}
}

// Run after controls and a steps block are completed to
// post process any data and provide downstream results
func _Lesson0_PCR_MultiAnalysis(_ctx context.Context, _input *Lesson0_PCR_MultiInput, _output *Lesson0_PCR_MultiOutput) {
}

// A block of tests to perform to validate that the sample was processed correctly
// Optionally, destructive tests can be performed to validate results on a
// dipstick basis
func _Lesson0_PCR_MultiValidation(_ctx context.Context, _input *Lesson0_PCR_MultiInput, _output *Lesson0_PCR_MultiOutput) {
}
func _Lesson0_PCR_MultiRun(_ctx context.Context, input *Lesson0_PCR_MultiInput) *Lesson0_PCR_MultiOutput {
	output := &Lesson0_PCR_MultiOutput{}
	_Lesson0_PCR_MultiSetup(_ctx, input)
	_Lesson0_PCR_MultiSteps(_ctx, input, output)
	_Lesson0_PCR_MultiAnalysis(_ctx, input, output)
	_Lesson0_PCR_MultiValidation(_ctx, input, output)
	return output
}

func Lesson0_PCR_MultiRunSteps(_ctx context.Context, input *Lesson0_PCR_MultiInput) *Lesson0_PCR_MultiSOutput {
	soutput := &Lesson0_PCR_MultiSOutput{}
	output := _Lesson0_PCR_MultiRun(_ctx, input)
	if err := inject.AssignSome(output, &soutput.Data); err != nil {
		panic(err)
	}
	if err := inject.AssignSome(output, &soutput.Outputs); err != nil {
		panic(err)
	}
	return soutput
}

func Lesson0_PCR_MultiNew() interface{} {
	return &Lesson0_PCR_MultiElement{
		inject.CheckedRunner{
			RunFunc: func(_ctx context.Context, value inject.Value) (inject.Value, error) {
				input := &Lesson0_PCR_MultiInput{}
				if err := inject.Assign(value, input); err != nil {
					return nil, err
				}
				output := _Lesson0_PCR_MultiRun(_ctx, input)
				return inject.MakeValue(output), nil
			},
			In:  &Lesson0_PCR_MultiInput{},
			Out: &Lesson0_PCR_MultiOutput{},
		},
	}
}

var (
	_ = execute.MixInto
	_ = wtype.FALSE
	_ = wunit.Make_units
)

type Lesson0_PCR_MultiElement struct {
	inject.CheckedRunner
}

type Lesson0_PCR_MultiInput struct {
	AdditiveToAdditiveVolume map[string]wunit.Volume
	DefaultBuffer            *wtype.LHComponent
	DefaultBufferConcinX     int
	DefaultDNTPS             *wtype.LHComponent
	DefaultDNTPVol           wunit.Volume
	DefaultPolymerase        *wtype.LHComponent
	DefaultPolymeraseVolume  wunit.Volume
	DefaultPrimerVolume      wunit.Volume
	DefaultReactionVolume    wunit.Volume
	DefaultTemplateVol       wunit.Volume
	DefaultWater             *wtype.LHComponent
	DefaultWaterVolume       wunit.Volume
	FwdPrimertype            *wtype.LHComponent
	Plate                    *wtype.LHPlate
	Projectname              string
	Reactiontoprimerpair     map[string][2]string
	Reactiontotemplate       map[string]string
	RevPrimertype            *wtype.LHComponent
	Templatetype             *wtype.LHComponent
}

type Lesson0_PCR_MultiOutput struct {
	Errors      []error
	ReactionMap map[string]*wtype.LHComponent
	Reactions   []*wtype.LHComponent
}

type Lesson0_PCR_MultiSOutput struct {
	Data struct {
		Errors []error
	}
	Outputs struct {
		ReactionMap map[string]*wtype.LHComponent
		Reactions   []*wtype.LHComponent
	}
}

func init() {
	if err := addComponent(component.Component{Name: "Lesson0_PCR_Multi",
		Constructor: Lesson0_PCR_MultiNew,
		Desc: component.ComponentDesc{
			Desc: "Perform multiple PCR reactions with common default parameters\n",
			Path: "src/github.com/antha-lang/elements/AnthaAcademy/Lesson0_Examples/AutoPCR/AutoPCR.an",
			Params: []component.ParamDesc{
				{Name: "AdditiveToAdditiveVolume", Desc: "look up table of additives to volumes of each additive; e.g. [\"DMSO\"]:\"3ul\"\n", Kind: "Parameters"},
				{Name: "DefaultBuffer", Desc: "", Kind: "Inputs"},
				{Name: "DefaultBufferConcinX", Desc: "e.g. for  10X Q5 buffer this would be 10\n", Kind: "Parameters"},
				{Name: "DefaultDNTPS", Desc: "", Kind: "Inputs"},
				{Name: "DefaultDNTPVol", Desc: "", Kind: "Parameters"},
				{Name: "DefaultPolymerase", Desc: "", Kind: "Inputs"},
				{Name: "DefaultPolymeraseVolume", Desc: "", Kind: "Parameters"},
				{Name: "DefaultPrimerVolume", Desc: "", Kind: "Parameters"},
				{Name: "DefaultReactionVolume", Desc: "Volume for each reaction\n", Kind: "Parameters"},
				{Name: "DefaultTemplateVol", Desc: "Volume of template in each reaction\n", Kind: "Parameters"},
				{Name: "DefaultWater", Desc: "", Kind: "Inputs"},
				{Name: "DefaultWaterVolume", Desc: "", Kind: "Parameters"},
				{Name: "FwdPrimertype", Desc: "", Kind: "Inputs"},
				{Name: "Plate", Desc: "", Kind: "Inputs"},
				{Name: "Projectname", Desc: "PCRprep parameters\n", Kind: "Parameters"},
				{Name: "Reactiontoprimerpair", Desc: "map of which reaction uses which primer pair e.g. [\"left homology arm\"]:\"fwdprimer\",\"revprimer\"\n", Kind: "Parameters"},
				{Name: "Reactiontotemplate", Desc: "map of which reaction uses which template e.g. [\"left homology arm\"]:\"templatename\"\n", Kind: "Parameters"},
				{Name: "RevPrimertype", Desc: "", Kind: "Inputs"},
				{Name: "Templatetype", Desc: "", Kind: "Inputs"},
				{Name: "Errors", Desc: "return an error message if an error is encountered\n", Kind: "Data"},
				{Name: "ReactionMap", Desc: "", Kind: "Outputs"},
				{Name: "Reactions", Desc: "", Kind: "Outputs"},
			},
		},
	}); err != nil {
		panic(err)
	}
}