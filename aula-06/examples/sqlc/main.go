package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	sqlcdb "db/examples/sqlc/db"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	ctx := context.Background()

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "postgres://app:app@localhost:5432/app?sslmode=disable"
	}

	pool, err := pgxpool.New(ctx, databaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	queries := sqlcdb.New(pool)

	err = resetDatabase(ctx, pool)
	if err != nil {
		log.Fatal(err)
	}

	email := fmt.Sprintf("sqlc.%d@example.com", time.Now().UnixNano())

	createdUser, err := queries.CreateUser(ctx, sqlcdb.CreateUserParams{
		Name:     text("User sqlc"),
		Email:    text(email),
		Password: text("password"),
	})
	if err != nil {
		log.Fatal(err)
	}
	printUser("criado", createdUser)

	foundUser, err := queries.GetUser(ctx, createdUser.ID)
	if err != nil {
		log.Fatal(err)
	}
	printUser("buscado", foundUser)

	updatedUser, err := queries.UpdateUser(ctx, sqlcdb.UpdateUserParams{
		ID:       createdUser.ID,
		Name:     text("User sqlc atualizado"),
		Email:    text(email),
		Password: text("new-password"),
	})
	if err != nil {
		log.Fatal(err)
	}
	printUser("atualizado", updatedUser)

	users, err := queries.ListUsers(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("total de usuarios:", len(users))

	rowsAffected, err := queries.DeleteUser(ctx, createdUser.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("linhas removidas:", rowsAffected)
}

func printUser(label string, user sqlcdb.UsersSqlcDemo) {
	fmt.Printf(
		"%s: id=%s name=%s email=%s password=%s\n",
		label,
		user.ID.String(),
		user.Name.String,
		user.Email.String,
		user.Password.String,
	)
}

func resetDatabase(ctx context.Context, pool *pgxpool.Pool) error {
	_, err := pool.Exec(ctx, "DROP TABLE IF EXISTS users_sqlc_demo")
	if err != nil {
		return err
	}

	_, err = pool.Exec(ctx, `
		CREATE TABLE users_sqlc_demo (
			id UUID PRIMARY KEY DEFAULT uuidv7(),
			name VARCHAR(255),
			email VARCHAR(255),
			password VARCHAR(255)
		)
	`)
	return err
}

func text(value string) pgtype.Text {
	return pgtype.Text{
		String: value,
		Valid:  true,
	}
}
