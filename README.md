# candig_mds

A toy metadata service

## Getting started

This project requires Go to be installed. On OS X with Homebrew you can just run `brew install go`.

Running it then should be as simple as:

```console
$ make get-deps
$ make build
$ ./bin/candig_mds
```

One can then run with

```console
$ ./bin/candig_mds --port=8000
```

and then in another terminal:

```console
$ curl http://127.0.0.1:8000/CanDIG/metadata/0.1.0/individuals/search
[]

$ curl -X POST -H 'Content-Type: application/json' \
  http://127.0.0.1:8000/CanDIG/metadata/0.1.0/individual \
  -d '{ "name" : "CANDIG_001", "description" : "foo" }'

$ curl http://127.0.0.1:8000/CanDIG/metadata/0.1.0/individuals/search
[{"attributes":null,"createdDate":"0001-01-01","description":"foo","id":"1",
  "name":"CANDIG_001","updatedDate":"0001-01-01"}]

$ curl -X POST -H 'Content-Type: application/json' \
  http://127.0.0.1:8000/CanDIG/metadata/0.1.0/biosample \
  -d '{ "individualId": "1", "name" : "CANDIG_001_sample", "description" : "foo",
        "collectionAge": "20" }'

$ curl http://127.0.0.1:8000/CanDIG/metadata/0.1.0/biosamples/search
[{"attributes":null,"collectionAge":"20","createdDate":"0001-01-01",
  "description":"foo","id":"1","individualId":"1","name":"CANDIG_001_sample",
  "updatedDate":"0001-01-01"}]
```
