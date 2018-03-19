# JTree

Database Access API for AMDL Results Database

[View On SwaggerHub](https://app.swaggerhub.com/apis/JTree/jtree-metadata_api/0.1.0)
</br>
[![Build Status](https://travis-ci.org/Bio-Core/JTree.svg?branch=master)](https://travis-ci.org/Bio-Core/JTree)
[![Go Report Card](https://goreportcard.com/badge/Bio-core/Jtree)](https://goreportcard.com/report/Bio-core/Jtree)

## Getting started

This project requires Go to be installed. On OS X with Homebrew you can just run `brew install go`.

Installing this code:
`$ go get github.com/Bio-core/JTree`

To get any missing dependencies run:
`$ bash go_get.sh`

Running it then should be as simple as:

```console
$ make database
$ make build
$ ./bin/jtree
```
To generate fake data, run `$ ./bin/jtree -g=100`, where 100 is the amount of dummy data requested
To run on a spesific port run `$ ./bin/jtree -p=8000`, where 8000 is the desired port



Endpoints:

```sh
# QUERIES
# This will return all of the columns in the database
$ curl http://127.0.0.1:8000/Jtree/metadata/0.1.0/columns

# This is an example query that will return all data from every table in the database
$ curl http://127.0.0.1:8000/Jtree/metadata/0.1.0/query -X POST -H "content-type:application/json" /
-d '{"selected_fields":["*"],"selected_tables":["samples", "patients","experiments", "results", "resultdetails"],"selected_conditions":[[]]}'

# This is an example query that will return all data where the date of birth is greater than 1950
$ curl http://127.0.0.1:8000/Jtree/metadata/0.1.0/query -X POST -H "content-type:application/json" -d /
'{"selected_fields":["samples.sample_id", "patients.dob"],"selected_tables":["samples", "patients","experiments", "results", "resultdetails"],"selected_conditions":[["AND", "patients.dob", "Greater than", "1950"]]}'


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

 # results
$ curl -X POST -H "Content-Type: application/json" /
 -d '{`# See models.results for object structure`}' 127.0.0.1:8000/Jtree/metadata/0.1.0/results

 # resultdetails
$ curl -X POST -H "Content-Type: application/json" /
 -d '{`# See models.resultdetails for object structure`}' 127.0.0.1:8000/Jtree/metadata/0.1.0/resultdetails

```
