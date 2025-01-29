package main

import "os"

func checkIfSingleToken(ch rune) bool {
	singleTokens := []rune{'{', '}', ',', ':'}

	for _, c := range singleTokens{
		if c == ch{
			return true
		}
	}
	return false
}

func checkIfDel(del []rune, ch rune) bool {
	for _, d := range del{
		if d == ch{
			return true
		}
	}
	return false
}

func readVal(del []rune, data string, pos int) int{
	tem := pos+1
	n := len(data)
	for tem<n && !checkIfDel(del, rune(data[tem])){
		tem+=1
	}
	if tem != n && data[tem] == '"'{
		tem+=1
	}
	return tem
}

func lexer(pathToFile string) []string {
	var tokens []string
	file, err := os.ReadFile(pathToFile)
	check(err)
	data := string(file)
	n := len(data)
	i:=0
	for i<n {
		if data[i] == ' ' || data[i] == '\n'{
			i+=1
			continue
		}
		if checkIfSingleToken(rune(data[i])) {
			tokens = append(tokens, string(data[i]))
			i+=1
		} else{
			del := []rune{'}', ',', '\n', ' '}
			if data[i] == '"'{
				del = []rune{'"'}
			}
			pos := readVal(del, data, i)
			tokens = append(tokens, data[i:pos])
			i=pos
		}
	}
	return tokens
}