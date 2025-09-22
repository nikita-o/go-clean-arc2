package database

import (
	"clean-arc-2/internal/infrastructure/logger"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

var log = logger.GetLogger()

func NewPostgresDatabase(connString string) *pgxpool.Pool {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Разбираем конфигурацию
	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		panic(fmt.Errorf("Failed to parse database URL: %v\n", err).Error())
	}

	// Настройка параметров пула соединений
	config.MaxConns = 3                        // Максимальное количество соединений
	config.MinConns = 1                        // Минимальное количество соединений
	config.MaxConnIdleTime = 5 * time.Minute   // Максимальное время простоя соединения
	config.MaxConnLifetime = 30 * time.Minute  // Максимальная продолжительность использования соединения
	config.HealthCheckPeriod = 1 * time.Minute // Частота проверки здоровья соединений

	// Создаем пул соединений
	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		panic(fmt.Errorf("Unable to create database connection pool: %v\n", err).Error())
	}

	// Проверка соединения
	err = pool.Ping(ctx)
	if err != nil {
		panic(fmt.Errorf("Unable to connect to database: %v\n", err).Error())
	}

	log.Info("Successfully connected to the database")
	return pool
}
