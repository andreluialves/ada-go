package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

const (
	databaseURL = "postgres://app:app@localhost:5432/app?sslmode=disable"
	table       = "users_main_demo"
)

func main() {
	ctx := context.Background()

	user := User{
		Name:     "John Doe",
		Email:    "john.doe@example.com",
		Password: "password",
	}

	// Exemplo 1: conexao simples usando pool com configuracao padrao.
	// pool, err := pgxpool.New(ctx, databaseURL)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer pool.Close()

	// Exemplo 2: pool configurado manualmente.
	config, err := pgxpool.ParseConfig(databaseURL)
	if err != nil {
		log.Fatal(err)
	}

	config.MaxConns = 10                      // limite maximo de conexoes abertas
	config.MinConns = 2                       // conexoes minimas sempre prontas
	config.MaxConnIdleTime = 10 * time.Second // fecha conexao parada por muito tempo
	config.MaxConnLifetime = 30 * time.Second // recicla conexao depois de certa idade

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	// Exemplo 3: pgx.Connect abre uma unica conexao, sem pool.
	// Para testar esse exemplo isolado, chame exampleSingleConnection(ctx).

	commandTag, err := createTable(ctx, pool)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Tabela pronta:", commandTag)

	user, err = createUser(ctx, pool, user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Usuario criado: %+v\n", user)

	foundUser, err := findUserByID(ctx, pool, user.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Usuario encontrado: %+v\n", foundUser)

	err = updateUserEmail(ctx, pool, user.ID, "john.updated@example.com")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Usuarios cadastrados:")
	printUsers(ctx, pool)

	err = deleteUser(ctx, pool, user.ID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Usuarios depois do delete:")
	printUsers(ctx, pool)
}

func createTable(ctx context.Context, pool *pgxpool.Pool) (pgconn.CommandTag, error) {
	return pool.Exec(ctx, fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			id SERIAL PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			email VARCHAR(255) NOT NULL,
			password VARCHAR(255) NOT NULL
		)
	`, table))
}

func createUser(ctx context.Context, pool *pgxpool.Pool, user User) (User, error) {
	err := pool.QueryRow(
		ctx,
		fmt.Sprintf(`
			INSERT INTO %s (name, email, password)
			VALUES ($1, $2, $3)
			RETURNING id
		`, table),
		user.Name,
		user.Email,
		user.Password,
	).Scan(&user.ID)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func findUserByID(ctx context.Context, pool *pgxpool.Pool, id int) (User, error) {
	var user User

	err := pool.QueryRow(
		ctx,
		fmt.Sprintf(`
			SELECT id, name, email, password
			FROM %s
			WHERE id = $1
		`, table),
		id,
	).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func updateUserEmail(ctx context.Context, pool *pgxpool.Pool, id int, email string) error {
	commandTag, err := pool.Exec(
		ctx,
		fmt.Sprintf(`
			UPDATE %s
			SET email = $1
			WHERE id = $2
		`, table),
		email,
		id,
	)
	if err != nil {
		return err
	}

	if commandTag.RowsAffected() == 0 {
		return fmt.Errorf("nenhum usuario encontrado com id %d", id)
	}

	return nil
}

func deleteUser(ctx context.Context, pool *pgxpool.Pool, id int) error {
	commandTag, err := pool.Exec(
		ctx,
		fmt.Sprintf(`
			DELETE FROM %s
			WHERE id = $1
		`, table),
		id,
	)
	if err != nil {
		return err
	}

	if commandTag.RowsAffected() == 0 {
		return fmt.Errorf("nenhum usuario encontrado com id %d", id)
	}

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

func exampleSingleConnection(ctx context.Context) error {
	conn, err := pgx.Connect(ctx, databaseURL)
	if err != nil {
		return err
	}
	defer conn.Close(ctx)

	var now time.Time
	err = conn.QueryRow(ctx, "SELECT NOW()").Scan(&now)
	if err != nil {
		return err
	}

	fmt.Println("Horario do banco:", now.Format(time.RFC3339))
	return nil
}
