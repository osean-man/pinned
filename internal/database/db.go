package database

import (
	"database/sql"
	"fmt"
	"github.com/charmbracelet/log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

type Pin struct {
	ID        int    `json:"id"`
	Command   string `json:"command"`
	Timestamp string `json:"timestamp"`
}

// InitializeDB creates the database if it doesn't exist and returns a connection handle
func InitializeDB() (*sql.DB, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("error getting home directory: %w", err)
	}

	configDir := filepath.Join(homeDir, ".config", "pinned")
	err = os.MkdirAll(configDir, 0755)
	if err != nil {
		return nil, fmt.Errorf("error creating config directory: %w", err)
	}

	dbPath := filepath.Join(configDir, "commands.db")
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	sqlStmt := ` 
    CREATE TABLE IF NOT EXISTS pins (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        command TEXT NOT NULL,
        timestamp DATETIME DEFAULT CURRENT_TIMESTAMP
    );
    `
	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, fmt.Errorf("%q: %w", err, sqlStmt)
	}

	return db, nil
}

func GetPinByID(db *sql.DB, id int) (string, error) {
	var command string
	err := db.QueryRow("SELECT command FROM pins WHERE id = ?", id).Scan(&command)
	if err != nil {
		return "", fmt.Errorf("error fetching command with ID %d: %w", id, err)
	}
	return command, nil
}

func AddPin(db *sql.DB, command string) error {
	log.Infof("Adding command: %v", command)
	stmt, err := db.Prepare("INSERT INTO pins(command) VALUES(?)")
	if err != nil {
		return fmt.Errorf("error preparing statement: %w", err)
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			log.Errorf("error closing statement: %v", err)
		}
	}(stmt)

	_, err = stmt.Exec(command)
	if err != nil {
		return fmt.Errorf("error executing insert: %w", err)
	}
	return nil
}

func GetPins(db *sql.DB) ([]Pin, error) {
	rows, err := db.Query("SELECT id, command, timestamp FROM pins")
	if err != nil {
		return nil, fmt.Errorf("error executing select: %w", err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Errorf("error closing rows: %v", err)
		}
	}(rows)

	var pins []Pin
	for rows.Next() {
		var pin Pin
		if err := rows.Scan(&pin.ID, &pin.Command, &pin.Timestamp); err != nil {
			return nil, err // Or handle potential error here
		}
		pins = append(pins, pin)
	}

	return pins, nil
}

func RemovePin(db *sql.DB, id int) error {
	stmt, err := db.Prepare("DELETE FROM pins WHERE id = ?")
	if err != nil {
		return fmt.Errorf("error preparing statement: %w", err)
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			log.Errorf("error closing statement: %v", err)
		}
	}(stmt)

	_, err = stmt.Exec(id)
	if err != nil {
		return fmt.Errorf("error executing delete: %w", err)
	}
	return nil
}

func UpdatePin(db *sql.DB, id int, command string) error {
	stmt, err := db.Prepare("UPDATE pins SET command = ? WHERE id = ?")
	if err != nil {
		return fmt.Errorf("error preparing statement: %w", err)
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			log.Errorf("error closing statement: %v", err)
		}
	}(stmt)

	_, err = stmt.Exec(command, id)
	if err != nil {
		return fmt.Errorf("error executing update: %w", err)
	}
	return nil
}
