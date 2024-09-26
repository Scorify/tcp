package tcp

import (
	"context"
	"encoding/json"
)

type Schema struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

func Run(ctx context.Context, config string) error {
	schema := Schema{}

	err := json.Unmarshal([]byte(config), &schema)
	if err != nil {
		return err
	}

	return nil
}
