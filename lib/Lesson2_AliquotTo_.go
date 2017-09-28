// The lowest level example protocol Lessonshowing The MixTo command being used to specify the specific wells to be aliquoted to;
// By doing this we are able to specify whether the aliqouts are pipetted by row or by column.
// In this case the user is still not specifying the well location (i.e. A1) in the parameters, although that would be possible to specify.
// We don't generally encourage this since Antha is designed to be prodiminantly a high level language which avoids the user specifying well locations but this possibility is there if necessary.
package lib

import (
	"context"
	"github.com/antha-lang/antha/antha/anthalib/mixer"
	"github.com/antha-lang/antha/antha/anthalib/wtype"
	"github.com/antha-lang/antha/antha/anthalib/wunit"
	"github.com/antha-lang/antha/antha/anthalib/wutil"
	"github.com/antha-lang/antha/component"
	"github.com/antha-lang/antha/execute"
	"github.com/antha-lang/antha/inject"
	"strconv"
)

// Input parameters for this protocol Lesson(data)

// Data which is returned from this protocol, and data types

// Physical Inputs to this protocol Lessonwith types

// we're now going to aliquot multiple solutions at the same time (but not mixing them)

// Physical outputs from this protocol Lessonwith types

func _Lesson2_AliquotToRequirements() {

}

// Conditions to run on startup
func _Lesson2_AliquotToSetup(_ctx context.Context, _input *Lesson2_AliquotToInput) {

}

// The core process for this protocol, with the steps to be performed
// for every input
func _Lesson2_AliquotToSteps(_ctx context.Context, _input *Lesson2_AliquotToInput, _output *Lesson2_AliquotToOutput) {

	number := _input.SolutionVolume.SIValue() / _input.VolumePerAliquot.SIValue()
	possiblenumberofAliquots, _ := wutil.RoundDown(number)
	if possiblenumberofAliquots < _input.NumberofAliquots {
		execute.Errorf(_ctx, "Not enough solution for this many aliquots")
	}

	aliquots := make([]*wtype.LHComponent, 0)

	// work out well coordinates for any plate
	wellpositionarray := make([]string, 0)

	alphabet := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J",
		"K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X",
		"Y", "Z", "AA", "AB", "AC", "AD", "AE", "AF"}

	if _input.ByRow {
		// add well positions to the array based upon the number of wells per column (OutPlate.WlsX) and row (OutPlate.WlsY) of the plate type in question
		for j := 0; j < _input.OutPlate.WlsY; j++ {
			for i := 0; i < _input.OutPlate.WlsX; i++ {

				// antha, like golang upon which it is built, is a strongly type language so an int must be converted to a string using the strconv package
				// as shown here, strings can be concatenated using +
				// other types can sometimes be converted more directly.
				// In particular an int can be converted to a float64 like this:
				// var myInt int = 1
				// var myFloat float64
				// myFloat = float64(myInt)
				wellposition := alphabet[j] + strconv.Itoa(i+1)

				wellpositionarray = append(wellpositionarray, wellposition)
			}

		}
	} else {
		for j := 0; j < _input.OutPlate.WlsX; j++ {
			for i := 0; i < _input.OutPlate.WlsY; i++ {

				wellposition := alphabet[i] + strconv.Itoa(j+1)

				wellpositionarray = append(wellpositionarray, wellposition)
			}

		}
	}

	// initialise a counter
	var counter int // an int is initialised as zero therefore this is the same as counter := 0 or var counter = 0

	for _, Solution := range _input.Solutions {
		for k := 0; k < _input.NumberofAliquots; k++ {

			if Solution.TypeName() == "dna" {
				Solution.Type = wtype.LTDoNotMix
			}
			aliquotSample := mixer.Sample(Solution, _input.VolumePerAliquot)

			// this time we're using counter as an index to go through the wellpositionarray one position at a time and ensuring the next free position is chosen
			// the platenumber is hardcoded to 1 here so if we tried to specify too many aliquots in the parameters the protocol Lessonwould fail
			// it would be better to create a platenumber variable of type int and use an if statement to increase platenumber by 1 if all well positions are filled up i.e.
			// if counter == len(wellpositionarray) {
			// 		platenumber++
			//}
			aliquot := execute.MixTo(_ctx, _input.OutPlate.Type, wellpositionarray[counter], 1, aliquotSample)
			aliquots = append(aliquots, aliquot)
			counter = counter + 1 // this is the same as using the more concise counter++
		}
		_output.Aliquots = aliquots

		// Exercise: refactor to use wtype.WellCoords instead of creating the well ids manually using alphabet and strconv
	}

}

// Run after controls and a steps block are completed to
// post process any data and provide downstream results
func _Lesson2_AliquotToAnalysis(_ctx context.Context, _input *Lesson2_AliquotToInput, _output *Lesson2_AliquotToOutput) {
}

// A block of tests to perform to validate that the sample was processed correctly
// Optionally, destructive tests can be performed to validate results on a
// dipstick basis
func _Lesson2_AliquotToValidation(_ctx context.Context, _input *Lesson2_AliquotToInput, _output *Lesson2_AliquotToOutput) {

}
func _Lesson2_AliquotToRun(_ctx context.Context, input *Lesson2_AliquotToInput) *Lesson2_AliquotToOutput {
	output := &Lesson2_AliquotToOutput{}
	_Lesson2_AliquotToSetup(_ctx, input)
	_Lesson2_AliquotToSteps(_ctx, input, output)
	_Lesson2_AliquotToAnalysis(_ctx, input, output)
	_Lesson2_AliquotToValidation(_ctx, input, output)
	return output
}

func Lesson2_AliquotToRunSteps(_ctx context.Context, input *Lesson2_AliquotToInput) *Lesson2_AliquotToSOutput {
	soutput := &Lesson2_AliquotToSOutput{}
	output := _Lesson2_AliquotToRun(_ctx, input)
	if err := inject.AssignSome(output, &soutput.Data); err != nil {
		panic(err)
	}
	if err := inject.AssignSome(output, &soutput.Outputs); err != nil {
		panic(err)
	}
	return soutput
}

func Lesson2_AliquotToNew() interface{} {
	return &Lesson2_AliquotToElement{
		inject.CheckedRunner{
			RunFunc: func(_ctx context.Context, value inject.Value) (inject.Value, error) {
				input := &Lesson2_AliquotToInput{}
				if err := inject.Assign(value, input); err != nil {
					return nil, err
				}
				output := _Lesson2_AliquotToRun(_ctx, input)
				return inject.MakeValue(output), nil
			},
			In:  &Lesson2_AliquotToInput{},
			Out: &Lesson2_AliquotToOutput{},
		},
	}
}

var (
	_ = execute.MixInto
	_ = wtype.FALSE
	_ = wunit.Make_units
)

type Lesson2_AliquotToElement struct {
	inject.CheckedRunner
}

type Lesson2_AliquotToInput struct {
	ByRow            bool
	NumberofAliquots int
	OutPlate         *wtype.LHPlate
	SolutionVolume   wunit.Volume
	Solutions        []*wtype.LHComponent
	VolumePerAliquot wunit.Volume
}

type Lesson2_AliquotToOutput struct {
	Aliquots []*wtype.LHComponent
}

type Lesson2_AliquotToSOutput struct {
	Data struct {
	}
	Outputs struct {
		Aliquots []*wtype.LHComponent
	}
}

func init() {
	if err := addComponent(component.Component{Name: "Lesson2_AliquotTo",
		Constructor: Lesson2_AliquotToNew,
		Desc: component.ComponentDesc{
			Desc: "The lowest level example protocol Lessonshowing The MixTo command being used to specify the specific wells to be aliquoted to;\nBy doing this we are able to specify whether the aliqouts are pipetted by row or by column.\nIn this case the user is still not specifying the well location (i.e. A1) in the parameters, although that would be possible to specify.\nWe don't generally encourage this since Antha is designed to be prodiminantly a high level language which avoids the user specifying well locations but this possibility is there if necessary.\n",
			Path: "src/github.com/antha-lang/elements/AnthaAcademy/Lesson2_mix/E_AliquotTo_wellpositions.an",
			Params: []component.ParamDesc{
				{Name: "ByRow", Desc: "", Kind: "Parameters"},
				{Name: "NumberofAliquots", Desc: "", Kind: "Parameters"},
				{Name: "OutPlate", Desc: "", Kind: "Inputs"},
				{Name: "SolutionVolume", Desc: "", Kind: "Parameters"},
				{Name: "Solutions", Desc: "we're now going to aliquot multiple solutions at the same time (but not mixing them)\n", Kind: "Inputs"},
				{Name: "VolumePerAliquot", Desc: "", Kind: "Parameters"},
				{Name: "Aliquots", Desc: "", Kind: "Outputs"},
			},
		},
	}); err != nil {
		panic(err)
	}
}
