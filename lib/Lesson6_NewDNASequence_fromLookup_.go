package lib

import (
	"context"
	"fmt"
	"github.com/antha-lang/antha/antha/AnthaStandardLibrary/Packages/igem"
	"github.com/antha-lang/antha/antha/AnthaStandardLibrary/Packages/sequences"
	"github.com/antha-lang/antha/antha/AnthaStandardLibrary/Packages/sequences/entrez"
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

func _Lesson6_NewDNASequence_fromLookupRequirements() {

}

// Actions to perform before protocol Lessonitself
func _Lesson6_NewDNASequence_fromLookupSetup(_ctx context.Context, _input *Lesson6_NewDNASequence_fromLookupInput) {

}

// Core process of the protocol: steps to be performed for each input
func _Lesson6_NewDNASequence_fromLookupSteps(_ctx context.Context, _input *Lesson6_NewDNASequence_fromLookupInput, _output *Lesson6_NewDNASequence_fromLookupOutput) {

	var err error
	var partdetails igem.Rsbpml

	if _input.EntrezID {

		if _input.DNAID {

			_output.DNA, err = entrez.RetrieveSequence(_input.ID, "nucleotide")

			if err != nil {
				execute.Errorf(_ctx, err.Error())
			}

			if err != nil {
				execute.Errorf(_ctx, "Error retrieving sequence %s from entrez nucleotide database: %s", _input.ID, err.Error())
			}

		}
	} else if _input.BiobrickID {

		partdetails = igem.LookUp([]string{_input.ID})

		seq := partdetails.Sequence(_input.ID)

		_output.DNA = wtype.MakeLinearDNASequence(_input.ID, seq)

	}

	if _input.AddORFS {
		orfs := sequences.FindallORFs(_output.DNA.Seq)
		features := sequences.ORFs2Features(orfs)
		_output.DNA = wtype.Annotate(_output.DNA, features)
	}

	if _input.BiobrickID {
		_output.Status = fmt.Sprintln(
			text.Print(_input.ID+" DNA_Seq: ", _output.DNA),
			text.Print(_input.ID+" ORFs: ", _output.DNA.Features),
			text.Print(_input.ID+" PartDescription", partdetails.Description(_input.ID)),
		)
		_output.Description = partdetails.Description(_input.ID)
	} else {
		_output.Status = fmt.Sprintln(
			text.Print(_input.ID+" DNA_Seq: ", _output.DNA),
			text.Print(_input.ID+" ORFs: ", _output.DNA.Features),
		)
	}
	_output.Warnings = err
	fmt.Println(_output.Status)
}

// Actions to perform after steps block to analyze data
func _Lesson6_NewDNASequence_fromLookupAnalysis(_ctx context.Context, _input *Lesson6_NewDNASequence_fromLookupInput, _output *Lesson6_NewDNASequence_fromLookupOutput) {

}

func _Lesson6_NewDNASequence_fromLookupValidation(_ctx context.Context, _input *Lesson6_NewDNASequence_fromLookupInput, _output *Lesson6_NewDNASequence_fromLookupOutput) {

}
func _Lesson6_NewDNASequence_fromLookupRun(_ctx context.Context, input *Lesson6_NewDNASequence_fromLookupInput) *Lesson6_NewDNASequence_fromLookupOutput {
	output := &Lesson6_NewDNASequence_fromLookupOutput{}
	_Lesson6_NewDNASequence_fromLookupSetup(_ctx, input)
	_Lesson6_NewDNASequence_fromLookupSteps(_ctx, input, output)
	_Lesson6_NewDNASequence_fromLookupAnalysis(_ctx, input, output)
	_Lesson6_NewDNASequence_fromLookupValidation(_ctx, input, output)
	return output
}

func Lesson6_NewDNASequence_fromLookupRunSteps(_ctx context.Context, input *Lesson6_NewDNASequence_fromLookupInput) *Lesson6_NewDNASequence_fromLookupSOutput {
	soutput := &Lesson6_NewDNASequence_fromLookupSOutput{}
	output := _Lesson6_NewDNASequence_fromLookupRun(_ctx, input)
	if err := inject.AssignSome(output, &soutput.Data); err != nil {
		panic(err)
	}
	if err := inject.AssignSome(output, &soutput.Outputs); err != nil {
		panic(err)
	}
	return soutput
}

func Lesson6_NewDNASequence_fromLookupNew() interface{} {
	return &Lesson6_NewDNASequence_fromLookupElement{
		inject.CheckedRunner{
			RunFunc: func(_ctx context.Context, value inject.Value) (inject.Value, error) {
				input := &Lesson6_NewDNASequence_fromLookupInput{}
				if err := inject.Assign(value, input); err != nil {
					return nil, err
				}
				output := _Lesson6_NewDNASequence_fromLookupRun(_ctx, input)
				return inject.MakeValue(output), nil
			},
			In:  &Lesson6_NewDNASequence_fromLookupInput{},
			Out: &Lesson6_NewDNASequence_fromLookupOutput{},
		},
	}
}

var (
	_ = execute.MixInto
	_ = wtype.FALSE
	_ = wunit.Make_units
)

type Lesson6_NewDNASequence_fromLookupElement struct {
	inject.CheckedRunner
}

type Lesson6_NewDNASequence_fromLookupInput struct {
	AddORFS    bool
	BiobrickID bool
	DNAID      bool
	EntrezID   bool
	ID         string
}

type Lesson6_NewDNASequence_fromLookupOutput struct {
	DNA         wtype.DNASequence
	Description string
	Status      string
	Warnings    error
}

type Lesson6_NewDNASequence_fromLookupSOutput struct {
	Data struct {
		DNA         wtype.DNASequence
		Description string
		Status      string
		Warnings    error
	}
	Outputs struct {
	}
}

func init() {
	if err := addComponent(component.Component{Name: "Lesson6_NewDNASequence_fromLookup",
		Constructor: Lesson6_NewDNASequence_fromLookupNew,
		Desc: component.ComponentDesc{
			Desc: "",
			Path: "src/github.com/antha-lang/elements/AnthaAcademy/Lesson6_DNA/C_NewDNASequence_fromLookup.an",
			Params: []component.ParamDesc{
				{Name: "AddORFS", Desc: "", Kind: "Parameters"},
				{Name: "BiobrickID", Desc: "", Kind: "Parameters"},
				{Name: "DNAID", Desc: "", Kind: "Parameters"},
				{Name: "EntrezID", Desc: "", Kind: "Parameters"},
				{Name: "ID", Desc: "", Kind: "Parameters"},
				{Name: "DNA", Desc: "", Kind: "Data"},
				{Name: "Description", Desc: "", Kind: "Data"},
				{Name: "Status", Desc: "", Kind: "Data"},
				{Name: "Warnings", Desc: "", Kind: "Data"},
			},
		},
	}); err != nil {
		panic(err)
	}
}
