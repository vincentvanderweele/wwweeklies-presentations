FROM golang:1

ENV RUN_DIR /go/src/github.com/vincentvanderweele/wwweeklies-presentations

RUN go get -v github.com/pilu/fresh

RUN mkdir -p $RUN_DIR
WORKDIR $RUN_DIR

COPY . $RUN_DIR
RUN go-wrapper install

CMD go-wrapper run
