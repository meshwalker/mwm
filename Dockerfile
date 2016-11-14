FROM golang

ADD . /go/src/meshwalker.com/mwm

RUN go install meshwalker.com/mwm

ENTRYPOINT ["mwm"]