package main

import (
	"pacotes-manipulacao-arquivos/exemplo1"
	"pacotes-manipulacao-arquivos/exemplo2"
	"pacotes-manipulacao-arquivos/exemplo3"
	"pacotes-manipulacao-arquivos/exemplo4"
)

func main() {
	println("Criando arquivo e gravando conteúdo")
	exemplo1.CreateFile("exemplo-01.txt")

	println("Lendo conteúdo com ReadFile()")
	exemplo2.ReadFile("exemplo-01.txt")

	println("Salvando bytes no arquivo")
	exemplo3.ByteWrite("exemplo-01.txt")

	println("Lendo conteúdo por bufferização")
	exemplo4.BuffRead("exemplo-01.txt")
}
