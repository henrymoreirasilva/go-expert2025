package exemplo4

import (
	"bufio"
	"os"
)

// Lendo conteúdo de um arquivo por BUFFER
func BuffRead(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	buffer := make([]byte, 3)

	for {
		qtd, err := reader.Read(buffer)
		if err != nil {
			break
		}

		println(string(buffer[:qtd]))
	}
}
