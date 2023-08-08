package main

import (
	"OLC2_Proyecto1_202111478/TswiftP"
	"fmt"
	"log"

	"github.com/antlr4-go/antlr/v4"
)

func main() {
	fmt.Println("hola mundo")
	//abrir archivos
	archivo, err := antlr.NewFileStream("entrada.txt")
	if err != nil {
		log.Fatal(err)
	}
	//se ejecuta el lexer
	lexer := TswiftP.NewTswift_Lexer(archivo)
	//se ven los tokens
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	//se ve el parser
	parser := TswiftP.NewTswift_GrammarParser(tokens)
	//se visita por medio de visitor
	visitor := NewTswiftPVisitorImp()
	//comenzar el arbol
	arbol := parser.S() // S es la produccion inicial

	fmt.Println(visitor.Visit(arbol))
}
