FROM golang:1.9

WORKDIR /go/src/github.com/Bio-core/jtree
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

# Import gouuid executable to generate uuid
RUN go get github.com/nu7hatch/gouuid

RUN make build

EXPOSE 8000

# May have to add entrypoint
CMD /go/src/github.com/Bio-core/jtree/bin/jtree

# RUN ./bin/jtree -g=100