# gtc - golang test colors

> A simple colorizer wrapper for the `go test` command

Go's default test output is uncolored and can be hard to read. This simple wrapper adds color to the output by piping the stdout and stderr of the `go test` command and outputting ansii colorized text.

![screenshot](screenshot.png)

This should work on any platform that supports ansii color codes, including Windows.

gtc uses your system's `go` command to run tests, so it should work with any version of Go.

## Installation

```bash
go install github.com/danawoodman/gtc
```

## Usage

Just replace `go test` with `gtc` in your test commands. All flags will be passed as-is to `go test`.

```bash
# single test run
gtc ./...

# with flags
gtc -v ./...
```

## License

[MIT](LICENSE)
