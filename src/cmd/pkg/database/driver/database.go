package database

import "github.com/cleyton1986/go-agenda-microservices/cmd/pkg/database/driver/sqlite"

// Database é uma interface que define operações comuns do banco de dados.
type Database interface {
    Connect() error
    Close() error
    // Outros métodos comuns de CRUD, transações, etc.
}

// NewDatabase cria uma instância do banco de dados com base no tipo de driver fornecido.
func NewDatabase(driver string) Database {
    switch driver {
    case "sqlite":
        return sqlite.NewSQLiteDatabase() // Chamando NewSQLiteDatabase do pacote sqlite
    // Outros casos para drivers de outros bancos de dados
    default:
        // Retornar um banco de dados padrão ou um erro, dependendo dos requisitos do seu projeto.
        return nil
    }
}
