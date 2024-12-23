package main

import (
	"os"
	"strings"
	"testing"
	"unicode"
)

func panicIfErr(err error){
	if err != nil{
		panic(err)
	}
}

func getAnswer(name string) string{
	val := strings.Split(name, ".")[0]
	i := len(val) - 1
	for i>=0{
		if unicode.IsDigit(rune(val[i])){
			i-=1
		}else{
			break
		}
	}
	val = val[:i+1]
	return val
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
			if result != getAnswer(testFileName){
				t.Errorf("File - %v in folder - %v failed the test.", testFileName, filePathName)
			}
		}
	}
}