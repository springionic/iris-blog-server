// Package db
// created by lilei at 2021/11/7
package db

import (
	"entgo.io/ent/dialect/sql"
	"iris-blog-server/ent"
	"sync"
	"time"
)

var pgOnce sync.Once
var pgClient *ent.Client

func getPgClient() *ent.Client {
	driver, err := sql.Open(
		"postgres",
		"user=postgres password=123456 dbname=iris_blog sslmode=disable",
	)
	if err != nil {
		panic(err.Error())
	}
	db := driver.DB()
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)
	return ent.NewClient(ent.Driver(driver))
}

func PgClient() *ent.Client {
	pgOnce.Do(func() {
		pgClient = getPgClient()
	})
	return pgClient
}
