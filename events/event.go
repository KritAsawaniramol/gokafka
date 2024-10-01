package events

import "reflect"

// reflect.TypeOf(OpenAccountEvent{}).Name() return name of type/struct "OpenAccountEvent"
var Topics = []string{
	reflect.TypeOf(OpenAccountEvent{}).Name(),
	reflect.TypeOf(DepositFundEvent{}).Name(),
	reflect.TypeOf(WithdrawFundEvent{}).Name(),
	reflect.TypeOf(CloseAccountEvent{}).Name(),
}

type Event interface {
}

type OpenAccountEvent struct {
	ID             string
	AccountHolder  string
	AccountType    int
	OpeningBalance float64
}

type DepositFundEvent struct {
	ID     string
	Amount float64
}

type WithdrawFundEvent struct {
	ID     string
	Amount float64
}

type CloseAccountEvent struct {
	ID string
}
