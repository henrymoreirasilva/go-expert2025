package encoder

import (
	"encoding/json"
	"os"
)

func EncoderConta(conta interface{}) {
	json.NewEncoder(os.Stdout).Encode(conta)
}
