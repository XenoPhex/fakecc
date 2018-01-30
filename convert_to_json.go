package main

import (
	"bytes"
	"encoding/json"
)

func ConvertToJSON(obj map[string]interface{}) string {
	buff := new(bytes.Buffer)
	encoder := json.NewEncoder(buff)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "  ")
	encoder.Encode(obj)
	return string(buff.Bytes())
}
