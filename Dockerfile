FROM golang:1.9

RUN apt-get update -qq && apt-get install -y libglpk-dev protobuf-compiler

RUN go get github.com/antha-lang/antha/...
ADD . /go/src/github.com/antha-lang/elements
RUN make -C /go/src/github.com/antha-lang/elements
