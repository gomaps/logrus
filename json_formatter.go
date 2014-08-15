package logrus

import (
	"encoding/json"
	"fmt"
)

type JSONFormatter struct {
}

func (f *JSONFormatter) Format(entry *Entry) ([]byte, error) {
	prefixFieldClashes(entry)

	// Copy our context fields to the standard data fields
	if entry.Logger.ContextFields != nil {
		for k, v := range *entry.Logger.ContextFields {
			if _, ok := entry.Data[k]; !ok {
				entry.Data[k] = v
			}
		}
	}

	serialized, err := json.Marshal(entry.Data)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal fields to JSON, %v", err)
	}
	return append(serialized, '\n'), nil
}
