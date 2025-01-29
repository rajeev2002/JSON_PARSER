package main

type Checks struct{
	canBeAtBegin bool
	canBeAtEnd bool
	previousTokens []string
	nextTokens []string
}

var checksMap = map[string]Checks{
	"{": {
	canBeAtBegin: true, 
	canBeAtEnd: false, 
	previousTokens: []string{""}, 
	nextTokens: []string{"string", "}"},
	},
	"}": {
	canBeAtBegin: false, 
	canBeAtEnd: true, 
	previousTokens: []string{"string", "value"}, 
	nextTokens: []string{""},
	},
	":": {
	canBeAtBegin: false, 
	canBeAtEnd: false, 
	previousTokens: []string{"string"}, 
	nextTokens: []string{"value"},
	},
	",": {
	canBeAtBegin: false, 
	canBeAtEnd: false, 
	previousTokens: []string{"value"}, 
	nextTokens: []string{"string"},
	},
	"string": {
	canBeAtBegin: false, 
	canBeAtEnd: false, 
	previousTokens: []string{",", "{"}, 
	nextTokens: []string{":"},
	},
	"value": {
	canBeAtBegin: false, 
	canBeAtEnd: false, 
	previousTokens: []string{":"}, 
	nextTokens: []string{"}"},
	},
}

func giveTokenKind(token string, previousToken string) string {
	if checkIfSingleToken(rune(token[0])){
		return token
	}else if token[0] == '"'{
		if previousToken == ":"{
			return "value"
		}
		return "string"
	}
	return "value"
}

func isIn(target string, list []string) bool {
	for _, val := range list{
		if val == target{
			return true
		}
	}
	return false
}

func stringCheck(val string) bool {
	n := len(val)
	return val[n-1] == '"'
}

func numberCheck(val string) bool {
	// for _, ch := range val{
	// }
	return true
}

func booleanCheck(val string) bool {
	return val == "true" || val == "false"
}

func nullCheck(val string) bool {
	return val == "null"
}

func valueCheck(val string) bool {
	if val[0] == '"'{
		return stringCheck(val)
	}else if (val[0]>='0' && val[0]<='9') || val[0] == '-'{
		return numberCheck(val)
	}else{
		return booleanCheck(val) || nullCheck(val)
	}
}

func checker(pos int, dataLen int, token string, previousToken string, nextToken string, tokenKind string) bool{
	checks := checksMap[tokenKind]
	if checks.canBeAtBegin != (pos == 0){
		return false
	}
	if checks.canBeAtEnd != (pos+1 == dataLen){
		return false
	}
	if !isIn(previousToken, checks.previousTokens) || !isIn(nextToken, checks.nextTokens){
		return false
	}
	if tokenKind == "string"{
		return stringCheck(token)
	}
	if tokenKind == "value"{
		return valueCheck(token)
	}
	return true
}

func parser(tokens []string) string {
	n := len(tokens)
	if n == 0{
		return  "invalid"
	}
	for i, token := range tokens {
		var nextToken, previousToken string
		if i+1<n{
			nextToken = tokens[i+1]
		}
		if i > 0{
			previousToken = tokens[i-1]
		}
		if !checker(i, n, token, previousToken, nextToken, giveTokenKind(token, previousToken)){
			return "invalid"
		}
	}
	return "valid"
}