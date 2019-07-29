FROM golang:1.9

WORKDIR /go/src/github.com/Bio-core/jtree
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN make build

EXPOSE 8000

CMD /go/src/github.com/Bio-core/jtree/bin/jtree

# RUN ./bin/jtree -g=100