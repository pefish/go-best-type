package go_best_type

import (
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
	Start(ask *AskType)
	ProcessOtherAsk(ask *AskType)
	Stop(ask *AskType)
	Terminal(ask *AskType)
	Name() string

	Ask(ask *AskType)
	AskForAnswer(ask *AskType) interface{}
	BestTypeManager() *BestTypeManager
}

type BaseBestType struct {
	logger          go_logger.InterfaceLogger
	askChan         chan *AskType
	bestTypeManager *BestTypeManager
}

func NewBaseBestType(
	myself IBestType,
	bestTypeManager *BestTypeManager,
	askChanCap int,
) *BaseBestType {
	b := &BaseBestType{
		logger:          go_logger.Logger.CloneWithPrefix(myself.Name()),
		askChan:         make(chan *AskType, askChanCap),
		bestTypeManager: bestTypeManager,
	}

	go func() {
		for ask := range b.askChan {
			switch ask.Action {
			case ActionType_Start:
				go myself.Start(ask)
			case ActionType_Stop:
				myself.Stop(ask)
				bestTypeManager.WaitGroup().Done()
				return
			case ActionType_Terminal:
				myself.Terminal(ask)
				bestTypeManager.WaitGroup().Done()
				return
			default:
				go myself.ProcessOtherAsk(ask)
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
