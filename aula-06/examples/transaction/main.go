package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type User struct {
	ID    int
	Name  string
	Email string
}

const table = "users_transaction_demo"

func main() {
	ctx := context.Background()

	pool, err := pgxpool.New(ctx, "postgres://app:app@localhost:5432/app")
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	err = resetDatabase(ctx, pool)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Transacao com erro:")
	err = createTwoUsers(ctx, pool,
		User{Name: "John Doe", Email: "john@example.com"},
		User{Name: "John Repetido", Email: "john@example.com"},
	)
	if err != nil {
		fmt.Println("erro:", err)
	}

	fmt.Println("Usuarios depois da transacao com erro:")
	printUsers(ctx, pool)

	fmt.Println("\nTransacao sem erro:")
	err = createTwoUsers(ctx, pool,
		User{Name: "Maria Silva", Email: "maria@example.com"},
		User{Name: "Ana Souza", Email: "ana@example.com"},
	)
	if err != nil {
		fmt.Println("erro:", err)
	}

	fmt.Println("Usuarios depois da transacao sem erro:")
	printUsers(ctx, pool)
}

func resetDatabase(ctx context.Context, pool *pgxpool.Pool) error {
	_, err := pool.Exec(ctx, fmt.Sprintf("DROP TABLE IF EXISTS %s", table))
	if err != nil {
		return err
	}

	_, err = pool.Exec(ctx, fmt.Sprintf(`
		CREATE TABLE %s (
			id SERIAL PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			email VARCHAR(255) NOT NULL UNIQUE
		)
	`, table))
	return err
}

func createTwoUsers(ctx context.Context, pool *pgxpool.Pool, first User, second User) error {
	tx, err := pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(
		ctx,
		fmt.Sprintf(`
			INSERT INTO %s (name, email)
			VALUES ($1, $2)
		`, table),
		first.Name,
		first.Email,
	)
	if err != nil {
		return fmt.Errorf("insert first user: %w", err)
	}

	_, err = tx.Exec(
		ctx,
		fmt.Sprintf(`
			INSERT INTO %s (name, email)
			VALUES ($1, $2)
		`, table),
		second.Name,
		second.Email,
	)
	if err != nil {
		return fmt.Errorf("insert second user: %w", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("commit transaction: %w", err)
	}

	return nil
}

func printUsers(ctx context.Context, pool *pgxpool.Pool) {
	rows, err := pool.Query(ctx, fmt.Sprintf(`
		SELECT id, name, email
		FROM %s
		ORDER BY id
	`, table))
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	count := 0
	for rows.Next() {
		var user User

		err = rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			log.Fatal(err)
		}

		count++
		fmt.Printf("%+v\n", user)
	}

	if rows.Err() != nil {
		log.Fatal(rows.Err())
	}

	if count == 0 {
		fmt.Println("Nenhum usuario encontrado")
	}
}
