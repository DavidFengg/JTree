# JTree Project Documentation

Find the source code on [GitHub](https://github.com/Bio-core/Jtree)
See the README on github for a overview of the endpoints available

## Building and Running the Code
In the base directory of the project folder
```sh
$ make database 
#This will tear down any existing database structure and create the new datbase structure
#Note that the database username and password may have to be changed in the stript for your system
$ make build
#This will build the go code and put it in the ./bin/ folder
$ ./bin/jtree 
#This is what runs the code
```
### Arguments for Running the Code
-g=X -> This will generate X number of patient records </br>
-s -> This will run the code with keycloak security on </br>

## Testing
All of the test cases can be run locally or will also run on travis every commit.  The travis build status can be found on the github readme document.
To run the test locally you need to build the test database strucure and then run the tests package.
```sh
$ bash ./tests/sql/DatabaseRebuild.sh
$ make test
```

## Node 38
The results db can be accessed on node 38 of mordor
The Jtree source code can be found within the _go_ directory
The front end code can be found within the _flask_ directory

## Docker


## Appendix
### Notes
 - The name of the database is JTree
 - The name of the test database is TestJTree