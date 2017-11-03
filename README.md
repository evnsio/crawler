# Crawler

## Overview

Simple app to crawl urls on a site and print to stdout.  

Doesn't follow exteral links, and can be depth limited.

### Example Call

```
./crawler -url http://www.evns.io/ -max-depth 1


.http://www.evns.io/
├── http://www.evns.io/
├── http://www.evns.io/about/
├── http://www.evns.io/2017/10/04/velocity-ny-microservices-at-scale.html
├── http://www.evns.io/2017/10/04/velocity-ny-helm-draft.html
├── http://www.evns.io/2017/10/04/velocity-ny-chemical-prog-secrets.html
├── http://www.evns.io/2017/02/14/ssl-setup.html
├── http://www.evns.io/2017/02/11/running-grav-in-docker.html
├── http://www.evns.io/2017/02/10/site-setup-with-terraform.html
├── http://www.evns.io/cdn-cgi/l/email-protection#f4979c869d87b491829a87da9d9b
```

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