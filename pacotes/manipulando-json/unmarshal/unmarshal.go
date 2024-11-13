package unmarshal

import (
	"encoding/json"
)

func MarchalConta(conta any) error {

	jsonByte := []byte(`{"n": 999, "s": 122}`)

	err := json.Unmarshal(jsonByte, &conta)
	if err != nil {
		return err
	}
	return nil
}
