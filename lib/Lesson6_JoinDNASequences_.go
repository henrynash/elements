// Demo protocol Lesson of how to create an array of dna types from parsing user inputs of various types
// scenarios handled:
// Biobrick IDS
// genbank files
// raw sequence
// inventory lookup
package lib

import (
	"strings"

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

func _Lesson6_JoinDNASequencesRequirements() {

}

// Actions to perform before protocol Lessonitself
func _Lesson6_JoinDNASequencesSetup(_ctx context.Context, _input *Lesson6_JoinDNASequencesInput) {

}

// Core process of the protocol: steps to be performed for each input
func _Lesson6_JoinDNASequencesSteps(_ctx context.Context, _input *Lesson6_JoinDNASequencesInput, _output *Lesson6_JoinDNASequencesOutput) {

	newSeq := _input.Seqsinorder[0]
	seqnames := make([]string, 0)

	for i, seq := range _input.Seqsinorder {
		if i != 0 {
			newSeq.Append(seq.Seq)
		}
		seqnames = append(seqnames, seq.Nm)
	}

	newSeq.Nm = strings.Join(seqnames, "_")

	_output.Seq = newSeq

}

// Actions to perform after steps block to analyze data
func _Lesson6_JoinDNASequencesAnalysis(_ctx context.Context, _input *Lesson6_JoinDNASequencesInput, _output *Lesson6_JoinDNASequencesOutput) {

}

func _Lesson6_JoinDNASequencesValidation(_ctx context.Context, _input *Lesson6_JoinDNASequencesInput, _output *Lesson6_JoinDNASequencesOutput) {

}
func _Lesson6_JoinDNASequencesRun(_ctx context.Context, input *Lesson6_JoinDNASequencesInput) *Lesson6_JoinDNASequencesOutput {
	output := &Lesson6_JoinDNASequencesOutput{}
	_Lesson6_JoinDNASequencesSetup(_ctx, input)
	_Lesson6_JoinDNASequencesSteps(_ctx, input, output)
	_Lesson6_JoinDNASequencesAnalysis(_ctx, input, output)
	_Lesson6_JoinDNASequencesValidation(_ctx, input, output)
	return output
}

func Lesson6_JoinDNASequencesRunSteps(_ctx context.Context, input *Lesson6_JoinDNASequencesInput) *Lesson6_JoinDNASequencesSOutput {
	soutput := &Lesson6_JoinDNASequencesSOutput{}
	output := _Lesson6_JoinDNASequencesRun(_ctx, input)
	if err := inject.AssignSome(output, &soutput.Data); err != nil {
		panic(err)
	}
	if err := inject.AssignSome(output, &soutput.Outputs); err != nil {
		panic(err)
	}
	return soutput
}

func Lesson6_JoinDNASequencesNew() interface{} {
	return &Lesson6_JoinDNASequencesElement{
		inject.CheckedRunner{
			RunFunc: func(_ctx context.Context, value inject.Value) (inject.Value, error) {
				input := &Lesson6_JoinDNASequencesInput{}
				if err := inject.Assign(value, input); err != nil {
					return nil, err
				}
				output := _Lesson6_JoinDNASequencesRun(_ctx, input)
				return inject.MakeValue(output), nil
			},
			In:  &Lesson6_JoinDNASequencesInput{},
			Out: &Lesson6_JoinDNASequencesOutput{},
		},
	}
}

var (
	_ = execute.MixInto
	_ = wtype.FALSE
	_ = wunit.Make_units
)

type Lesson6_JoinDNASequencesElement struct {
	inject.CheckedRunner
}

type Lesson6_JoinDNASequencesInput struct {
	Seqsinorder []wtype.DNASequence
}

type Lesson6_JoinDNASequencesOutput struct {
	Seq wtype.DNASequence
}

type Lesson6_JoinDNASequencesSOutput struct {
	Data struct {
		Seq wtype.DNASequence
	}
	Outputs struct {
	}
}

func init() {
	if err := addComponent(component.Component{Name: "Lesson6_JoinDNASequences",
		Constructor: Lesson6_JoinDNASequencesNew,
		Desc: component.ComponentDesc{
			Desc: "Demo protocol Lesson of how to create an array of dna types from parsing user inputs of various types\nscenarios handled:\nBiobrick IDS\ngenbank files\nraw sequence\ninventory lookup\n",
			Path: "src/github.com/antha-lang/elements/AnthaAcademy/Lesson6_DNA/E_JoinDNASequences.an",
			Params: []component.ParamDesc{
				{Name: "Seqsinorder", Desc: "", Kind: "Parameters"},
				{Name: "Seq", Desc: "", Kind: "Data"},
			},
		},
	}); err != nil {
		panic(err)
	}
}
