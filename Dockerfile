# Build Stage
FROM lacion/docker-alpine:gobuildimage AS build-stage

LABEL app="build-candig_mds"
LABEL REPO="https://github.com/ljdursi/candig_mds"

ENV GOROOT=/usr/lib/go \
    GOPATH=/gopath \
    GOBIN=/gopath/bin \
    PROJPATH=/gopath/src/github.com/ljdursi/candig_mds

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:$GOROOT/bin:$GOPATH/bin

ADD . /gopath/src/github.com/ljdursi/candig_mds
WORKDIR /gopath/src/github.com/ljdursi/candig_mds

RUN apk add glide
RUN make get-deps
RUN make build-alpine

# Final Stage
FROM lacion/docker-alpine:latest

ARG GIT_COMMIT
ARG VERSION
LABEL REPO="https://github.com/ljdursi/candig_mds"
LABEL GIT_COMMIT=$GIT_COMMIT
LABEL VERSION=$VERSION

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:/opt/candig_mds/bin

WORKDIR /opt/candig_mds/bin

COPY --from=build-stage /gopath/src/github.com/ljdursi/candig_mds/bin/candig_mds /opt/candig_mds/bin/
RUN chmod +x /opt/candig_mds/bin/candig_mds

CMD /opt/candig_mds/bin/candig_mds
