package utils

import (
	"encoding/json"
	"fmt"
)

// PrettyPrint -
func PrettyPrint(val interface{}) {

	b, _ := json.MarshalIndent(val, "", " ")

	fmt.Printf("%s\n", b)
}
