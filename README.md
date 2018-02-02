# JTree

**NOTE: Fork of jtree project for now...will have to slowly move things over**
[![Build Status](https://travis-ci.org/jtree/jtree.svg?branch=master)](https://travis-ci.org/jtree/jtree)

A toy metadata service

[View On SwaggerHub](https://app.swaggerhub.com/apis/jtree/jtree-metadata_api/0.1.0)

## Getting started

This project requires Go to be installed. On OS X with Homebrew you can just run `brew install go`.

Running it then should be as simple as:

```console
$ make get-deps
$ make build
$ ./bin/jtree --port=8000
```

Alternately, one can run

```console
$ docker pull quay.io/jtree/jtree
```
and replace 8000 below with the port output.

Query the server in another terminal:

```console
$ curl http://127.0.0.1:8000/jtree/metadata/0.1.0/individuals/search
[]

$ curl -X POST -H 'Content-Type: application/json' \
  http://127.0.0.1:8000/jtree/metadata/0.1.0/individual \
  -d '{ "name" : "jtree_001", "description" : "foo" }'

$ curl http://127.0.0.1:8000/jtree/metadata/0.1.0/individuals/search
[{"attributes":null,"createdDate":"0001-01-01","description":"foo","id":"1",
  "name":"jtree_001","updatedDate":"0001-01-01"}]

$ curl -X POST -H 'Content-Type: application/json' \
  http://127.0.0.1:8000/jtree/metadata/0.1.0/biosample \
  -d '{ "individualId": "1", "name" : "jtree_001_sample", "description" : "foo", \
        "collectionAge": "20" }'

$ curl http://127.0.0.1:8000/jtree/metadata/0.1.0/biosamples/search
[{"attributes":null,"collectionAge":"20","createdDate":"0001-01-01",
  "description":"foo","id":"1","individualId":"1","name":"jtree_001_sample",
  "updatedDate":"0001-01-01"}]
```
