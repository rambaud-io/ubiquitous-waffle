ubiquitous-waffle

- A simple HTTP server using go


Pre-requisites:
1. must have go runtime

```
$ git clone https://github.com/rambaud-io/ubiquitous-waffle.git
$ cd ubiquitous-waffle
$ go build ./...
$ ./init
```

At the moment, the HTTP server has these endpoints:
1. GET /hacktoberfest
2. GET /healthcheck

Improvements:
[ ] format response body as json