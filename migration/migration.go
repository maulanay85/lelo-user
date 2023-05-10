package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	configuration "lelo-user/config"
	"os"
	"strings"
	"time"

	"github.com/jackc/pgx/v4"
)

func main() {

	err := configuration.ReadConfiguration("..")
	if err != nil {
		fmt.Printf("[migration] error read configuration: %#v\n", err)
		return
	}
	fmt.Println("initial connection to db")
	ctx := context.Background()
	urlDatabase := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		configuration.CredentialData.Database.User,
		configuration.CredentialData.Database.Password,
		configuration.ConfigData.Database.Host,
		configuration.ConfigData.Database.Port,
		configuration.ConfigData.Database.Name)
	conn, err := pgx.Connect(ctx, urlDatabase)
	if err != nil {
		fmt.Printf("[migration] error connecting to db: %#v", err)
	}
	println("successfully connect to db")

	defer conn.Close(ctx)

	queryCheckTableIsExist := `SELECT EXISTS(
			SELECT FROM information_schema.tables
			WHERE table_schema = 'public'
			AND table_name = 'schema_migration'
		);`

	var isMigrationTableExist bool
	conn.QueryRow(ctx, queryCheckTableIsExist).Scan(&isMigrationTableExist)

	if err := checkMigrationTable(ctx, conn); err != nil {
		fmt.Print(err)
		return
	}

	if err := checkFolderSqlIsExist(); err != nil {
		fmt.Print(err)
		return
	}

	var flagCmd = flag.String("type", "", "")
	flag.Parse()
	if *flagCmd == "create" {
		err := createMigrationFile()
		if err != nil {
			fmt.Printf("error: %#v", err)
			return
		}
	} else if *flagCmd == "up" {
		err := upMigration(ctx, conn)
		if err != nil {
			fmt.Printf("error up: %#v", err)
			return

		}
	}

	println("finish migration")
}

func checkMigrationTable(ctx context.Context, conn *pgx.Conn) error {
	queryCheckTableIsExist := `SELECT EXISTS(
		SELECT FROM information_schema.tables
		WHERE table_schema = 'public'
		AND table_name = 'schema_migration'
	);`

	var isMigrationTableExist bool
	conn.QueryRow(ctx, queryCheckTableIsExist).Scan(&isMigrationTableExist)

	if isMigrationTableExist {
		fmt.Println("table migration is exist, continue run migration")
		return nil
	}

	fmt.Println("table migrations is not exist, create table...")
	createTable := `CREATE TABLE schema_migration(
		id serial,
		migration_name varchar(100),
		migration_date timestamp DEFAULT CURRENT_TIMESTAMP
	);`

	_, err := conn.Exec(ctx, createTable)
	if err != nil {
		wrap := fmt.Errorf("error create table migration: %#v", err)
		return wrap
	}
	return nil
}

func checkFolderSqlIsExist() error {
	path := "./sql"
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			return fmt.Errorf("error create folder sql: %#v", err)
		}
	}
	return nil
}

func createMigrationFile() error {
	fileName := time.Now().UnixMilli()

	f, err := os.Create(fmt.Sprintf("./sql/%d.sql", fileName))
	if err != nil {
		wrap := fmt.Errorf("error create file migration: %#v", err)
		return wrap
	}
	defer f.Close()
	return nil
}

func upMigration(ctx context.Context, conn *pgx.Conn) error {
	fmt.Println("process up migration")
	rows, err := conn.Query(ctx, "SELECT migration_name from schema_migration")

	if err != nil {
		wrap := fmt.Errorf("error on up migration: %#v", rows)
		return wrap
	}
	defer rows.Close()

	row := make(map[string]int)

	for rows.Next() {
		var migrationName string
		if err := rows.Scan(&migrationName); err != nil {
			wrap := fmt.Errorf("error on read migration row: %#v", rows)
			return wrap
		}
		row[migrationName] = 1
	}

	files, err := ioutil.ReadDir("./sql")
	if err != nil {
		wrap := fmt.Errorf("error list sql file: %#v", err)
		return wrap
	}

	tx, err := conn.Begin(ctx)
	if err != nil {
		wrap := fmt.Errorf("error initial: %#v", err)
		return wrap
	}
	defer func() {
		tx.Rollback(ctx)
	}()

	for _, file := range files {
		fileName := strings.Split(file.Name(), ".")[0]
		if _, ok := row[fileName]; !ok {
			readFile, err := ioutil.ReadFile(fmt.Sprintf("./sql/%s", file.Name()))
			if err != nil {
				wrap := fmt.Errorf("error read migration file: %#v", err)
				return wrap
			}
			for _, q := range strings.Split(string(readFile), ";") {
				q := strings.TrimSpace(q)
				if q == "" {
					continue
				}
				if _, err := tx.Exec(ctx, q); err != nil {
					wrap := fmt.Errorf("error run migration file: %#v", err)
					return wrap
				}
			}

			_, err = conn.Exec(ctx, "INSERT INTO schema_migration (migration_name) VALUES($1)", fileName)
			if err != nil {
				wrap := fmt.Errorf("error save history migration file: %#v", err)
				return wrap
			}
		}
	}

	return tx.Commit(ctx)
}
