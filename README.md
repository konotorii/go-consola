# Go-Consola
Colorful console outputs to make life easier.

This repo is essentially the Go version of [npm consola](https://www.npmjs.com/package/consola).


## Features
Currently there are only 6 basic functions: log, debug, info, success, error, warning.

These are called in the following:
```go
consola.log("MESSAGE", "MESSAGE2")

consola.debug("MESSAGE", "MESSAGE2")
```