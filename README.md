# Crawler

## Overview

Simple app to crawl urls on a site and print to stdout.  

Doesn't follow exteral links, and can be depth limited.

## Usage 

### Build

* `make build` or
* `go build -o crawler` 

### Run

* `ARGS="-url http://www.evns.io -max-depth 2" make run` or
* `./crawler -url http://www.evns.io -max-depth 2` 

### Test

* `make test` or
* `go test -v`

## Limitations

* Currently the crawler doesn't de-duplicate pages, so the same page may be added multiple times.
* Cyclic loops aren't guarded against
* Handling of errors is poor