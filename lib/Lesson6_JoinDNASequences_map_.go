package lib

import (
	"context"
	"fmt"
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

func _Lesson6_JoinDNASequences_mapRequirements() {

}

// Actions to perform before protocol Lessonitself
func _Lesson6_JoinDNASequences_mapSetup(_ctx context.Context, _input *Lesson6_JoinDNASequences_mapInput) {

}

// Core process of the protocol: steps to be performed for each input
func _Lesson6_JoinDNASequences_mapSteps(_ctx context.Context, _input *Lesson6_JoinDNASequences_mapInput, _output *Lesson6_JoinDNASequences_mapOutput) {

	// make a lookup table of []DNASequences; in go these are called maps
	// in this case the map is called sequencepairs; the key is of type string; the values returned will be an array of DNASequences
	// e.g. we can add two sequences eGFP and sfGFP to the map like so:
	// sequencepairs["GFPSequences"] = DNASequence{eGFP,sfGFP}
	// we can return the two stored sequences to a variable called gfpseqscopy using the key "GFPSequences" by typing
	// gfpseqscopy := sequencepairs["GFPSequences"]

	sequencepairs := make(map[string][]wtype.DNASequence)

	// Step 1. Call antha element for turning a string array of parts into []DNASequence for each map entry
	// We can range through maps in a similar fashion to arrays; the important difference is after for the key value is used rather than the index as is the case with array
	for key, values := range _input.Pairs {

		// this is how we call an antha element from within an element
		seqs := Lesson6_NewDNASequencesRunSteps(_ctx, &Lesson6_NewDNASequencesInput{Seqsinorder: values,
			BlastSeqswithNoName: _input.BlastSeqswithNoName,
			Vectors:             _input.Vectors},
		)

		sequencepairs[key] = seqs.Data.Parts
	}

	//  make an array of seqs to export for each map combination
	seqstoexport := make([]wtype.DNASequence, 0)

	// Step 2. Range through the map created in step 1.
	for key, Seqsinorder := range sequencepairs {

		newSeq := Seqsinorder[0]
		fmt.Println("seq?", Seqsinorder[0].Nm)
		//seqnames := make([]string,0)

		// Step 2a. Each set of sequences we'll range through and concatenate the sequence with the next sequence
		for i, seq := range Seqsinorder {
			fmt.Println("seq[i]?", Seqsinorder[i].Nm)
			if i != 0 {
				newSeq.Append(seq.Seq)
			}
			//seqnames = append(seqnames,seq.Nm)
		}

		// Step 2b. Name the new DNAParts using map key. This could also name by concatenating but we'll use key for now
		newSeq.Nm = key //strings.Join(seqnames,"_")
		seqstoexport = append(seqstoexport, newSeq)
	}

	_output.JoinedSeqs = seqstoexport

}

// Actions to perform after steps block to analyze data
func _Lesson6_JoinDNASequences_mapAnalysis(_ctx context.Context, _input *Lesson6_JoinDNASequences_mapInput, _output *Lesson6_JoinDNASequences_mapOutput) {

}

func _Lesson6_JoinDNASequences_mapValidation(_ctx context.Context, _input *Lesson6_JoinDNASequences_mapInput, _output *Lesson6_JoinDNASequences_mapOutput) {

}
func _Lesson6_JoinDNASequences_mapRun(_ctx context.Context, input *Lesson6_JoinDNASequences_mapInput) *Lesson6_JoinDNASequences_mapOutput {
	output := &Lesson6_JoinDNASequences_mapOutput{}
	_Lesson6_JoinDNASequences_mapSetup(_ctx, input)
	_Lesson6_JoinDNASequences_mapSteps(_ctx, input, output)
	_Lesson6_JoinDNASequences_mapAnalysis(_ctx, input, output)
	_Lesson6_JoinDNASequences_mapValidation(_ctx, input, output)
	return output
}

func Lesson6_JoinDNASequences_mapRunSteps(_ctx context.Context, input *Lesson6_JoinDNASequences_mapInput) *Lesson6_JoinDNASequences_mapSOutput {
	soutput := &Lesson6_JoinDNASequences_mapSOutput{}
	output := _Lesson6_JoinDNASequences_mapRun(_ctx, input)
	if err := inject.AssignSome(output, &soutput.Data); err != nil {
		panic(err)
	}
	if err := inject.AssignSome(output, &soutput.Outputs); err != nil {
		panic(err)
	}
	return soutput
}

func Lesson6_JoinDNASequences_mapNew() interface{} {
	return &Lesson6_JoinDNASequences_mapElement{
		inject.CheckedRunner{
			RunFunc: func(_ctx context.Context, value inject.Value) (inject.Value, error) {
				input := &Lesson6_JoinDNASequences_mapInput{}
				if err := inject.Assign(value, input); err != nil {
					return nil, err
				}
				output := _Lesson6_JoinDNASequences_mapRun(_ctx, input)
				return inject.MakeValue(output), nil
			},
			In:  &Lesson6_JoinDNASequences_mapInput{},
			Out: &Lesson6_JoinDNASequences_mapOutput{},
		},
	}
}

var (
	_ = execute.MixInto
	_ = wtype.FALSE
	_ = wunit.Make_units
)

type Lesson6_JoinDNASequences_mapElement struct {
	inject.CheckedRunner
}

type Lesson6_JoinDNASequences_mapInput struct {
	BlastSeqswithNoName bool
	Pairs               map[string][]string
	Vectors             bool
}

type Lesson6_JoinDNASequences_mapOutput struct {
	JoinedSeqs []wtype.DNASequence
}

type Lesson6_JoinDNASequences_mapSOutput struct {
	Data struct {
		JoinedSeqs []wtype.DNASequence
	}
	Outputs struct {
	}
}

func init() {
	if err := addComponent(component.Component{Name: "Lesson6_JoinDNASequences_map",
		Constructor: Lesson6_JoinDNASequences_mapNew,
		Desc: component.ComponentDesc{
			Desc: "",
			Path: "src/github.com/antha-lang/elements/AnthaAcademy/Lesson6_DNA/G_JoinSequencePairs.an",
			Params: []component.ParamDesc{
				{Name: "BlastSeqswithNoName", Desc: "", Kind: "Parameters"},
				{Name: "Pairs", Desc: "", Kind: "Parameters"},
				{Name: "Vectors", Desc: "", Kind: "Parameters"},
				{Name: "JoinedSeqs", Desc: "", Kind: "Data"},
			},
		},
	}); err != nil {
		panic(err)
	}
}