package data

import (
	"context"
	"database/sql"

	"gorm.io/gorm"
)

type TransactionManager struct {
	ctx context.Context
}

func NewTransactionManager(ctx context.Context) *TransactionManager {
	return &TransactionManager{ctx: ctx}
}

// 执行需要事务的操作
func (m *TransactionManager) ExecuteTransaction(fn func(txCtx context.Context) error, opts ...*sql.TxOptions) error {
	return globalDb.Transaction(func(tx *gorm.DB) error {

		txCtx := context.WithValue(m.ctx, txKey, tx)

		return fn(txCtx)
	}, opts...)
}
