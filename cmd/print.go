package cmd

import (
	"github.com/fatih/color"
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type colorPrint func(format string, a ...interface{})

// Color Print type
var (
	Cyan  = color.Cyan
	White = color.White
)

func printPrettyJSON(print colorPrint, value interface{}) error {
	bytes, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		return err
	}
	print(string(bytes))
	return nil
}
