package pprint

import (
	"github.com/fatih/color"
	jsoniter "github.com/json-iterator/go"
)

var jsonL = jsoniter.ConfigCompatibleWithStandardLibrary

type colorPrint func(format string, a ...interface{})

var (
	Cyan = color.Cyan
)

func PrintPrettyJSON(print colorPrint, value interface{}) error {
	bytes, err := jsonL.MarshalIndent(value, "", "  ")
	if err != nil {
		return err
	}
	print(string(bytes))
	return nil
}
