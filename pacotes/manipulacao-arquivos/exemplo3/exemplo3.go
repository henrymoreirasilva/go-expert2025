package exemplo3

import "os"

// Escrita em formato de bytes
func ByteWrite(fileName string) {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	byteContent := []byte("Escrevendo em formato BYTE no arquivo")
	qtd, err := file.Write(byteContent)
	if err != nil {
		panic(err)
	}
	println("Foram gravados ", qtd, " bytes")

}
