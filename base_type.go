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
	Ask(ask *AskType)
	AskForAnswer(ask *AskType) interface{}
	ProcessAsk(ctx context.Context, ask *AskType)
	Name() string
	BtsCollect() map[string]IBestType
	OnExited()
}

type BaseBestType struct {
	ctx        context.Context
	logger     go_logger.InterfaceLogger
	askChan    chan *AskType
	btsCollect map[string]IBestType
}

func NewBaseBestType(
	ctx context.Context,
	myself IBestType,
	btsCollect map[string]IBestType,
	askChanCap int,
) *BaseBestType {
	b := &BaseBestType{
		ctx:        ctx,
		logger:     go_logger.Logger.CloneWithPrefix(myself.Name()),
		askChan:    make(chan *AskType, askChanCap),
		btsCollect: btsCollect,
	}
	go func() {
		for {
			select {
			case ask := <-b.askChan:
				myself.ProcessAsk(ctx, ask)
			case <-b.ctx.Done():
				myself.OnExited()
				return
			}
		}
	}()
	return b
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

func (b *BaseBestType) BtsCollect() map[string]IBestType {
	return b.btsCollect
}
