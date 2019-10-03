# Use golang:1.13 when becomes available
FROM golang:1.12

WORKDIR /go/src/github.com/Bio-core/jtree
COPY . .

# RUN go get ./...
RUN go get -d -v ./...
RUN go install -v ./...

# Import gouuid executable to generate uuid
RUN go get github.com/nu7hatch/gouuid

RUN make build

EXPOSE 8000

ENTRYPOINT [ "/go/src/github.com/Bio-core/jtree/bin/jtree" ]