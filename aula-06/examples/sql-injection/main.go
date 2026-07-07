package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

const table = "users_injection_demo"

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

	fmt.Println("Antes do DELETE vulneravel:")
	printUsers(ctx, pool)

	attackerInput := "' OR TRUE --"

	err = deleteByEmailVulnerable(ctx, pool, attackerInput)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\nDepois do DELETE vulneravel:")
	printUsers(ctx, pool)

	err = resetDatabase(ctx, pool)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\nAntes do DELETE seguro:")
	printUsers(ctx, pool)

	err = deleteByEmailSafe(ctx, pool, attackerInput)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\nDepois do DELETE seguro:")
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
			email VARCHAR(255) NOT NULL,
			password VARCHAR(255) NOT NULL
		)
	`, table))
	if err != nil {
		return err
	}

	users := []User{
		{Name: "John Doe", Email: "john.doe@example.com", Password: "123"},
		{Name: "Maria Silva", Email: "maria@example.com", Password: "456"},
		{Name: "Ana Souza", Email: "ana@example.com", Password: "789"},
	}

	for _, user := range users {
		_, err = pool.Exec(
			ctx,
			fmt.Sprintf(`
				INSERT INTO %s (name, email, password)
				VALUES ($1, $2, $3)
			`, table),
			user.Name,
			user.Email,
			user.Password,
		)
		if err != nil {
			return err
		}
	}

	return nil
}

func deleteByEmailVulnerable(ctx context.Context, pool *pgxpool.Pool, email string) error {
	sql := fmt.Sprintf(
		"DELETE FROM %s WHERE email = '%s'",
		table,
		email,
	)

	
	fmt.Println("\nSQL vulneravel gerado:")
	fmt.Println(sql)

	commandTag, err := pool.Exec(ctx, sql)
	if err != nil {
		return err
	}

	fmt.Println("Linhas apagadas:", commandTag.RowsAffected())
	return nil
}

func deleteByEmailSafe(ctx context.Context, pool *pgxpool.Pool, email string) error {
	sql := fmt.Sprintf(
		"DELETE FROM %s WHERE email = $1",
		table,
	)

	fmt.Println("\nSQL seguro:")
	fmt.Println(sql)
	fmt.Println("Valor enviado separado:", email)

	commandTag, err := pool.Exec(ctx, sql, email)
	if err != nil {
		return err
	}

	fmt.Println("Linhas apagadas:", commandTag.RowsAffected())
	return nil
}

func printUsers(ctx context.Context, pool *pgxpool.Pool) {
	rows, err := pool.Query(ctx, fmt.Sprintf(`
		SELECT id, name, email, password
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

		err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
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
