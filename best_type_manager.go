package go_best_type

import (
	"sync"
)

type BestTypeManager struct {
	wg         sync.WaitGroup
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

func (b *BestTypeManager) Set(name string, bestType IBestType) {
	b.wg.Add(1)
	b.btsCollect.Store(name, bestType)
}

func (b *BestTypeManager) StopOneAsync(name string) {
	v, ok := b.btsCollect.Load(name)
	if !ok {
		return
	}
	bestType := v.(IBestType)
	bestType.Ask(&AskType{
		Action: ActionType_Stop,
	})
}

func (b *BestTypeManager) StopAllAsync() {
	b.btsCollect.Range(func(key any, value any) bool {
		bestType := value.(IBestType)
		bestType.Ask(&AskType{
			Action: ActionType_Stop,
		})
		return true
	})
}

func (b *BestTypeManager) TerminalOneAsync(name string) {
	v, ok := b.btsCollect.Load(name)
	if !ok {
		return
	}
	bestType := v.(IBestType)
	bestType.Ask(&AskType{
		Action: ActionType_Terminal,
	})
}

func (b *BestTypeManager) TerminalAllAsync() {
	b.btsCollect.Range(func(key any, value any) bool {
		bestType := value.(IBestType)
		bestType.Ask(&AskType{
			Action: ActionType_Terminal,
		})
		return true
	})
}

func (b *BestTypeManager) Wait() {
	b.wg.Wait()
}

func (b *BestTypeManager) WaitGroup() *sync.WaitGroup {
	return &b.wg
}
