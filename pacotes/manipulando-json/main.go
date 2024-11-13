package main

import (
	"pacotes-manipulando-json/encoder"
	"pacotes-manipulando-json/marshal"
	"pacotes-manipulando-json/unmarshal"
)

type Conta struct {
	Numero int `json:"n"`
	Saldo  int `json:"s"`
}

func main() {
	conta := Conta{
		Numero: 123,
		Saldo:  100,
	}

	res, err := marshal.MarchalConta(conta)
	if err != nil {
		panic(err)
	}
	println(res)

	encoder.EncoderConta(conta)

	contaJson := Conta{}
	err = unmarshal.MarchalConta(&contaJson)
	if err != nil {
		panic(err)
	}
	println(contaJson.Numero)
}
