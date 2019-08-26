package db

import (
	"context"
	"database/sql"
	"fmt"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "my-super-secret"
	dbname   = "doctor_visit"

	// DefaultLimit is the default value of numbers of data returned from the database
	// when doing GetAll
	DefaultLimit = 10
)

// ConnectionConfig is the struct for configuring the connection to the database
type ConnectionConfig struct {
	host     string
	port     int
	user     string
	password string
	dbname   string
}

// NewConnectionConfig returns a connection config based on the parameters it's given.
func NewConnectionConfig(host string, port int, user, password, dbname string) *ConnectionConfig {
	return &ConnectionConfig{
		host:     host,
		port:     port,
		user:     user,
		password: password,
		dbname:   dbname,
	}
}

// DB is the database struct used for interacting with the database
// it extends the functionality of original *sql.DB
type DB struct {
	*sql.DB

	ctx context.Context
}

// New accepts a connectionConfig as the parameter and therefore create the connection to the database
func New(connConfig *ConnectionConfig) (*DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		connConfig.host, connConfig.port, connConfig.user, connConfig.password, connConfig.dbname)

	return Newf(psqlInfo)
}

// Newf accepts connection string as the parameter and therefore create the connection to the database
func Newf(connString string) (*DB, error) {
	database, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, fmt.Errorf("Failed to open connection to database: %v", err)
	}

	// Check connection by ping-ing
	err = database.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Successfully connect to the database")

	return &DB{database, context.Background()}, nil
}

// NewDefaultConnection create connection based on the default value
func NewDefaultConnection() (*DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	return Newf(psqlInfo)
}
