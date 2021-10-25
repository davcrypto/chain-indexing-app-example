package view

import (
	"fmt"

	"github.com/davcrypto/chain-indexing-app-example/internal/json"

	"github.com/crypto-com/chain-indexing/usecase/coin"

	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	_ "github.com/crypto-com/chain-indexing/test/factory"
)

type Examples interface {
	Insert(*ExampleRow) error
}

type ExamplesView struct {
	rdb *rdb.Handle
}

func NewExamplesView(handle *rdb.Handle) Examples {
	return &ExamplesView{
		handle,
	}
}

func (exampleView *ExamplesView) Insert(example *ExampleRow) error {
	sql, sqlArgs, err := exampleView.rdb.StmtBuilder.
		Insert(
			"view_examples",
		).
		Columns(
			"address",
			"balance",
		).
		Values(
			example.Address,
			json.MustMarshalToString(example.Balance),
		).
		ToSql()

	if err != nil {
		return fmt.Errorf("error building examples insertion sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	result, err := exampleView.rdb.Exec(sql, sqlArgs...)
	if err != nil {
		return fmt.Errorf("error inserting example into the table: %v: %w", err, rdb.ErrWrite)
	}
	if result.RowsAffected() != 1 {
		return fmt.Errorf("error inserting example into the table: no rows inserted: %w", rdb.ErrWrite)
	}

	return nil
}

type ExampleRow struct {
	Address string     `json:"address"`
	Balance coin.Coins `json:"balance"`
}
