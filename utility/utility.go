package utility

import (
	"encoding/json"
	"fmt"
)

func ObjectToJson(obj interface{}) string {
	b, err := json.Marshal(obj)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	return string(b)
}
