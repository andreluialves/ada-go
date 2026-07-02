package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
)

func consultarUsuario(ctx context.Context, db *sql.DB, userID string) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	row := db.QueryRowContext(
		ctx,
		"SELECT nome FROM usuarios WHERE id = ?",
		userID,
	)

	var nome string

	err := row.Scan(&nome)
	if err != nil {
		return err
	}

	fmt.Println("Nome:", nome)

	return nil
}

func main() {
	db, err := sql.Open("sqlite3", "usuarios.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	consultarUsuario(context.Background(), db, "123456")
}
