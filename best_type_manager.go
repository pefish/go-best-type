package go_best_type

import (
	"sync"

	go_logger "github.com/pefish/go-logger"
)

type BestTypeManager struct {
	logger     go_logger.InterfaceLogger
	btsCollect sync.Map
}

func NewBestTypeManager(logger go_logger.InterfaceLogger) *BestTypeManager {
	return &BestTypeManager{
		logger: logger,
	}
}

func (b *BestTypeManager) Get(name string) IBestType {
	v, ok := b.btsCollect.Load(name)
	if !ok {
		return nil
	}
	return v.(IBestType)
}

func (b *BestTypeManager) Set(bestType IBestType) {
	bestType.SetBestTypeManager(b)
	b.btsCollect.Store(bestType.Name(), bestType)
}

func (b *BestTypeManager) ExitOne(name string, exitType ExitType) {
	v, ok := b.btsCollect.Load(name)
	if !ok {
		return
	}
	bestType := v.(IBestType)
	b.logger.DebugF("Ask <%s> <%s> for anwser.", name, ActionType_ExitAndReply)
	bestType.AskForAnswer(&AskType{
		Action: ActionType_ExitAndReply,
		Data:   exitType,
	})
	b.logger.DebugF("Ask <%s> <%s> for anwser done.", name, ActionType_ExitAndReply)
	b.btsCollect.Delete(name)
}

func (b *BestTypeManager) ExitAll(exitType ExitType) {
	b.btsCollect.Range(func(key any, value any) bool {
		bestType := value.(IBestType)
		b.logger.DebugF("Ask <%s> <%s> for anwser.", bestType.Name(), ActionType_ExitAndReply)
		bestType.AskForAnswer(&AskType{
			Action: ActionType_ExitAndReply,
			Data:   exitType,
		})
		b.logger.DebugF("Ask <%s> <%s> for anwser done.", bestType.Name(), ActionType_ExitAndReply)
		b.btsCollect.Delete(bestType.Name())
		return true
	})
}
