FROM tvtamas/golang

RUN mkdir -p /go/src/github.com/meshwalker/mwm
WORKDIR /go/src/github.com/meshwalker/mwm

COPY glide.yaml
COPY glide.lock

# install dependencies with glide
RUN glide install -s -v

# copy source and build
COPY . /go/src/github.com/meshwalker/mwm
RUN go-wrapper build

# run the built binary
CMD ["go-wrapper", "run"]