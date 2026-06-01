package main

import (
	"errors"
	"flag"
	"fmt"

	// Библиотека для миграций
	"github.com/golang-migrate/migrate/v4"
	// Драйвер для выполнения миграций PostgreSQL
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	// Драйвер для получения миграций из файлов
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	var (
		dbHost          string
		dbPort          int
		dbUser          string
		dbPassword      string
		dbName          string
		dbSSLMode       string
		migrationsPath  string
		migrationsTable string
	)

	// Получаем параметры подключения к PostgreSQL из флагов запуска
	flag.StringVar(&dbHost, "db-host", "localhost", "PostgreSQL host")
	flag.IntVar(&dbPort, "db-port", 5432, "PostgreSQL port")
	flag.StringVar(&dbUser, "db-user", "", "PostgreSQL user")
	flag.StringVar(&dbPassword, "db-password", "", "PostgreSQL password")
	flag.StringVar(&dbName, "db-name", "", "PostgreSQL database name")
	flag.StringVar(&dbSSLMode, "db-sslmode", "disable", "PostgreSQL SSL mode (disable, require, etc.)")

	// Путь до папки с миграциями
	flag.StringVar(&migrationsPath, "migrations-path", "", "path to migrations")
	// Таблица, в которой будет храниться информация о миграциях
	flag.StringVar(&migrationsTable, "migrations-table", "schema_migrations", "name of migrations table")

	flag.Parse() // Выполняем парсинг флагов

	// Валидация обязательных параметров
	if dbUser == "" {
		panic("db-user is required")
	}
	if dbName == "" {
		panic("db-name is required")
	}
	if migrationsPath == "" {
		panic("migrations-path is required")
	}

	// Формируем строку подключения (DSN) для golang-migrate.
	// Параметр x-migrations-table задает имя служебной таблицы.
	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s&x-migrations-table=%s",
		dbUser, dbPassword, dbHost, dbPort, dbName, dbSSLMode, migrationsTable,
	)

	// Создаем объект мигратора
	m, err := migrate.New(
		"file://"+migrationsPath,
		dbURL,
	)
	if err != nil {
		panic(err)
	}

	// Выполняем миграции до последней версии
	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("no migrations to apply")
			return
		}

		panic(err)
	}

	fmt.Println("migrations applied successfully")
}
