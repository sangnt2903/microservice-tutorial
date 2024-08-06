package json_helper

import "encoding/json"

func SafeString(req any) string {
	bytesData, _ := json.Marshal(req)
	return string(bytesData)
}
