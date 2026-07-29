package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           int
	Name         string
	Email        string
	PasswordHash string
}

const table = "users_hash_demo"

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

	user, err := createUser(ctx, pool, "John Doe", "john.doe@example.com", "password")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Usuario criado: ID=%d Name=%s Email=%s\n", user.ID, user.Name, user.Email)
	fmt.Println("Hash salvo no banco:", user.PasswordHash)
	fmt.Println("Senha original nao foi salva no banco.")

	fmt.Println("\nTentando login com senha errada:")
	err = login(ctx, pool, "john.doe@example.com", "senha-errada")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("\nTentando login com senha correta:")
	err = login(ctx, pool, "john.doe@example.com", "password")
	if err != nil {
		fmt.Println(err)
	}
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
			email VARCHAR(255) NOT NULL UNIQUE,
			password_hash VARCHAR(255) NOT NULL
		)
	`, table))
	return err
}

func createUser(ctx context.Context, pool *pgxpool.Pool, name, email, password string) (User, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return User{}, err
	}
   // 
	var user User

	err = pool.QueryRow(
		ctx,
		fmt.Sprintf(`
			INSERT INTO %s (name, email, password_hash)
			VALUES ($1, $2, $3)
			RETURNING id, name, email, password_hash
		`, table),
		name,
		email,
		string(passwordHash),
	).Scan(&user.ID, &user.Name, &user.Email, &user.PasswordHash)

	if err != nil {
		return User{}, err
	}

	return user, nil
}

func login(ctx context.Context, pool *pgxpool.Pool, email, password string) error {
	var user User

	err := pool.QueryRow(
		ctx,
		fmt.Sprintf(`
			SELECT id, name, email, password_hash
			FROM %s
			WHERE email = $1
		`, table),
		email,
	).Scan(&user.ID, &user.Name, &user.Email, &user.PasswordHash)

	if err == pgx.ErrNoRows {
		return fmt.Errorf("usuario nao encontrado")
	}

	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return fmt.Errorf("senha invalida")
	}

	fmt.Printf("login aprovado para %s\n", user.Email)
	return nil
}
