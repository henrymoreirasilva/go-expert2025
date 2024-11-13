package marshal

import (
	"encoding/json"
)

func MarchalConta(conta interface{}) (string, error) {

	jsonConta, err := json.Marshal(conta)
	if err != nil {
		return "", err
	}

	return string(jsonConta), nil

	// encoder := json.NewEncoder(os.Stdout)

	// println("Resultado NewEncoder")
	// encoder.Encode(conta)
}
