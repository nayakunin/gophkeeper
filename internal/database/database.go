package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// DBPool is an interface for pgxpool.Pool
type DBPool interface {
	Acquire(ctx context.Context) (*pgxpool.Conn, error)
	Begin(ctx context.Context) (pgx.Tx, error)
	Ping(ctx context.Context) error
}

// Storage is a struct of the database.
type Storage struct {
	Pool DBPool
}

func initDB(ctx context.Context, conn *pgxpool.Conn) error {
	sqlCommands := `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			username VARCHAR(50) UNIQUE NOT NULL,
			password_hash BYTEA NOT NULL,
			encrypted_master_key BYTEA NOT NULL,
			created_at TIMESTAMP DEFAULT current_timestamp,
			updated_at TIMESTAMP DEFAULT current_timestamp
		);
		CREATE TABLE IF NOT EXISTS login_password_pairs (
			id SERIAL PRIMARY KEY,
			user_id INT REFERENCES users(id),
			service_name VARCHAR(50) NOT NULL,
			login VARCHAR(50) NOT NULL,
			encrypted_password BYTEA NOT NULL,
			description TEXT,
			created_at TIMESTAMP DEFAULT current_timestamp,
			updated_at TIMESTAMP DEFAULT current_timestamp
		);
		CREATE TABLE IF NOT EXISTS text_data (
			id SERIAL PRIMARY KEY,
			user_id INT REFERENCES users(id),
			description TEXT,
			encrypted_text BYTEA NOT NULL,
			created_at TIMESTAMP DEFAULT current_timestamp,
			updated_at TIMESTAMP DEFAULT current_timestamp
		);
		CREATE TABLE IF NOT EXISTS binary_data (
			id SERIAL PRIMARY KEY,
			user_id INT REFERENCES users(id),
			description TEXT,
			encrypted_data BYTEA NOT NULL,
			created_at TIMESTAMP DEFAULT current_timestamp,
			updated_at TIMESTAMP DEFAULT current_timestamp
		);
		CREATE TABLE IF NOT EXISTS bank_card_details (
			id SERIAL PRIMARY KEY,
			user_id INT REFERENCES users(id),
			description TEXT,
			card_name VARCHAR(50) NOT NULL,
			encrypted_card_number BYTEA NOT NULL,
			encrypted_expiry_date BYTEA NOT NULL,
			encrypted_cvc BYTEA NOT NULL,
			created_at TIMESTAMP DEFAULT current_timestamp,
			updated_at TIMESTAMP DEFAULT current_timestamp
		);
	`

	if _, err := conn.Exec(ctx, sqlCommands); err != nil {
		return fmt.Errorf("create tables: %w", err)
	}

	return nil
}

// NewStorage creates a new Storage struct
func NewStorage(ctx context.Context, databaseURL string) (*Storage, error) {
	pool, err := pgxpool.New(context.Background(), databaseURL)
	if err != nil {
		return nil, fmt.Errorf("create db pool: %w", err)
	}

	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return nil, fmt.Errorf("acquire db connection: %w", err)
	}

	if err = initDB(ctx, conn); err != nil {
		return nil, fmt.Errorf("init db: %w", err)
	}

	return &Storage{
		Pool: pool,
	}, nil
}
