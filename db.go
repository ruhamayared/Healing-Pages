package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
)

func ConnectDB() (*pgx.Conn, error) {
	// Connect to the database using the connection string
	conn, err := pgx.Connect(context.Background(), "postgresql://ruhamayared:v2_42GRu_69gS2m3p4WiNKVKeb4N7jYj@db.bit.io:5432/ruhamayared/healing-pages")
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	return conn, nil
}
