package exemplo2

import (
	"os"
)

// Recuperando o conteúdo de um arquivo de uma só vez (ideal para arquivos pequenos)
func ReadFile(nameFile string) {
	byteContent, err := os.ReadFile(nameFile)
	if err != nil {
		panic(err)
	}

	println(string(byteContent))
}
