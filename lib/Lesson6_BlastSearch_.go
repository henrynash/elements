// Example element demonstrating how to perform a BLAST search using the megablast algorithm
package lib

import (
	"context"
	"fmt"
	"github.com/antha-lang/antha/antha/AnthaStandardLibrary/Packages/sequences"
	biogo "github.com/antha-lang/antha/antha/AnthaStandardLibrary/Packages/sequences/biogo/ncbi/blast"
	"github.com/antha-lang/antha/antha/AnthaStandardLibrary/Packages/sequences/blast"
	"github.com/antha-lang/antha/antha/anthalib/wtype"
	"github.com/antha-lang/antha/antha/anthalib/wunit"
	"github.com/antha-lang/antha/component"
	"github.com/antha-lang/antha/execute"
	"github.com/antha-lang/antha/inject"
)

// Input parameters for this protocol

// Data which is returned from this protocol; output data

// Physical inputs to this protocol

// Physical outputs from this protocol

func _Lesson6_BlastSearchRequirements() {

}

// Actions to perform before protocol Lessonitself
func _Lesson6_BlastSearchSetup(_ctx context.Context, _input *Lesson6_BlastSearchInput) {

}

// Core process of the protocol: steps to be performed for each input
func _Lesson6_BlastSearchSteps(_ctx context.Context, _input *Lesson6_BlastSearchInput, _output *Lesson6_BlastSearchOutput) {

	var err error
	var hits []biogo.Hit
	var hitsummary string
	var identity float64
	var coverage float64
	var besthitsummary string

	_output.AnthaSeq = _input.DNA

	// look for orfs
	orf, orftrue := sequences.FindORF(_output.AnthaSeq.Seq)

	if orftrue == true && len(orf.DNASeq) == len(_output.AnthaSeq.Seq) {
		// if open reading frame is detected, we'll perform a blastP search'
		fmt.Println("ORF detected:", "full sequence length: ", len(_output.AnthaSeq.Seq), "ORF length: ", len(orf.DNASeq))
		hits, err = blast.MegaBlastP(orf.ProtSeq)
	} else {
		// otherwise we'll blast the nucleotide sequence
		hits, err = _output.AnthaSeq.Blast()
	}
	if err != nil {
		fmt.Println(err.Error())

	}

	_output.ExactHits, hitsummary, err = blast.AllExactMatches(hits)

	if len(_output.ExactHits) == 0 {
		hitsummary, err = blast.HitSummary(hits, 10, 10)
	}
	_output.BestHit, identity, coverage, besthitsummary, err = blast.FindBestHit(hits)

	//	AllHits = hits
	_output.Hitssummary = hitsummary
	fmt.Println(hitsummary)
	fmt.Println(besthitsummary)
	// Rename Sequence with ID of top blast hit

	if coverage == 100 && identity == 100 {
		_output.AnthaSeq.Nm = _output.BestHit.Id
	}
	_output.Warning = err
	_output.Identity = identity
	_output.Coverage = coverage

}

// Actions to perform after steps block to analyze data
func _Lesson6_BlastSearchAnalysis(_ctx context.Context, _input *Lesson6_BlastSearchInput, _output *Lesson6_BlastSearchOutput) {

}

func _Lesson6_BlastSearchValidation(_ctx context.Context, _input *Lesson6_BlastSearchInput, _output *Lesson6_BlastSearchOutput) {

}
func _Lesson6_BlastSearchRun(_ctx context.Context, input *Lesson6_BlastSearchInput) *Lesson6_BlastSearchOutput {
	output := &Lesson6_BlastSearchOutput{}
	_Lesson6_BlastSearchSetup(_ctx, input)
	_Lesson6_BlastSearchSteps(_ctx, input, output)
	_Lesson6_BlastSearchAnalysis(_ctx, input, output)
	_Lesson6_BlastSearchValidation(_ctx, input, output)
	return output
}

func Lesson6_BlastSearchRunSteps(_ctx context.Context, input *Lesson6_BlastSearchInput) *Lesson6_BlastSearchSOutput {
	soutput := &Lesson6_BlastSearchSOutput{}
	output := _Lesson6_BlastSearchRun(_ctx, input)
	if err := inject.AssignSome(output, &soutput.Data); err != nil {
		panic(err)
	}
	if err := inject.AssignSome(output, &soutput.Outputs); err != nil {
		panic(err)
	}
	return soutput
}

func Lesson6_BlastSearchNew() interface{} {
	return &Lesson6_BlastSearchElement{
		inject.CheckedRunner{
			RunFunc: func(_ctx context.Context, value inject.Value) (inject.Value, error) {
				input := &Lesson6_BlastSearchInput{}
				if err := inject.Assign(value, input); err != nil {
					return nil, err
				}
				output := _Lesson6_BlastSearchRun(_ctx, input)
				return inject.MakeValue(output), nil
			},
			In:  &Lesson6_BlastSearchInput{},
			Out: &Lesson6_BlastSearchOutput{},
		},
	}
}

var (
	_ = execute.MixInto
	_ = wtype.FALSE
	_ = wunit.Make_units
)

type Lesson6_BlastSearchElement struct {
	inject.CheckedRunner
}

type Lesson6_BlastSearchInput struct {
	DNA wtype.DNASequence
}

type Lesson6_BlastSearchOutput struct {
	AnthaSeq    wtype.DNASequence
	BestHit     biogo.Hit
	Coverage    float64
	ExactHits   []biogo.Hit
	Hitssummary string
	Identity    float64
	Warning     error
}

type Lesson6_BlastSearchSOutput struct {
	Data struct {
		AnthaSeq    wtype.DNASequence
		BestHit     biogo.Hit
		Coverage    float64
		ExactHits   []biogo.Hit
		Hitssummary string
		Identity    float64
		Warning     error
	}
	Outputs struct {
	}
}

func init() {
	if err := addComponent(component.Component{Name: "Lesson6_BlastSearch",
		Constructor: Lesson6_BlastSearchNew,
		Desc: component.ComponentDesc{
			Desc: "Example element demonstrating how to perform a BLAST search using the megablast algorithm\n",
			Path: "src/github.com/antha-lang/elements/AnthaAcademy/Lesson6_DNA/BlastSearch_wtype.an",
			Params: []component.ParamDesc{
				{Name: "DNA", Desc: "", Kind: "Parameters"},
				{Name: "AnthaSeq", Desc: "", Kind: "Data"},
				{Name: "BestHit", Desc: "", Kind: "Data"},
				{Name: "Coverage", Desc: "", Kind: "Data"},
				{Name: "ExactHits", Desc: "", Kind: "Data"},
				{Name: "Hitssummary", Desc: "", Kind: "Data"},
				{Name: "Identity", Desc: "", Kind: "Data"},
				{Name: "Warning", Desc: "", Kind: "Data"},
			},
		},
	}); err != nil {
		panic(err)
	}
}
