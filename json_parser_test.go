package main

import (
	"os"
	"strings"
	"testing"
)

func panicIfErr(err error){
	if err != nil{
		panic(err)
	}
}

func TestJsonParser(t *testing.T){
	files, err := os.ReadDir("./tests")
	panicIfErr(err)
	for _, filePath := range files{
		filePathName := filePath.Name()
		test_files, err := os.ReadDir("./tests/"+ filePathName)
		panicIfErr(err)
		for _, testFile := range test_files{
			testFileName := testFile.Name()
			testFilePath := "./tests/" + filePath.Name() + "/" + testFileName
			result := jsonParser(testFilePath)
			if result != strings.Split(testFileName, ".")[0]{
				t.Errorf("File - %v in folder - %v failed the test.", testFileName, filePathName)
			}
		}
	}
}