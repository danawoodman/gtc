package internal_test

import (
	"fmt"
	"testing"

	"github.com/danawoodman/gtc/internal"
	"github.com/fatih/color"
	"github.com/stretchr/testify/assert"
)

func TestPackage(t *testing.T) {
	// Required to turn on colorized output.
	color.NoColor = false

	tests := []struct {
		name, input, expected string
	}{
		//------------------------------------------------
		// Success output
		//------------------------------------------------
		{
			"matches passed tests",
			"--- PASS: TestFoo (0.00s)",
			color.GreenString("--- PASS: TestFoo (0.00s)"),
		},
		{
			"matches passed sub-tests",
			"    --- PASS: TestFoo (0.00s)",
			color.GreenString("    --- PASS: TestFoo (0.00s)"),
		},
		{
			"matches error diff expected value",
			"    expected: true",
			color.GreenString("    expected: true"),
		},
		{
			"matches ok tests",
			"ok      github.com/foo/bar/baz 0.123s",
			color.GreenString("ok      github.com/foo/bar/baz 0.123s"),
		},

		//------------------------------------------------
		// Warning output
		//------------------------------------------------
		{
			"matches skipped tests",
			"--- SKIP: TestFoo (0.00s)",
			color.YellowString("--- SKIP: TestFoo (0.00s)"),
		},
		{
			"matches skipped sub-tests",
			"    --- SKIP: TestFoo (0.00s)",
			color.YellowString("    --- SKIP: TestFoo (0.00s)"),
		},

		//------------------------------------------------
		// Muted output
		//------------------------------------------------

		{
			"matches verbose run output",
			"=== RUN    TestFoo/bar",
			color.BlackString("=== RUN    TestFoo/bar"),
		},
		{
			"matches verbose cont output",
			"=== CONT    TestFoo/bar",
			color.BlackString("=== CONT    TestFoo/bar"),
		},
		{
			"matches verbose pause output",
			"=== PAUSE    TestFoo/bar",
			color.BlackString("=== PAUSE    TestFoo/bar"),
		},
		{
			"matches no test files message",
			"?      github.com/foo/bar/baz   [no test files]",
			color.BlackString("?      github.com/foo/bar/baz   [no test files]"),
		},

		//------------------------------------------------
		// Error output
		//------------------------------------------------

		{
			"matches failed tests result",
			"FAIL",
			color.RedString("FAIL"),
		},
		{
			"matches failed test filename output",
			"FAIL    github.com/foo/bar/baz    0.893s",
			color.RedString("FAIL    github.com/foo/bar/baz    0.893s"),
		},
		{
			"matches error type message",
			"    Error:    Not Equal:",
			color.RedString("    Error:    Not Equal:"),
		},
		{
			"matches error trace type message",
			"    Error Trace:    /foo/bar/baz.go:17",
			color.RedString("    Error Trace:    /foo/bar/baz.go:17"),
		},
		{
			"matches error diff actual value",
			"    actual  : true",
			color.RedString("    actual  : true"),
		},

		//------------------------------------------------
		// Highlighted output
		//------------------------------------------------

		{
			"matches file names",
			"/some/path/to/file.go:123",
			fmt.Sprintf("/some/path/to/%s", color.New(color.FgCyan, color.Bold).Sprint("file.go:123")),
		},
		{
			"matches test name output",
			"Test:   TestSomeThing/cool",
			color.CyanString("Test:   TestSomeThing/cool"),
		},

		//------------------------------------------------
		// Prevent false positives
		//------------------------------------------------

		{
			"doesn't match invalid test pass message",
			"hello --- PASS: Foo (0.00s)",
			"hello --- PASS: Foo (0.00s)",
		},
		{
			"doesn't match invalid test fail message",
			"WORLD --- FAIL: Foo (0.00s)",
			"WORLD --- FAIL: Foo (0.00s)",
		},
		{
			"doesn't match invalid test skip message",
			"HelloWorld --- SKIP: Foo (0.00s)",
			"HelloWorld --- SKIP: Foo (0.00s)",
		},
		{
			"doesn't match error diff expected value",
			"foo expected: true",
			"foo expected: true",
		},
		{
			"doesn't match error diff actual value",
			"something actual  : foo",
			"something actual  : foo",
		},
		{
			"doesn't match invalid test missing message",
			"waht's up doc?    github.com/foo/bar/baz 0.123s",
			"waht's up doc?    github.com/foo/bar/baz 0.123s",
		},
		{
			"doesn't match invalid error prefix",
			"This is an Error: foobar",
			"This is an Error: foobar",
		},
		{
			"doesn't match invalid test prefix",
			"just a Test:   false",
			"just a Test:   false",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := internal.NewFormatter()
			assert.Exactly(t, tt.expected, f.Format(tt.input))
		})
	}
}

// func escapeString(s string) string {
// 	result := ""
// 	for _, c := range s {
// 		if c < 32 || c == 127 {
// 			result += fmt.Sprintf("\\x%02x", c)
// 		} else {
// 			result += string(c)
// 		}
// 	}
// 	return result
// }
