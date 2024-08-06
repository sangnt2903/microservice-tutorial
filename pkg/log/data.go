package log

type data struct {
	Level    string                 `json:"level"`
	Time     string                 `json:"time"`
	Event    string                 `json:"event"`
	Metadata map[string]interface{} `json:"metadata"`
}
