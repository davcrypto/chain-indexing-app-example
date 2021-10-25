package view

import (
	"fmt"

	"github.com/davcrypto/chain-indexing-app-example/internal/json"

	"github.com/crypto-com/chain-indexing/usecase/coin"

	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	_ "github.com/crypto-com/chain-indexing/test/factory"
)

type Examples interface {
	Upsert(*ExampleRow) error
}

type ExamplesView struct {
	rdb *rdb.Handle
}

func NewExamplesView(handle *rdb.Handle) Examples {
	return &ExamplesView{
		handle,
	}
}

func (exampleView *ExamplesView) Upsert(example *ExampleRow) error {
	sql, sqlArgs, err := exampleView.rdb.StmtBuilder.
		Insert(
			"view_examples",
		).
		Columns(
			"address",
			"account_type",
			"name",
			"pubkey",
			"account_number",
			"sequence_number",
			"balance",
		).
		Values(
			example.Address,
			example.Type,
			example.MaybeName,
			example.MaybePubkey,
			example.AccountNumber,
			example.SequenceNumber,
			json.MustMarshalToString(example.Balance),
		).
		Suffix("ON CONFLICT(address) DO UPDATE SET balance = EXCLUDED.balance").
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
	Address        string     `json:"address"`
	Type           string     `json:"type"`
	MaybeName      *string    `json:"name"`
	MaybePubkey    *string    `json:"pubkey"`
	AccountNumber  string     `json:"accountNumber"`
	SequenceNumber string     `json:"sequenceNumber"`
	Balance        coin.Coins `json:"balance"`
}
