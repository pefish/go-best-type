package go_best_type

import (
	"sync"
)

type BestTypeManager struct {
	btsCollect sync.Map
}

func NewBestTypeManager() *BestTypeManager {
	return &BestTypeManager{}
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
	bestType.AskForAnswer(&AskType{
		Action: ActionType_ExitAndReply,
		Data:   exitType,
	})
	b.btsCollect.Delete(name)
}

func (b *BestTypeManager) ExitAll(exitType ExitType) {
	b.btsCollect.Range(func(key any, value any) bool {
		bestType := value.(IBestType)
		bestType.AskForAnswer(&AskType{
			Action: ActionType_ExitAndReply,
			Data:   exitType,
		})
		b.btsCollect.Delete(bestType.Name())
		return true
	})
}
