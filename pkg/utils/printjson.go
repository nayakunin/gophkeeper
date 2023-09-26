package utils

import (
	"encoding/json"
	"fmt"
)

func PrintJSON(v interface{}) error {
	jsonData, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return fmt.Errorf("unable to marshal json: %v", err)
	}

	fmt.Println(string(jsonData))

	return err
}
