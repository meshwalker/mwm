FROM golang

ADD . /go/src/github.com/meshwalker/mwm

RUN go install github.com/meshwalker/mwm

ENTRYPOINT ["mwm"]