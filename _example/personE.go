package main

import (
	"context"
	"fmt"
)

type PersonEType struct {
	BaseBestType

	cancelFunc context.CancelFunc
}

func NewPersonE(ctx context.Context, cancelFunc context.CancelFunc) *PersonEType {
	return &PersonEType{
		BaseBestType: BaseBestType{
			ctx:     ctx,
			askChan: make(chan *AskType),
		},
		cancelFunc: cancelFunc,
	}
}

func (p *PersonEType) ProcessAsk(ask *AskType, bts map[string]IBestType) {
	switch ask.Action {
	case ActionType_Finished:
		fmt.Printf("【CEO】产品开发完成，恭喜各位，可以休息了！！！\n")
		p.cancelFunc()
		return
	}
}

func (p *PersonEType) Start(personA *PersonAType) {
	personA.Ask(&AskType{
		Action: ActionType_InitNeed,
	})
}

func (p *PersonEType) Exited() {
	fmt.Printf("CEO 下班了\n")
}
