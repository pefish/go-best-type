package main

import (
	"context"
)

type ActionType string

type AskType struct {
	Action ActionType
	Data   interface{}
}

type IBestType interface {
	Ask(ask *AskType)
	Listen(myself IBestType, bts map[string]IBestType)
	ProcessAsk(ask *AskType, bts map[string]IBestType)
	Exited()
}

type BaseBestType struct {
	ctx     context.Context
	askChan chan *AskType
}

func (b *BaseBestType) Ask(ask *AskType) {
	b.askChan <- ask
}

func (b *BaseBestType) Listen(
	myself IBestType,
	bts map[string]IBestType,
) {
	for {
		select {
		case ask := <-b.askChan:
			myself.ProcessAsk(ask, bts)
		case <-b.ctx.Done():
			myself.Exited()
			return
		}
	}
}
