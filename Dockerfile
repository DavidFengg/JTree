FROM golang:1.12

WORKDIR /go/src/github.com/Bio-core/jtree
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN make build

CMD /go/src/github.com/Bio-core/jtree/bin/jtree
