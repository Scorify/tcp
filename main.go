package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
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

	connStr := fmt.Sprintf("%s:%d", schema.Host, schema.Port)

	dialer := &net.Dialer{}
	conn, err := dialer.DialContext(ctx, "tcp", connStr)
	if err != nil {
		return fmt.Errorf("failed to connect to %q; %w", connStr, err)
	}
	defer conn.Close()

	return nil
}
