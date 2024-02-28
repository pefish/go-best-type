package go_best_type

import (
	"context"

	go_logger "github.com/pefish/go-logger"
)

type ActionType string

const (
	ActionType_Start    ActionType = "start"
	ActionType_Stop     ActionType = "stop"
	ActionType_Terminal ActionType = "terminal"
)

type AskType struct {
	Action     ActionType
	AnswerChan chan interface{}
	Data       interface{}
}

type IBestType interface {
	Start(stopCtx context.Context, terminalCtx context.Context, ask *AskType)
	ProcessOtherAsk(stopCtx context.Context, terminalCtx context.Context, ask *AskType)
	Name() string

	Ask(ask *AskType)
	AskForAnswer(ask *AskType) interface{}
	BestTypeManager() *BestTypeManager
}

type BaseBestType struct {
	logger          go_logger.InterfaceLogger
	askChan         chan *AskType
	bestTypeManager *BestTypeManager

	stopCancel     context.CancelFunc
	terminalCancel context.CancelFunc
}

func NewBaseBestType(
	myself IBestType,
	bestTypeManager *BestTypeManager,
	askChanCap int,
) *BaseBestType {
	stopCtx, stopCancel := context.WithCancel(context.Background())
	terminalCtx, terminalCancel := context.WithCancel(context.Background())

	b := &BaseBestType{
		logger:          go_logger.Logger.CloneWithPrefix(myself.Name()),
		askChan:         make(chan *AskType, askChanCap),
		bestTypeManager: bestTypeManager,
		stopCancel:      stopCancel,
		terminalCancel:  terminalCancel,
	}

	go func() {
		for ask := range b.askChan {
			switch ask.Action {
			case ActionType_Start:
				bestTypeManager.WaitGroup().Add(1)
				go func(ask *AskType) {
					defer bestTypeManager.WaitGroup().Done()
					myself.Start(stopCtx, terminalCtx, ask)
				}(ask)
			case ActionType_Stop:
				stopCancel()
				return
			case ActionType_Terminal:
				terminalCancel()
				return
			default:
				bestTypeManager.WaitGroup().Add(1)
				go func(ask *AskType) {
					defer bestTypeManager.WaitGroup().Done()
					myself.ProcessOtherAsk(stopCtx, terminalCtx, ask)
				}(ask)
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

func (b *BaseBestType) BestTypeManager() *BestTypeManager {
	return b.bestTypeManager
}
