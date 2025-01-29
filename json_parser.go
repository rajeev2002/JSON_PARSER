package main

func check(err error){
	if err != nil{
		panic(err)
	}
}

func jsonParser(pathToFile string) string {
	tokens := lexer(pathToFile)
	return parser(tokens)
}