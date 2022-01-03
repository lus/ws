package out

import (
	"fmt"
	"github.com/logrusorgru/aurora/v3"
	"os"
)

const (
	prefixInfo  = "INFO  "
	prefixError = "ERROR "

	separator = " ::  "
)

var au = aurora.NewAurora(true)

// DisableFormatting disables output formatting
func DisableFormatting() {
	au = aurora.NewAurora(false)
}

// Info prints an info message
func Info(msg string) {
	fmt.Printf("%s%s%s\n", au.Bold(au.Blue(prefixInfo)), au.Gray(10, separator), msg)
}

// Error prints an error message and calls os.Exit(-1) afterwards
func Error(msg string) {
	fmt.Printf("%s%s%s\n", au.Bold(au.Red(prefixError)), au.Gray(10, separator), msg)
	os.Exit(-1)
}
