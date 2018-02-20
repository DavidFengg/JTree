# JTree

Database Access API

[View On SwaggerHub](https://app.swaggerhub.com/apis/JTree/jtree-metadata_api/0.1.0)
</br>
[![Build Status](https://travis-ci.org/Bio-Core/JTree.svg?branch=master)](https://travis-ci.org/Bio-Core/JTree)
[![Go Report Card](https://goreportcard.com/badge/Bio-core/Jtree)](https://goreportcard.com/report/Bio-core/Jtree)

## Getting started

This project requires Go to be installed. On OS X with Homebrew you can just run `brew install go`.

Running it then should be as simple as:

```console
$ make get-deps
$ make database
$ make build
$ ./bin/jtree
```
To generate fake data, run `$ ./bin/jtree -g=100` instead, where 100 is the amount of dummy data requested


Endpoints:

```sh
# This will return all of the columns in the database
$ curl http://127.0.0.1:8000/Jtree/metadata/0.1.0/columns
[]

# This is an example query that will return all data from the samples and patients tables
$ curl -X POST -H "Content-Type: application/json" /
 -d '{"selected_tables":["samples", "patients"], "selected_fields":["*"], "selected_conditions":[]}' /
 127.0.0.1:8000/Jtree/metadata/0.1.0/query

# INSERTS
# samples
$ curl -X POST -H "Content-Type: application/json" /
 -d '{`# See models.samples for object structure`}' 127.0.0.1:8000/Jtree/metadata/0.1.0/samples

# patients
$ curl -X POST -H "Content-Type: application/json" /
 -d '{`# See models.patients for object structure`}' 127.0.0.1:8000/Jtree/metadata/0.1.0/patients

# experiments
$ curl -X POST -H "Content-Type: application/json" /
 -d '{`# See models.experiments for object structure`}' 127.0.0.1:8000/Jtree/metadata/0.1.0/experiments


```
