package go_best_type

import (
	"context"
	"sync"
)

type BestTypeManager struct {
	ctx        context.Context
	wg         sync.WaitGroup
	btsCollect sync.Map
}

func NewBestTypeManager(
	ctx context.Context,
) *BestTypeManager {
	return &BestTypeManager{
		ctx: ctx,
	}
}

func (b *BestTypeManager) Get(name string) IBestType {
	v, ok := b.btsCollect.Load(name)
	if !ok {
		return nil
	}
	return v.(IBestType)
}

func (b *BestTypeManager) Set(name string, bestType IBestType) {
	b.btsCollect.Store(name, bestType)
	b.wg.Add(1)
	go func() {
		defer b.wg.Done()
		for {
			select {
			case ask := <-bestType.AskChan():
				bestType.ProcessAsk(b.ctx, ask)
			case <-b.ctx.Done():
				bestType.OnExited()
				return
			}
		}
	}()
}

func (b *BestTypeManager) Wait() {
	b.wg.Wait()
}
