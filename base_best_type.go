package go_best_type

import (
	"context"

	go_logger "github.com/pefish/go-logger"
)

type ActionType string

type AskType struct {
	Action     ActionType
	AnswerChan chan interface{}
	Data       interface{}
}

type IBestType interface {
	AskChan() chan *AskType
	Ask(ask *AskType)
	AskForAnswer(ask *AskType) interface{}
	ProcessAsk(ctx context.Context, ask *AskType)
	Name() string
	BestTypeManager() *BestTypeManager
	OnExited()
}

type BaseBestType struct {
	ctx             context.Context
	logger          go_logger.InterfaceLogger
	askChan         chan *AskType
	bestTypeManager *BestTypeManager
}

func NewBaseBestType(
	ctx context.Context,
	myself IBestType,
	bestTypeManager *BestTypeManager,
	askChanCap int,
) *BaseBestType {
	return &BaseBestType{
		ctx:             ctx,
		logger:          go_logger.Logger.CloneWithPrefix(myself.Name()),
		askChan:         make(chan *AskType, askChanCap),
		bestTypeManager: bestTypeManager,
	}
}

func (b *BaseBestType) AskChan() chan *AskType {
	return b.askChan
}

func (b *BaseBestType) Logger() go_logger.InterfaceLogger {
	return b.logger
}

func (b *BaseBestType) Ask(ask *AskType) {
	b.askChan <- ask
}

func (b *BaseBestType) AskForAnswer(ask *AskType) interface{} {
	if ask.AnswerChan == nil {
		ask.AnswerChan = make(chan interface{})
	}
	b.askChan <- ask
	return <-ask.AnswerChan
}

func (b *BaseBestType) BestTypeManager() *BestTypeManager {
	return b.bestTypeManager
}
