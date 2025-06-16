package db

import (
	"context"
	"log"
	"main/db/ent"
)

type DBOperation func(ctx context.Context, client *ent.Client) error

type TransactionManager interface {
	Client() *ent.Client
	Exec(ctx context.Context, operation DBOperation) error
}

// Real
type EntTransactionManager struct {
	entClient *ent.Client
}

func NewEntTransactionManager(client *ent.Client) TransactionManager {
	return &EntTransactionManager{entClient: client}
}

func (etm *EntTransactionManager) Exec(ctx context.Context, operation DBOperation) error {
	return operation(ctx, etm.entClient)
}

func (etm *EntTransactionManager) Client() *ent.Client {
	return etm.entClient
}

// Mock
type MockTransactionManager struct{}

func NewMockTransactionManager() TransactionManager {
	log.Println("Вызов NewMockTransactionManager, создание MockTransactionManager")
	return &MockTransactionManager{}
}

func (mtm *MockTransactionManager) Exec(ctx context.Context, operation DBOperation) error {
	log.Println("Вызов Exec() MockTransactionManager, возврат nil")
	return nil
}

func (mtm *MockTransactionManager) Client() *ent.Client {
	log.Println("Вызов Client(), возврат nil")
	return nil
}
