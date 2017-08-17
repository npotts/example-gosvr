# example-gosvr

Dead Simple Go HTTP Server example with basic handlers for web app tinkering and learning exercises

# Installing & Usage

Since this is a ```go``` app, you need to have go installed.  On debian systems, simply do a ```apt install golang```.  rpm based systems are similar.  Once that is complete:

```sh
  # If you haven't installed any Go-based tooling, you need to have a GOPATH
  # set which is the directory Go will install local binaries and sources
   export GOPATH="${HOME}/go"

   #go get fetches and installs the dependancies needed by example-gosvr
   go get -u http://github.com/npotts/example-gosvr

   #go install builds packages, and installs the object code in $GOPATH/pkg
   # if it a package, or $GOPATH/bin if it is an executable binary.
   go install github.com/npotts/example-gosvr

   # CD to wherever your HTML sources are
   cd /where/ever/your/html/stuff/is

   # as part of the demo, the example-gosvr hosts the contents of
   # /css and /js from [cwd]/css and [cwd]/js
   mkdir css js

   # Show usage help
   $GOPATH/bin/example-gosvr --help

   # Starts the binary.  It doesn't daemonize
   $GOPATH/bin/example-gosvr

   #point your browser to http://localhost:8080
   open http://localhost:8080

   # There are "special" routes, http://localhost:8080/time that
   # presents a JSON object of the local time and time since daemon was started
   # http://localhost:8080/random returns a JSON object with a random integer
   # and a float in the range of [0.0, 1.0)

```
