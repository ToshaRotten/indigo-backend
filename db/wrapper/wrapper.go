package wrapper

import (
	"context"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"main/db/ent"
)

type DatabaseWrapper struct {
	Client *ent.Client
}

func OpenConnection(dataSourceName string) (*DatabaseWrapper, error) {
	client, err := ent.Open("postgres", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("не удалось открыть соединение с БД: %w", err)
	}

	ctx := context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		return nil, fmt.Errorf("не удалось применить миграцию: %w", err)
	}

	fmt.Println("Клиент БД успешно инициализирован и создан")
	return &DatabaseWrapper{Client: client}, nil
}

func (dw *DatabaseWrapper) CloseConnection() error {
	if dw.Client != nil {
		log.Println("Закрытие соединения с БД")
		return dw.Client.Close()
	}

	return nil
}
