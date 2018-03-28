FROM golang:1.8

WORKDIR /go/src/github.com/Bio-core/Jtree
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN make build

CMD /go/src/github.com/Bio-core/Jtree/bin/jtree