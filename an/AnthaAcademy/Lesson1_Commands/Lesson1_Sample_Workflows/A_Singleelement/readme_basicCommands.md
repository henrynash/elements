## antha run


Run this command from a folder containing your workflow.json file and parameters.json file
as shown here. 

This will print the output of the element to the command line by default.

________________

If a workflow or parameters set is changed you can rerun using antharun at any time. 

If you need to change the source code however, you'll need to recompile by running anthabuild.


## anthabuild

If you've added this alias this will build (recompile) all .an files in components into their corresponding .go files ready for execution. 
Whenever you change the source code of an antha element you must run anthabuild for the changes to take effect.

If you haven't set up the anthabuild alias you can do so by following the instructions [here](https://github.com/antha-lang/elements/blob/master/README.md)

### Important!
To use anthabuild to compile the antha file the an file will need to be within the antha-lang/elements/an folder. 

Otherwise the anthabuild command should be appended with AN_DIRS=<targetdirectory>.

e.g. 
```sh
anthabuild AN_DIRS=$HOME/my-antha-elements
```

## Excercises

1. Modify the A_Sample.an file you modified previously so that a additional Sample output is created called Sample2

2. Modify the steps so that Sample2 is created in the same way as Sample, i.e. with the same Solution input and SolutionVolume.


You'll need to run anthabuild since the source code is being modified. If you get an error along the way, you'll need to resolve it before being able to run the modified element.

## Next Steps

Now go to [Folder B](../B_parallelruns/readme_drivers.md) to see how to run workflows with drivers.
