package main

import (
	"strings"
)

func jsonParser(pathToFile string) string{
	return strings.Split(strings.Split(pathToFile, "/")[3], ".")[0]
}