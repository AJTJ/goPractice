CURRENT STATUS
check basic_types

Executable commands must always use "package main"

### updating golang

- extract tarball from http://golang.org/dl/
- will install/replace/update automatically in default dir

### STARTING A PROJECT

- create the folder with something like:
  `mkdir $GOPATH/src/github.com/user/hello`
- add a `something.go` file to it
- install it with `go install`
- run it by typing the file name in the command line

### SOME COMMANDS

- `go run FILENAME`
- while in the dir of the file, this is used to run a file.

- `go build`
- won't produce an ouptut. Instead saves the compiled package in the local build cache alongside the source code.
- It is used for confirming that it compiles/builds.

- `go clean -i`
- cleans up the "go build" cached build.

- `go install`
- returns an output file in the "bin" directory.
- automatically cleans the cache as well

- `filename`
- typing in a file name will run it directly since our bin directory is in our PATH.

- `go get`
- will fetch, build, and install a repository automatically

- `go test`
- for running the test file

### imports

```
import (
  "fmt"
  "math
)
```

- The above is a 'factored' import statement
