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

// InitDBFunc is a function that is used to initialize database
type InitDBFunc func(conn *pgxpool.Conn) error

type DBStorage struct {
	Pool DBPool
}

func initDB(conn *pgxpool.Conn) error {
	if _, err := conn.Exec(context.Background(), `CREATE TABLE users (
		id SERIAL PRIMARY KEY,
		username VARCHAR(50) UNIQUE NOT NULL,
		email VARCHAR(50) UNIQUE NOT NULL,
		password_hash VARCHAR(256) NOT NULL,
		created_at TIMESTAMP DEFAULT current_timestamp,
		updated_at TIMESTAMP DEFAULT current_timestamp
	)`); err != nil {
		return fmt.Errorf("create users table: %w", err)
	}

	if _, err := conn.Exec(context.Background(), `CREATE TABLE login_password_pairs (
		id SERIAL PRIMARY KEY,
		user_id INT REFERENCES users(id),
		service_name VARCHAR(50),
		login VARCHAR(50),
		password_hash VARCHAR(256),
		description TEXT,
		created_at TIMESTAMP DEFAULT current_timestamp,
		updated_at TIMESTAMP DEFAULT current_timestamp
	);`); err != nil {
		return fmt.Errorf("create login_password_pairs table: %w", err)
	}

	if _, err := conn.Exec(context.Background(), `CREATE TABLE text_data (
		id SERIAL PRIMARY KEY,
		user_id INT REFERENCES users(id),
		description TEXT,
		encrypted_text TEXT,
		created_at TIMESTAMP DEFAULT current_timestamp,
		updated_at TIMESTAMP DEFAULT current_timestamp
	)`); err != nil {
		return fmt.Errorf("create text_data table: %w", err)
	}

	if _, err := conn.Exec(context.Background(), `CREATE TABLE binary_data (
		id SERIAL PRIMARY KEY,
		user_id INT REFERENCES users(id),
		description TEXT,
		encrypted_data BYTEA,
		created_at TIMESTAMP DEFAULT current_timestamp,
		updated_at TIMESTAMP DEFAULT current_timestamp
	)`); err != nil {
		return fmt.Errorf("create binary_data table: %w", err)
	}

	if _, err := conn.Exec(context.Background(), `CREATE TABLE bank_card_details (
		id SERIAL PRIMARY KEY,
		user_id INT REFERENCES users(id),
		description TEXT,
		card_name VARCHAR(50),
		encrypted_card_number VARCHAR(50),
		encrypted_expiry_date VARCHAR(10),
		encrypted_cvc VARCHAR(5),
		created_at TIMESTAMP DEFAULT current_timestamp,
		updated_at TIMESTAMP DEFAULT current_timestamp
	);`); err != nil {
		return fmt.Errorf("create bank_card_details table: %w", err)
	}

	return nil
}

func NewDBStorage(databaseURL string) (*DBStorage, error) {
	pool, err := pgxpool.New(context.Background(), databaseURL)
	if err != nil {
		return nil, fmt.Errorf("create db pool: %w", err)
	}

	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return nil, fmt.Errorf("acquire db connection: %w", err)
	}

	if err = initDB(conn); err != nil {
		return nil, fmt.Errorf("init db: %w", err)
	}

	return &DBStorage{
		Pool: pool,
	}, nil
}
