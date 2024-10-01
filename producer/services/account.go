package services

import (
	"errors"
	"events"
	"log"

	"github.com/google/uuid"
	"github.com/kritAsawaniramol/gokafka/producer/commands"
)

type AccountService interface {
	OpenAccount(command commands.OpenAccountCommand) (id string, err error)
	DepositFund(command commands.DepositFundCommand) error
	WithdrawFund(command commands.WithdrawFundCommand) error
	CloseAccount(command commands.CloseAccountCommand) error
}

type accountService struct {
	eventProducer EventProducer
}

// CloseAccount implements AccountService.
func (a accountService) CloseAccount(command commands.CloseAccountCommand) error {
	if command.ID == "" {
		return errors.New("bad request")
	}

	event := events.CloseAccountEvent{
		ID: command.ID,
	}

	log.Printf("%#v", event)
	return a.eventProducer.Produce(event)
}

// DepositFund implements AccountService.
func (a accountService) DepositFund(command commands.DepositFundCommand) error {
	if command.ID == "" || command.Amount == 0 {
		return errors.New("bad request")
	}

	event := events.DepositFundEvent{
		ID:     command.ID,
		Amount: command.Amount,
	}

	log.Printf("%#v", event)
	return a.eventProducer.Produce(event)
}

// OpenAccount implements AccountService.
func (a accountService) OpenAccount(command commands.OpenAccountCommand) (id string, err error) {
	if command.AccountHolder == "" || command.AccountType == 0 || command.OpeningBalance == 0 {
		return "", errors.New("bad request")
	}

	event := events.OpenAccountEvent{
		ID:             uuid.NewString(),
		AccountHolder:  command.AccountHolder,
		AccountType:    command.AccountType,
		OpeningBalance: command.OpeningBalance,
	}

	log.Printf("%#v", event)
	return event.ID, a.eventProducer.Produce(event)
}

// WithdrawFund implements AccountService.
func (a accountService) WithdrawFund(command commands.WithdrawFundCommand) error {
	if command.ID == "" || command.Amount == 0 {
		return errors.New("bad request")
	}

	event := events.WithdrawFundEvent{
		ID:     command.ID,
		Amount: command.Amount,
	}

	log.Printf("%#v", event)
	return a.eventProducer.Produce(event)
}

func NewAccountService(eventProducer EventProducer) AccountService {
	return accountService{eventProducer: eventProducer}
}
