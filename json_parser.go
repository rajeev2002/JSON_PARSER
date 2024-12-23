package main

import (
	"os"
)

func check(err error){
	if err != nil{
		panic(err)
	}
}

func lexer(pathToFile string) []string {
	var tokens []string
	file, err := os.ReadFile(pathToFile)
	check(err)
	data := string(file)
	n := len(data)
	i:=0
	for i<n {
		if data[i] == ' '{
			i+=1
			continue
		}
		if data[i] == '{'{
			tokens = append(tokens, string(data[i]))
		} else if data[i] == '}'{
			tokens = append(tokens, string(data[i]))
		}
		i+=1
	}
	return tokens
}

func parser(tokens []string) string {
	if len(tokens) == 0{
		return "invalid"
	}
	for i, token := range tokens {
		if token == "{"{
			if i+1 == len(tokens) || tokens[i+1] != "}"{
				return "invalid"
			}
		}else if token == "}"{
			if i+1 != len(tokens){
				return "invalid"
			}
		}
	}
	return "valid"
}

func jsonParser(pathToFile string) string {
	tokens := lexer(pathToFile)
	return parser(tokens)
}