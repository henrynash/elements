// protocol Lessonfor creating a DNASequence from a sequence file format. // Supported formats: .gdx .fasta .gb
package lib

import (
	"context"
	"fmt"
	"github.com/antha-lang/antha/antha/AnthaStandardLibrary/Packages/sequences"
	"github.com/antha-lang/antha/antha/AnthaStandardLibrary/Packages/sequences/parse"
	"github.com/antha-lang/antha/antha/AnthaStandardLibrary/Packages/text"
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

func _Lesson6_NewDNASequence_fromfileRequirements() {

}

// Actions to perform before protocol Lessonitself
func _Lesson6_NewDNASequence_fromfileSetup(_ctx context.Context, _input *Lesson6_NewDNASequence_fromfileInput) {

}

// Core process of the protocol: steps to be performed for each input
func _Lesson6_NewDNASequence_fromfileSteps(_ctx context.Context, _input *Lesson6_NewDNASequence_fromfileInput, _output *Lesson6_NewDNASequence_fromfileOutput) {

	seqs, err := parse.DNAFileToDNASequence(_input.SequenceFile)
	if err != nil {
		execute.Errorf(_ctx, err.Error())
	}
	if len(seqs) == 1 && err == nil {

		_output.DNA = seqs[0]

		_output.DNA.Nm = _input.Gene_name
		_output.DNA.Plasmid = _input.Plasmid

	}

	orfs := sequences.FindallORFs(_output.DNA.Seq)

	if len(_output.DNA.Features) == 0 {
		features := sequences.ORFs2Features(orfs)

		_output.DNA = wtype.Annotate(_output.DNA, features)
	}

	_output.Status = fmt.Sprintln(
		text.Print("DNA_Seq: ", _output.DNA),
		text.Print("ORFs: ", orfs),
	)

	_output.Warnings = err

}

// Actions to perform after steps block to analyze data
func _Lesson6_NewDNASequence_fromfileAnalysis(_ctx context.Context, _input *Lesson6_NewDNASequence_fromfileInput, _output *Lesson6_NewDNASequence_fromfileOutput) {

}

func _Lesson6_NewDNASequence_fromfileValidation(_ctx context.Context, _input *Lesson6_NewDNASequence_fromfileInput, _output *Lesson6_NewDNASequence_fromfileOutput) {

}
func _Lesson6_NewDNASequence_fromfileRun(_ctx context.Context, input *Lesson6_NewDNASequence_fromfileInput) *Lesson6_NewDNASequence_fromfileOutput {
	output := &Lesson6_NewDNASequence_fromfileOutput{}
	_Lesson6_NewDNASequence_fromfileSetup(_ctx, input)
	_Lesson6_NewDNASequence_fromfileSteps(_ctx, input, output)
	_Lesson6_NewDNASequence_fromfileAnalysis(_ctx, input, output)
	_Lesson6_NewDNASequence_fromfileValidation(_ctx, input, output)
	return output
}

func Lesson6_NewDNASequence_fromfileRunSteps(_ctx context.Context, input *Lesson6_NewDNASequence_fromfileInput) *Lesson6_NewDNASequence_fromfileSOutput {
	soutput := &Lesson6_NewDNASequence_fromfileSOutput{}
	output := _Lesson6_NewDNASequence_fromfileRun(_ctx, input)
	if err := inject.AssignSome(output, &soutput.Data); err != nil {
		panic(err)
	}
	if err := inject.AssignSome(output, &soutput.Outputs); err != nil {
		panic(err)
	}
	return soutput
}

func Lesson6_NewDNASequence_fromfileNew() interface{} {
	return &Lesson6_NewDNASequence_fromfileElement{
		inject.CheckedRunner{
			RunFunc: func(_ctx context.Context, value inject.Value) (inject.Value, error) {
				input := &Lesson6_NewDNASequence_fromfileInput{}
				if err := inject.Assign(value, input); err != nil {
					return nil, err
				}
				output := _Lesson6_NewDNASequence_fromfileRun(_ctx, input)
				return inject.MakeValue(output), nil
			},
			In:  &Lesson6_NewDNASequence_fromfileInput{},
			Out: &Lesson6_NewDNASequence_fromfileOutput{},
		},
	}
}

var (
	_ = execute.MixInto
	_ = wtype.FALSE
	_ = wunit.Make_units
)

type Lesson6_NewDNASequence_fromfileElement struct {
	inject.CheckedRunner
}

type Lesson6_NewDNASequence_fromfileInput struct {
	Gene_name      string
	Linear         bool
	Plasmid        bool
	SequenceFile   wtype.File
	SingleStranded bool
}

type Lesson6_NewDNASequence_fromfileOutput struct {
	DNA      wtype.DNASequence
	Status   string
	Warnings error
}

type Lesson6_NewDNASequence_fromfileSOutput struct {
	Data struct {
		DNA      wtype.DNASequence
		Status   string
		Warnings error
	}
	Outputs struct {
	}
}

func init() {
	if err := addComponent(component.Component{Name: "Lesson6_NewDNASequence_fromfile",
		Constructor: Lesson6_NewDNASequence_fromfileNew,
		Desc: component.ComponentDesc{
			Desc: "protocol Lessonfor creating a DNASequence from a sequence file format. // Supported formats: .gdx .fasta .gb\n",
			Path: "src/github.com/antha-lang/elements/AnthaAcademy/Lesson6_DNA/B_NewDNASequence_fromfile.an",
			Params: []component.ParamDesc{
				{Name: "Gene_name", Desc: "", Kind: "Parameters"},
				{Name: "Linear", Desc: "", Kind: "Parameters"},
				{Name: "Plasmid", Desc: "", Kind: "Parameters"},
				{Name: "SequenceFile", Desc: "", Kind: "Parameters"},
				{Name: "SingleStranded", Desc: "", Kind: "Parameters"},
				{Name: "DNA", Desc: "", Kind: "Data"},
				{Name: "Status", Desc: "", Kind: "Data"},
				{Name: "Warnings", Desc: "", Kind: "Data"},
			},
		},
	}); err != nil {
		panic(err)
	}
}
