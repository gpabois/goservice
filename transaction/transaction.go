package transaction

import (
	"github.com/gpabois/gostd/option"
	"github.com/gpabois/gostd/result"
)

type ITransactionable[Tx ITransaction] interface {
	Begin(root option.Option[Transaction]) result.Result[Tx]
}

type ITransaction interface {
	Commit() result.Result[bool]
	Rollback() result.Result[bool]
	// Returns true if the transaction is managed by a parent transaction
	IsManaged() bool
}

// A transaction which handle multiple sub transactions
type Transaction struct {
	Children map[string]ITransaction
	Managed  bool
}

// Get or create the transaction
func (rtx Transaction) GetOrCreate(txName string, new func() result.Result[ITransaction]) result.Result[ITransaction] {
	tx, ok := rtx.Children[txName]

	if !ok {
		txRes := new()
		if txRes.HasFailed() {
			return result.Result[ITransaction]{}.Failed(txRes.UnwrapError())
		}
		rtx.Children[txName] = txRes.Expect()
	}

	return result.Success(tx)
}

func (rtx Transaction) Commit() result.Result[bool] {
	for _, tx := range rtx.Children {
		res := tx.Commit()
		if res.HasFailed() {
			return result.Result[bool]{}.Failed(res.UnwrapError())
		}
	}

	return result.Success(true)
}

func (rtx Transaction) Rollback() result.Result[bool] {
	for _, tx := range rtx.Children {
		tx.Rollback()
	}
	return result.Success(true)
}

func (rtx Transaction) IsManaged() bool {
	return rtx.Managed
}

type TransactionFunc[Tx ITransaction, R any] func(tx Tx) result.Result[R]

func With[Tx ITransaction, R any](txRes result.Result[Tx], transaction TransactionFunc[Tx, R]) result.Result[R] {
	if txRes.HasFailed() {
		return result.Result[R]{}.Failed(txRes.UnwrapError())
	}

	tx := txRes.Expect()

	// Rollback on panic
	if !tx.IsManaged() {
		defer func() {
			if r := recover(); r != nil {
				tx.Rollback()
			}
		}()

	}

	res := transaction(tx)

	if !tx.IsManaged() {
		if res.HasFailed() {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}

	return res
}

func Begin[R any](transaction TransactionFunc[Transaction, R]) result.Result[R] {
	tx := Transaction{}
	return With(result.Success(tx), transaction)
}
