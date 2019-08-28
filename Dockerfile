# Use golang:1.13 when becomes available
FROM golang:1.12

WORKDIR /go/src/github.com/Bio-core/jtree
COPY . .

# This is getting master branch
RUN go get -d -v ./...
RUN go get github.com/Bio-core/keycloakgo
RUN go get github.com/go-openapi/errors
RUN go get github.com/go-openapi/loads
RUN go get github.com/go-openapi/runtime
RUN go get github.com/go-openapi/runtime/flagext
RUN go get github.com/go-openapi/runtime/middleware
RUN go get github.com/go-openapi/runtime/security
RUN go get github.com/go-openapi/spec
RUN go get github.com/go-openapi/strfmt
RUN go get github.com/go-openapi/swag
RUN go get github.com/go-sql-driver/mysql

# RUN git checkout add_docker_compose
# RUN git checkout add_docker_compose
RUN go install -v ./...
RUN git checkout add_docker_compose

# Import gouuid executable to generate uuid
RUN go get github.com/nu7hatch/gouuid

RUN make build

EXPOSE 8000

ENTRYPOINT [ "/go/src/github.com/Bio-core/jtree/bin/jtree" ]