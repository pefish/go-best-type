package go_best_type

import (
	"sync"

	go_logger "github.com/pefish/go-logger"
)

type ActionType string

const (
	ActionType_Start        ActionType = "start"          // 开始你的工作
	ActionType_ExitAndReply ActionType = "exit_and_reply" // 你回家吧，事情交接下
)

type ExitType string

const (
	ExitType_System ExitType = "system"
	ExitType_User   ExitType = "user"
)

type AskType struct {
	Action     ActionType
	AnswerChan chan interface{}
	Data       interface{}
}

type IBestType interface {
	Start(exitChan <-chan ExitType, ask *AskType)
	ProcessOtherAsk(exitChan <-chan ExitType, ask *AskType)
	Name() string

	// 每个人只能通过 ask 来沟通
	Ask(ask *AskType)
	AskForAnswer(ask *AskType) interface{}

	SetBestTypeManager(btm *BestTypeManager)
	BestTypeManager() *BestTypeManager
}

type BaseBestType struct {
	logger  go_logger.InterfaceLogger
	askChan chan *AskType
	wg      sync.WaitGroup

	btm *BestTypeManager

	exitChans []chan ExitType
}

func NewBaseBestType(
	myself IBestType,
) *BaseBestType {
	b := &BaseBestType{
		logger:    go_logger.Logger.CloneWithPrefix(myself.Name()),
		askChan:   make(chan *AskType),
		exitChans: make([]chan ExitType, 0),
	}

	go func() {
		for ask := range b.askChan {
			switch ask.Action {
			case ActionType_Start:
				exitChan := make(chan ExitType)
				b.exitChans = append(b.exitChans, exitChan)
				b.wg.Add(1)
				go func(ask *AskType) {
					defer b.wg.Done()
					myself.Start(exitChan, ask)
				}(ask)
			case ActionType_ExitAndReply:
				exitType := ask.Data.(ExitType)
				b.exit(exitType)
				ask.AnswerChan <- true
				return
			default:
				exitChan := make(chan ExitType)
				b.exitChans = append(b.exitChans, exitChan)
				b.wg.Add(1)
				go func(ask *AskType) {
					defer b.wg.Done()
					myself.ProcessOtherAsk(exitChan, ask)
				}(ask)
			}
		}
	}()

	return b
}

func (b *BaseBestType) exit(exitType ExitType) {
	for _, exitChan := range b.exitChans {
		exitChan <- exitType
	}
	b.wg.Wait()
}

func (b *BaseBestType) Logger() go_logger.InterfaceLogger {
	return b.logger
}

func (b *BaseBestType) BestTypeManager() *BestTypeManager {
	return b.btm
}

func (b *BaseBestType) SetBestTypeManager(btm *BestTypeManager) {
	b.btm = btm
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
