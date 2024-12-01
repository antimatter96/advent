package common

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/mattn/go-colorable"
	"github.com/rs/zerolog"
)

var Log zerolog.Logger

func init() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	adventDebug, ok := os.LookupEnv("ADVENT_DEBUG")
	if ok && adventDebug == "true" {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	coloredStdOut := colorable.NewColorableStdout()
	output := zerolog.ConsoleWriter{Out: coloredStdOut}

	output.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("[ %-6s]", i))
	}
	output.FormatMessage = func(i interface{}) string {
		if reflect.TypeOf(i) == nil {
			return ""
		}
		return fmt.Sprintf("%s", i)
	}
	output.FormatFieldName = func(i interface{}) string {
		return fmt.Sprintf("%s: ", i)
	}
	output.FormatFieldValue = func(i interface{}) string {
		return fmt.Sprintf("%s", i)
	}
	output.FormatTimestamp = func(i interface{}) string {
		return ""
	}

	Log = zerolog.New(output)
}
