package exemplo1

import (
	"log"
	"os"
)

// Criação e escrita de conteúdo
func CreateFile(fileName string) {
	f, err := os.Create(fileName)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	qtd, err := f.WriteString("Exemplo 1")
	if err != nil {
		panic(err)
	}

	log.Print("Gravados ", qtd, " bytes")
}
