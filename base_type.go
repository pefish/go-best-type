package go_best_type

import (
	"context"

	go_logger "github.com/pefish/go-logger"
)

type ActionType string

type AskType struct {
	Action     ActionType
	AnswerChan chan<- interface{}
	Data       interface{}
}

type IBestType interface {
	Ask(ask *AskType)
	Listen(myself IBestType, bts map[string]IBestType)
	ProcessAsk(ask *AskType, bts map[string]IBestType)
	Name() string
	OnExited()
}

type BaseBestType struct {
	ctx     context.Context
	logger  go_logger.InterfaceLogger
	askChan chan *AskType
}

func NewBaseBestType(
	ctx context.Context,
	myself IBestType,
	askChanCap int,
) *BaseBestType {
	return &BaseBestType{
		ctx:     ctx,
		logger:  go_logger.Logger.CloneWithPrefix(myself.Name()),
		askChan: make(chan *AskType, askChanCap),
	}
}

func (b *BaseBestType) Logger() go_logger.InterfaceLogger {
	return b.logger
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
			myself.OnExited()
			return
		}
	}
}
