//Some examples functions
// Calculate rate of reaction, V, of enzyme displaying Micahelis-Menten kinetics with Vmax, Km and [S] declared
// Calculating [S] and V from g/l concentration and looking up molecular weight of named substrate
// Calculating [S] and V from g/l concentration of DNA of known sequence
// Calculating [S] and V from g/l concentration of Protein product of DNA of known sequence
package lib

import (
	"context"
	"fmt"
	"github.com/antha-lang/antha/antha/AnthaStandardLibrary/Packages/Pubchem"
	"github.com/antha-lang/antha/antha/AnthaStandardLibrary/Packages/sequences"
	"github.com/antha-lang/antha/antha/anthalib/wtype"
	"github.com/antha-lang/antha/antha/anthalib/wunit"
	"github.com/antha-lang/antha/component"
	"github.com/antha-lang/antha/execute"
	"github.com/antha-lang/antha/inject"
)

// Input parameters for this protocol

//Amount
// i.e. Moles, M

//Amount

// Data which is returned from this protocol

// Physical inputs to this protocol

// Physical outputs from this protocol

func _Lesson5_DatacrunchRequirements() {

}

// Actions to perform before protocol Lessonitself
func _Lesson5_DatacrunchSetup(_ctx context.Context, _input *Lesson5_DatacrunchInput) {

}

// Core process of the protocol: steps to be performed for each input
func _Lesson5_DatacrunchSteps(_ctx context.Context, _input *Lesson5_DatacrunchInput, _output *Lesson5_DatacrunchOutput) {
	// Work out rate of reaction, V of enzyme with Michaelis-Menten kinetics and [S], Km and Vmax declared
	//Using declared values for S and unit of S
	km := wunit.NewAmount(_input.Km, _input.Kmunit) //.SIValue()
	s := wunit.NewAmount(_input.S, _input.Sunit)    //.SIValue()

	_output.V = ((s.SIValue() * _input.Vmax) / (s.SIValue() + km.SIValue()))

	// Now working out Molarity of Substrate based on conc and looking up molecular weight in pubchem

	// Look up properties
	substrate_mw, err := pubchem.MakeMolecule(_input.Substrate_name)
	if err != nil {
		execute.Errorf(_ctx, err.Error())
	}

	// calculate moles
	submoles := sequences.Moles(_input.SubstrateConc, substrate_mw.MolecularWeight, _input.SubstrateVol)
	// calculate molar concentration
	submolarconc := sequences.GtoMolarConc(_input.SubstrateConc, substrate_mw.MolecularWeight)

	// make a new amount
	s = wunit.NewAmount(submolarconc, "M")

	// use michaelis menton equation
	v_substrate_name := ((s.SIValue() * _input.Vmax) / (s.SIValue() + km.SIValue()))

	// Now working out Molarity of Substrate from DNA Sequence
	// calculate molar concentration
	dna_mw := sequences.MassDNA(_input.DNA_seq, false, false)
	dnamolarconc := sequences.GtoMolarConc(_input.DNAConc, dna_mw)

	// make a new amount
	s = wunit.NewAmount(dnamolarconc, "M")

	// use michaelis menton equation
	v_dna := ((s.SIValue() * _input.Vmax) / (s.SIValue() + km.SIValue()))

	// Now working out Molarity of Substrate from Protein product of dna Sequence

	// translate
	orf, orftrue := sequences.FindORF(_input.DNA_seq)
	var protein_mw float64
	if orftrue == true {
		protein_mw_kDA := sequences.Molecularweight(orf)
		protein_mw = protein_mw_kDA * 1000
		_output.Orftrue = orftrue
	}

	// calculate molar concentration
	proteinmolarconc := sequences.GtoMolarConc(_input.ProteinConc, protein_mw)

	// make a new amount
	s = wunit.NewAmount(submolarconc, "M")

	// use michaelis menton equation
	v_protein := ((s.SIValue() * _input.Vmax) / (s.SIValue() + km.SIValue()))

	// print report
	_output.Status = fmt.Sprintln(
		"Rate, V of enzyme at substrate conc", _input.S, _input.Sunit,
		"of enzyme with Km", km.ToString(),
		"and Vmax", _input.Vmax, _input.Vmaxunit,
		"=", _output.V, _input.Vunit, ".",
		"Substrate =", _input.Substrate_name, ". We have", _input.SubstrateVol.ToString(), "of", _input.Substrate_name, "at concentration of", _input.SubstrateConc.ToString(),
		"Therefore... Moles of", _input.Substrate_name, "=", submoles, "Moles.",
		"Molar Concentration of", _input.Substrate_name, "=", submolarconc, "Mol/L.",
		"Rate, V = ", v_substrate_name, _input.Vmaxunit,
		"Substrate =", "DNA Sequence of", _input.Gene_name, "We have", "concentration of", _input.DNAConc.ToString(),
		"Therefore... Molar conc", "=", dnamolarconc, "Mol/L",
		"Rate, V = ", v_dna, _input.Vmaxunit,
		"Substrate =", "protein from DNA sequence", _input.Gene_name, ".",
		"We have", "concentration of", _input.ProteinConc.ToString(),
		"Therefore... Molar conc", "=", proteinmolarconc, "Mol/L",
		"Rate, V = ", v_protein, _input.Vmaxunit)
}

// Actions to perform after steps block to analyze data
func _Lesson5_DatacrunchAnalysis(_ctx context.Context, _input *Lesson5_DatacrunchInput, _output *Lesson5_DatacrunchOutput) {

}

func _Lesson5_DatacrunchValidation(_ctx context.Context, _input *Lesson5_DatacrunchInput, _output *Lesson5_DatacrunchOutput) {

}
func _Lesson5_DatacrunchRun(_ctx context.Context, input *Lesson5_DatacrunchInput) *Lesson5_DatacrunchOutput {
	output := &Lesson5_DatacrunchOutput{}
	_Lesson5_DatacrunchSetup(_ctx, input)
	_Lesson5_DatacrunchSteps(_ctx, input, output)
	_Lesson5_DatacrunchAnalysis(_ctx, input, output)
	_Lesson5_DatacrunchValidation(_ctx, input, output)
	return output
}

func Lesson5_DatacrunchRunSteps(_ctx context.Context, input *Lesson5_DatacrunchInput) *Lesson5_DatacrunchSOutput {
	soutput := &Lesson5_DatacrunchSOutput{}
	output := _Lesson5_DatacrunchRun(_ctx, input)
	if err := inject.AssignSome(output, &soutput.Data); err != nil {
		panic(err)
	}
	if err := inject.AssignSome(output, &soutput.Outputs); err != nil {
		panic(err)
	}
	return soutput
}

func Lesson5_DatacrunchNew() interface{} {
	return &Lesson5_DatacrunchElement{
		inject.CheckedRunner{
			RunFunc: func(_ctx context.Context, value inject.Value) (inject.Value, error) {
				input := &Lesson5_DatacrunchInput{}
				if err := inject.Assign(value, input); err != nil {
					return nil, err
				}
				output := _Lesson5_DatacrunchRun(_ctx, input)
				return inject.MakeValue(output), nil
			},
			In:  &Lesson5_DatacrunchInput{},
			Out: &Lesson5_DatacrunchOutput{},
		},
	}
}

var (
	_ = execute.MixInto
	_ = wtype.FALSE
	_ = wunit.Make_units
)

type Lesson5_DatacrunchElement struct {
	inject.CheckedRunner
}

type Lesson5_DatacrunchInput struct {
	DNAConc        wunit.Concentration
	DNA_seq        string
	Gene_name      string
	Km             float64
	Kmunit         string
	ProteinConc    wunit.Concentration
	S              float64
	SubstrateConc  wunit.Concentration
	SubstrateVol   wunit.Volume
	Substrate_name string
	Sunit          string
	Vmax           float64
	Vmaxunit       string
	Vunit          string
}

type Lesson5_DatacrunchOutput struct {
	Orftrue bool
	Status  string
	V       float64
}

type Lesson5_DatacrunchSOutput struct {
	Data struct {
		Orftrue bool
		Status  string
		V       float64
	}
	Outputs struct {
	}
}

func init() {
	if err := addComponent(component.Component{Name: "Lesson5_Datacrunch",
		Constructor: Lesson5_DatacrunchNew,
		Desc: component.ComponentDesc{
			Desc: "Some examples functions\nCalculate rate of reaction, V, of enzyme displaying Micahelis-Menten kinetics with Vmax, Km and [S] declared\nCalculating [S] and V from g/l concentration and looking up molecular weight of named substrate\nCalculating [S] and V from g/l concentration of DNA of known sequence\nCalculating [S] and V from g/l concentration of Protein product of DNA of known sequence\n",
			Path: "src/github.com/antha-lang/elements/AnthaAcademy/Lesson5_Units2/A_Datacrunch.an",
			Params: []component.ParamDesc{
				{Name: "DNAConc", Desc: "", Kind: "Parameters"},
				{Name: "DNA_seq", Desc: "", Kind: "Parameters"},
				{Name: "Gene_name", Desc: "", Kind: "Parameters"},
				{Name: "Km", Desc: "Amount\n", Kind: "Parameters"},
				{Name: "Kmunit", Desc: "i.e. Moles, M\n", Kind: "Parameters"},
				{Name: "ProteinConc", Desc: "", Kind: "Parameters"},
				{Name: "S", Desc: "Amount\n", Kind: "Parameters"},
				{Name: "SubstrateConc", Desc: "", Kind: "Parameters"},
				{Name: "SubstrateVol", Desc: "", Kind: "Parameters"},
				{Name: "Substrate_name", Desc: "", Kind: "Parameters"},
				{Name: "Sunit", Desc: "", Kind: "Parameters"},
				{Name: "Vmax", Desc: "", Kind: "Parameters"},
				{Name: "Vmaxunit", Desc: "", Kind: "Parameters"},
				{Name: "Vunit", Desc: "", Kind: "Parameters"},
				{Name: "Orftrue", Desc: "", Kind: "Data"},
				{Name: "Status", Desc: "", Kind: "Data"},
				{Name: "V", Desc: "", Kind: "Data"},
			},
		},
	}); err != nil {
		panic(err)
	}
}
