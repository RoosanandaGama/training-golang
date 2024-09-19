package main

import (
	"context"
	"fmt"
	"log"
	"training-golang/session-6-db-pgx/entity"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	dsn := "postgresql://postgres:P@ssw0rd@localhost:5432/training_golang"
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatalln(err)
	}

	err = pool.Ping(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("successfully connect to db")

	var u entity.User
	err = pool.QueryRow(ctx, "select id,name from users order by id desc limit 1").Scan(&u.ID, &u.Name)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("user retrieved", u)

	_, err = pool.Exec(ctx, "insert into users(name,email,password,created_at,updated_at) "+
		"values ('test','test@email.com','test123',NOW(),NOW())")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("succesfully inserted")
	var users []entity.User
	rows, err := pool.Query(ctx, "select id,name from users order by id desc")
	if err != nil {
		log.Panicln(err)
	}
	for rows.Next() {
		var user entity.User
		rows.Scan(&user.ID, &user.Name)
		users = append(users, user)
	}
	fmt.Println("all user retrieved", users)
}
