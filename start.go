package main

import (
	"context"
	"entdemo/ent/user"
	"fmt"
	"log"

	"entdemo/ent"

	_ "github.com/mattn/go-sqlite3"
)

func CreateUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.User.
		Create().
		SetAge(30).
		SetName("a8m").
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("user was created: ", u)
	return u, nil
}

func QueryUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.User.
		Query().
		Where(user.Name("a8m")).
		// `Only` fails if no user found,
		// or more than 1 user returned.
		First(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}
	log.Println("user returned: ", u)
	return u, nil
}

func QuerySql(ctx context.Context, client *ent.Client, sql string) error {
	rows, err := client.QueryContext(ctx, sql)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	names := make([]string, 0)

	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			// Check for a scan error.
			// Query rows will be closed with defer.
			log.Fatal(err)
		}
		names = append(names, name)
	}

	log.Println("user returned: ", names)
	return err
}

func main() {
	client, err := ent.Open("sqlite3", "file:ent.db?_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	//
	ctx := context.Background()
	_, err = CreateUser(ctx, client)
	if err != nil {
		fmt.Println(err)
		return
	}

	//
	_, err = QueryUser(ctx, client)
	if err != nil {
		fmt.Println(err)
		return
	}

	// sql
	sql := "select name from users"
	err = QuerySql(ctx, client, sql)
	if err != nil {
		fmt.Println(err)
		return
	}
}
