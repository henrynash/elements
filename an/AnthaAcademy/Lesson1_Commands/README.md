## Lesson 1: Key Concepts

This tutorial will teach you the basics you need to start running and writing Antha protocols. 

Here're the core concepts of how to run your first Antha programme:

### Antha elements (.an files)
Antha elements are the building blocks from which we assemble experimental workflows in Antha. 
The .an files found here show the structure of antha elements. 


```go
// Example protocol demonstrating the use of the Sample function
protocol Sample // this is the name of the protocol that will be called in a workflow or other antha element

```


```go
// we need to import the wtype package to use the LHComponent type
// the mixer package is required to use the Sample function
import (
	"github.com/antha-lang/antha/antha/anthalib/wtype"
	"github.com/antha-lang/antha/antha/anthalib/mixer"
)
```




### inputs


The Parameters and Inputs sections of these files represent the inputs to the element. 


```go
// Input parameters for this protocol (data)
Parameters (
	// antha, like golang is a strongly typed language in which the type of a variable must be declared.
	// In this case we're creating a variable called SampleVolume which is of type Volume;
	// the type system allows the antha compiler to catch many types of common errors before the programme is run
	// the antha type system extends this to biological types such as volumes here.
	// functions require inputs of particular types to be adhered to
	SampleVolume Volume 
)
```


```go
// Physical Inputs to this protocol with types
Inputs (
	// the LHComponent (i.e. Liquid Handling Component) is the principal liquidhandling type in antha 
	// the * signifies that this is a pointer to the component rather than the component itself
	// most key antha functions such as Sample and Mix use *LHComponent rather than LHComponent
	// since the type is imported from the wtype package we need to use  *wtype.LHComponent rather than simply *LHComponent
	Solution *wtype.LHComponent
)
```


### outputs


The Data and Outputs represent the outputs. 



```go

// Data which is returned from this protocol, and data types
Data (
	// Antha inherits all standard primitives valid in golang (e.g. bool, float64, string, int etc...); 
	//for example the string type shown here used to return a textual message 
	Status string
)
```


```go
// Physical outputs from this protocol with types
Outputs (
	// An output LHComponent variable is created called Sample
	Sample *wtype.LHComponent
)
```


### steps

The steps block defines how the inputs are converted into outputs. 

```go
// The core process for this protocol, with the steps to be performed
// for every input
Steps {
	
	// the Sample function is imported from the mixer library
	// in the mixer library the function signature can be found, here it is:
	// func Sample(l *wtype.LHComponent, v wunit.Volume) *wtype.LHComponent {
	// The function signature  shows that the function requires a *LHComponent and a Volume and returns an *LHComponent	
	Sample = mixer.Sample(Solution,SampleVolume)
	
	// The Sample function is not sufficient to generate liquid handling instructions alone,
	// We would need a Mix command to instruct where to put the sample
	
	// we can also create data outputs as a string like this
	// This can either be by using the ToString() method which can be used on units, such as volumes, and the .Name() method on an LHComponent  
	Status = SampleVolume.ToString() + " of " + Solution.Name() + " sampled"
	
}
```


Take a look at the three .an files in this folder and read through the comments explaining how the element is put together. 


## Excercises

1. Modify the Lesson1_Sample/element.an file by adding a step in the steps block to take the Sample produced and Mix it to an output location by adding the following line:

Sample = Mix(Sample)

2. Modify the Status message accordingly

3. Compile the changes to the source code


(a) Command line:

```bash
anthabuild AN_DIRS=$GOPATH/src/repos.antha.com/antha-ninja/elements-westeros
```

(b) AnthaOS: 

- create a new branch
- edit the file
- click commit changes button
- push and wait for the pipeline to build.

4. Now run protocol

(a) Command line:

```bash
antha run --bundle Lesson1_Sample/Lesson1_Sample.bundle.json
```

(b) AnthaOS:

Select the element set with the name of your new branch 
Select the Sample element 
Modify the parameters
click simulate

## Next Steps

Now Move to [workflows](readme_Lesson1_runningworkflows.md) to find out how to use the Antha element in a workflow with real parameters.

Or go straight to learning about writing antha elements with exercises [here](../Exercises/LiquidHandlingExercises/README.md)

