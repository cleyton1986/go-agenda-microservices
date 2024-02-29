package sqlite

import (
	"database/sql"
	"fmt"
)

// SQLiteDatabase é uma estrutura que implementa a interface Database para SQLite.
type SQLiteDatabase struct {
    DB *sql.DB
}

// NewSQLiteDatabase cria uma nova instância do banco de dados SQLite.
func NewSQLiteDatabase() *SQLiteDatabase {
    return &SQLiteDatabase{}
}

// Connect conecta ao banco de dados SQLite.
func (s *SQLiteDatabase) Connect() error {
    db, err := sql.Open("sqlite3", "data/sqlite/banco.db")
    if err != nil {
        return fmt.Errorf("falha ao conectar ao banco de dados SQLite: %v", err)
    }
    s.DB = db
    return nil
}

// Close fecha a conexão com o banco de dados SQLite.
func (s *SQLiteDatabase) Close() error {
    if s.DB != nil {
        return s.DB.Close()
    }
    return nil
}

// Implementar outros métodos para CRUD, transações, etc., conforme necessário.
