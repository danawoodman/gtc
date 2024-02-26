package internal

import (
	"regexp"

	"github.com/fatih/color"
)

type (
	Formatter interface {
		Format(string) string
	}
	Colorizer func(format string, a ...interface{}) string

	formatter struct{}
)

func NewFormatter() Formatter {
	return &formatter{}
}

func (f formatter) Format(line string) string {
	boldCyan := color.New(color.FgCyan, color.Bold)
	tokenPatterns := map[string]Colorizer{
		`@@ -\d+ \+\d+ @@`: color.BlackString,
	}

	for pattern, colorizer := range tokenPatterns {
		line = replaceFragment(pattern, line, colorizer)
	}

	// highlight file/line numbers, except for error messages (as they're already highlighted)
	errorTraceLine, _ := regexp.MatchString(`^\s*Error Trace:`, line)
	errorLine, _ := regexp.MatchString(`^\s*Error:`, line)
	if !errorTraceLine && !errorLine {
		line = replaceFragment(`[a-zA-Z0-9_-]+\.go:\d+`, line, boldCyan.SprintfFunc())
	}

	linePatterns := map[string]Colorizer{
		//-------------------------------------------------------
		// Successful outputs
		//-------------------------------------------------------
		`^\s*--- PASS:`:     color.GreenString,
		`^PASS$`:            color.GreenString,
		`^\s*--- Expected$`: color.GreenString,
		`^\s*expected:`:     color.GreenString,
		`^ok  `:             color.GreenString,
		//-------------------------------------------------------
		// Failure outputs
		//-------------------------------------------------------
		`^\s*--- FAIL:`:      color.RedString,
		`^FAIL`:              color.RedString,
		`^\s*Error Trace:`:   color.RedString,
		`^\s*Error:`:         color.RedString,
		`^\s*\+\+\+ Actual$`: color.RedString,
		`^\s*actual  :`:      color.RedString,
		//-------------------------------------------------------
		// Highlight outputs
		//-------------------------------------------------------
		`^\s*Test:   `: color.CyanString,
		//-------------------------------------------------------
		// Skipped test output
		//-------------------------------------------------------
		`^\s*--- SKIP:`: color.YellowString,
		//-------------------------------------------------------
		// Muted outputs
		//-------------------------------------------------------
		`^=== RUN`:   color.BlackString,
		`^=== PAUSE`: color.BlackString,
		`^=== CONT`:  color.BlackString,
		`^\?   `:     color.BlackString,
	}

	for pattern, colorizer := range linePatterns {
		if match, _ := regexp.MatchString(pattern, line); match {
			line = colorizer(line)
		}
	}

	// Fallback to printing the line as is
	return line
}

func replaceFragment(pattern, line string, colorizer Colorizer) string {
	matches := regexp.MustCompile(pattern).FindAllStringIndex(line, -1)
	for i := len(matches) - 1; i >= 0; i-- {
		start, end := matches[i][0], matches[i][1]
		line = line[:start] + colorizer(line[start:end]) + line[end:]
	}
	return line
}
