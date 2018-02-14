# Antha Elements

[![GoDoc](http://godoc.org/github.com/antha-lang/elements?status.svg)](http://godoc.org/github.com/antha-lang/elements)
[![Build Status](https://travis-ci.org/antha-lang/elements.svg?branch=master)](https://travis-ci.org/antha-lang/elements)

This repo is for storing and running Antha protocols. 

## Installation
Main instructions are in [antha-lang/antha](https://github.com/antha-lang/antha).

## Update submodels
```sh
git submodule update --init
```

## Build
To build or update elements in the elements folder:
```sh
make
# or
make current
```

or run this command from anywhere:
```sh
make -C "$(go list -f '{{.Dir}}' github.com/antha-lang/elements)"
```

By default, `make` will download and update any dependent libraries. If you
have any modifications to these dependencies (e.g., non-master branches), `make
current` will build elements without updating any dependent libraries.


If your elements are stored elsewhere you can change the target directory to comile with make by adding `AN_DIRS=<your-directory-here>` 
e.g...
```bash
make -C $HOME/go/src/github.com/antha-lang/elements AN_DIRS=$HOME/Documents
```

### Setting up anthabuild as an alias
The tutorial material will refer to using anthabuild as a command to recompile all antha elements.
You can set up the anthabuild alias by running this command:

#### On Mac
```sh
cat<<EOF>>$HOME/.bash_profile
alias anthabuild='make -C "$(go list -f '{{.Dir}}' github.com/antha-lang/elements)"'
EOF
source ~/.bash_profile
```

#### On Linux
```sh
cat<<EOF>>$HOME/.bashrc
alias anthabuild='make -C "$(go list -f '{{.Dir}}' github.com/antha-lang/elements)"'
EOF
source ~/.bashrc
```

Important: this will build elements stored in the elements folder, if your elements are stored elsewhere 
the anthabuild command should be appended with AN_DIRS=<targetdirectory> when running or initially setting up the anthabuild alias.

e.g. 
```sh
anthabuild AN_DIRS=$GOPATH/src/github.com/my-antha-elements
```

## Test
To run tests ok all example workflows:
```sh
make test
```

## Run 
```sh
antha run --bundle workflow-and-parameters.json
```

## Help
```sh
antha --help
```

## Academy
Go to the [Antha Academy](an/AnthaAcademy/README.md) page to be guided through how to use antha in more detail.
