package example

import (
	"fmt"

	applogger "github.com/crypto-com/chain-indexing/external/logger"

	//example_view "github.com/davcrypto/chain-indexing-app-example/projections/example/view"

	"github.com/crypto-com/chain-indexing/appinterface/projection/rdbprojectionbase"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	event_entity "github.com/crypto-com/chain-indexing/entity/event"
	event_usecase "github.com/crypto-com/chain-indexing/usecase/event"
)

type AdditionalExampleProjection struct {
	*rdbprojectionbase.Base

	rdbConn      rdb.Conn
	logger       applogger.Logger
}

func NewAdditionalProjection(
	logger applogger.Logger,
	rdbConn rdb.Conn,
) *AdditionalExampleProjection {
	return &AdditionalExampleProjection{
		rdbprojectionbase.NewRDbBase(rdbConn.ToHandle(), "Account"),
		rdbConn,
		logger,
	}
}

var (
	NewAccountsView              = example_view.NewExamplesView
	UpdateLastHandledEventHeight = (*AdditionalExampleProjection).UpdateLastHandledEventHeight
)

func (_ *AdditionalExampleProjection) GetEventsToListen() []string {
	return []string{
		event_usecase.ACCOUNT_TRANSFERRED,
	}
}

func (projection *AdditionalExampleProjection) OnInit() error {
	return nil
}

func (projection *AdditionalExampleProjection) HandleEvents(height int64, events []event_entity.Event) error {
	rdbTx, err := projection.rdbConn.Begin()
	if err != nil {
		return fmt.Errorf("error beginning transaction: %v", err)
	}

	committed := false
	defer func() {
		if !committed {
			_ = rdbTx.Rollback()
		}
	}()

	rdbTxHandle := rdbTx.ToHandle()

	accountsView := NewAccountsView(rdbTxHandle)

	for _, event := range events {
		if accountCreatedEvent, ok := event.(*event_usecase.AccountTransferred); ok {
			if handleErr := projection.handleSomeEvent(accountsView, accountCreatedEvent); handleErr != nil {
				return fmt.Errorf("error handling AccountCreatedEvent: %v", handleErr)
			}
		}
	}

	if err = UpdateLastHandledEventHeight(projection, rdbTxHandle, height); err != nil {
		return fmt.Errorf("error updating last handled event height: %v", err)
	}

	if err = rdbTx.Commit(); err != nil {
		return fmt.Errorf("error committing changes: %v", err)
	}
	committed = true

	return nil
}

func (projection *AdditionalExampleProjection) handleSomeEvent(accountsView account_view.Accounts, event *event_usecase.AccountTransferred) error {

	recipienterr := projection.writeAccountInfo(accountsView, event.Recipient)
	if recipienterr != nil {
		return recipienterr
	}

	sendererr := projection.writeAccountInfo(accountsView, event.Sender)
	if sendererr != nil {
		return sendererr
	}

	return nil
}
