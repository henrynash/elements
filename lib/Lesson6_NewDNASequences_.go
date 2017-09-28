// Demo protocol Lessonof how to create an array of dna types from parsing user inputs of various types
// scenarios handled:
// Biobrick IDS
// raw sequence
// inventory lookup
package lib

import (
	"context"
	inventory "github.com/antha-lang/antha/antha/AnthaStandardLibrary/Packages/Inventory"
	"github.com/antha-lang/antha/antha/AnthaStandardLibrary/Packages/igem"
	"github.com/antha-lang/antha/antha/AnthaStandardLibrary/Packages/sequences"
	"github.com/antha-lang/antha/antha/anthalib/wtype"
	"github.com/antha-lang/antha/antha/anthalib/wunit"
	"github.com/antha-lang/antha/component"
	"github.com/antha-lang/antha/execute"
	"github.com/antha-lang/antha/inject"
	"strconv"
	"strings"
)

// Input parameters for this protocol

// Data which is returned from this protocol

// Physical inputs to this protocol

// Physical outputs from this protocol

func _Lesson6_NewDNASequencesRequirements() {

}

// Actions to perform before protocol Lessonitself
func _Lesson6_NewDNASequencesSetup(_ctx context.Context, _input *Lesson6_NewDNASequencesInput) {

}

// Core process of the protocol: steps to be performed for each input
func _Lesson6_NewDNASequencesSteps(_ctx context.Context, _input *Lesson6_NewDNASequencesInput, _output *Lesson6_NewDNASequencesOutput) {

	// set warnings reported back to user to none initially
	warnings := make([]string, 0)

	// make an empty array of DNA Sequences ready to fill
	partsinorder := make([]wtype.DNASequence, 0)

	var partDNA wtype.DNASequence

	_output.Status = "all parts available"
	for i, part := range _input.Seqsinorder {
		// check if biobrick part
		if strings.Contains(part, "BBa_") {
			nm := _input.PartPrefix + "_" + "Part " + strconv.Itoa(i) + "_" + part
			part = igem.GetSequence(part)

			if _input.Vectors {
				partDNA = wtype.MakePlasmidDNASequence(nm, part)
			} else {
				partDNA = wtype.MakeLinearDNASequence(nm, part)
			}
			// check if in inventory
		} else if inventoryDNA, found := inventory.Partslist()[part]; found {
			partDNA = inventoryDNA

			// else treat as DNA sequence and check
		} else {

			if _input.Vectors {
				partDNA = wtype.MakePlasmidDNASequence(_input.PartPrefix+"_"+"Part "+strconv.Itoa(i), part)
			} else {
				partDNA = wtype.MakeLinearDNASequence(_input.PartPrefix+"_"+"Part "+strconv.Itoa(i), part)
			}

			// test for illegal nucleotides
			pass, illegals, _ := sequences.Illegalnucleotides(partDNA)

			if !pass {
				var newstatus = make([]string, 0)
				for _, illegal := range illegals {

					newstatus = append(newstatus, _input.PartPrefix+"_"+"part: "+partDNA.Nm+" "+partDNA.Seq+": contains illegalnucleotides:"+illegal.ToString())
				}
				warnings = append(warnings, strings.Join(newstatus, ""))
				execute.Errorf(_ctx, strings.Join(newstatus, ""))
			}
			if pass && _input.BlastSeqswithNoName {
				// run a blast search on the sequence to get the name
				blastsearch := Lesson6_BlastSearchRunSteps(_ctx, &Lesson6_BlastSearchInput{DNA: partDNA})
				partDNA.Nm = blastsearch.Data.AnthaSeq.Nm
			}

		}
		partsinorder = append(partsinorder, partDNA)
	}

	_output.Parts = partsinorder

	_output.Warnings = warnings

}

// Actions to perform after steps block to analyze data
func _Lesson6_NewDNASequencesAnalysis(_ctx context.Context, _input *Lesson6_NewDNASequencesInput, _output *Lesson6_NewDNASequencesOutput) {

}

func _Lesson6_NewDNASequencesValidation(_ctx context.Context, _input *Lesson6_NewDNASequencesInput, _output *Lesson6_NewDNASequencesOutput) {

}
func _Lesson6_NewDNASequencesRun(_ctx context.Context, input *Lesson6_NewDNASequencesInput) *Lesson6_NewDNASequencesOutput {
	output := &Lesson6_NewDNASequencesOutput{}
	_Lesson6_NewDNASequencesSetup(_ctx, input)
	_Lesson6_NewDNASequencesSteps(_ctx, input, output)
	_Lesson6_NewDNASequencesAnalysis(_ctx, input, output)
	_Lesson6_NewDNASequencesValidation(_ctx, input, output)
	return output
}

func Lesson6_NewDNASequencesRunSteps(_ctx context.Context, input *Lesson6_NewDNASequencesInput) *Lesson6_NewDNASequencesSOutput {
	soutput := &Lesson6_NewDNASequencesSOutput{}
	output := _Lesson6_NewDNASequencesRun(_ctx, input)
	if err := inject.AssignSome(output, &soutput.Data); err != nil {
		panic(err)
	}
	if err := inject.AssignSome(output, &soutput.Outputs); err != nil {
		panic(err)
	}
	return soutput
}

func Lesson6_NewDNASequencesNew() interface{} {
	return &Lesson6_NewDNASequencesElement{
		inject.CheckedRunner{
			RunFunc: func(_ctx context.Context, value inject.Value) (inject.Value, error) {
				input := &Lesson6_NewDNASequencesInput{}
				if err := inject.Assign(value, input); err != nil {
					return nil, err
				}
				output := _Lesson6_NewDNASequencesRun(_ctx, input)
				return inject.MakeValue(output), nil
			},
			In:  &Lesson6_NewDNASequencesInput{},
			Out: &Lesson6_NewDNASequencesOutput{},
		},
	}
}

var (
	_ = execute.MixInto
	_ = wtype.FALSE
	_ = wunit.Make_units
)

type Lesson6_NewDNASequencesElement struct {
	inject.CheckedRunner
}

type Lesson6_NewDNASequencesInput struct {
	BlastSeqswithNoName bool
	PartPrefix          string
	Seqsinorder         []string
	Vectors             bool
}

type Lesson6_NewDNASequencesOutput struct {
	Parts    []wtype.DNASequence
	Status   string
	Warnings []string
}

type Lesson6_NewDNASequencesSOutput struct {
	Data struct {
		Parts    []wtype.DNASequence
		Status   string
		Warnings []string
	}
	Outputs struct {
	}
}

func init() {
	if err := addComponent(component.Component{Name: "Lesson6_NewDNASequences",
		Constructor: Lesson6_NewDNASequencesNew,
		Desc: component.ComponentDesc{
			Desc: "Demo protocol Lessonof how to create an array of dna types from parsing user inputs of various types\nscenarios handled:\nBiobrick IDS\nraw sequence\ninventory lookup\n",
			Path: "src/github.com/antha-lang/elements/AnthaAcademy/Lesson6_DNA/D_NewDNASequences.an",
			Params: []component.ParamDesc{
				{Name: "BlastSeqswithNoName", Desc: "", Kind: "Parameters"},
				{Name: "PartPrefix", Desc: "", Kind: "Parameters"},
				{Name: "Seqsinorder", Desc: "", Kind: "Parameters"},
				{Name: "Vectors", Desc: "", Kind: "Parameters"},
				{Name: "Parts", Desc: "", Kind: "Data"},
				{Name: "Status", Desc: "", Kind: "Data"},
				{Name: "Warnings", Desc: "", Kind: "Data"},
			},
		},
	}); err != nil {
		panic(err)
	}
}