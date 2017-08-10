# shorturl
Example URL shortener written in Go

## What is this?

This is an example web service for getting short urls. This example was written for use during my introductory code talk for the Morgantown Codes 
tech meetup on August 10th 2017. 

## Getting Started

- Install Go from binaries available at [the Go website](https://golang.org/).
- Use `go get github.com/kelcecil/shorturl` or clone to your `$GOPATH`.
- Do `go run *.go`.

## Adding a URL

Just send a POST to the root with a JSON payload containing an attribute called "url":

```json
{
  "url": "http://github.com"
}
```

You can use curl to easily add a URL from command line like so.

```bash
curl -X POST http://localhost:8080/ -H "Accept: application/json" -d '{"url":"http://github.com"}'
```

The JSON payload in response will containing a JSON attribute called `key` that you'll use for your short url.

```json
{
  "key": "a"
}
```

## Redirect to a URL

Just call the root like so using the `key` you received when submitting the URL.

```bash
curl -I -X GET http://localhost:8080/a
```

Your output will look something like:

```
HTTP/1.1 307 Temporary Redirect
Location: http://github.com
Date: Thu, 10 Aug 2017 21:21:03 GMT
Content-Length: 50
Content-Type: text/html; charset=utf-8
```

Try it in a web browser for maximum effect.

## Potential Improvements

This example is intended to lead into a workshop session where people try running the application and make small improvements to the code 
to try their hand at Go programming. A few possibilities (in ascending order of difficulty) are:

- Adding more logging. ([Explore the `log` standard library package.](https://golang.org/pkg/log/))
- Add a flag to configure the port to listen on. ([Explore the `flag` standard library package.](https://golang.org/pkg/flag/))
- Add tests to test the HTTP endpoints. ([Explore the `httptest` standard library package.](https://golang.org/pkg/net/http/httptest/))
- Add benchmarks to evaluate the speed of the encoding. ([Explore the `testing` standard library package.](https://golang.org/pkg/testing/))
- Write code to ensure that the symmetric key will never produce a curse word. (Think http://localhost:8080/butts)
- Implement a new storage backend for URLStorage. ([sqlite3 might be a good option](https://github.com/mattn/go-sqlite3)).
- Implement middleware to add rate limiting or metrics collection. 

## Is this production ready?

Absolutely. 
