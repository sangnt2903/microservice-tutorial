package log

import (
	"encoding/json"
	"time"
)

type Field struct {
	Key   string
	Value interface{}
}

func NewField(key string, value interface{}) Field {
	return Field{Key: key, Value: value}
}

func Error(err error) Field {
	return Field{Key: "error", Value: err.Error()}
}

func getData(level string, event func() string, fields ...Field) string {
	metadata := make(map[string]interface{})
	for _, field := range fields {
		metadata[field.Key] = field.Value
	}

	var d = &data{
		Level:    level,
		Time:     time.Now().UTC().Format(time.RFC3339),
		Event:    event(),
		Metadata: metadata,
	}

	dataBytes, _ := json.Marshal(d)
	return string(dataBytes)
}

type Tag struct {
	Key   string
	Value string
}
