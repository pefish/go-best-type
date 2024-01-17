package main

import (
	"context"
	"fmt"
	go_best_type "github.com/pefish/go-best-type"
)

type PersonEType struct {
	go_best_type.BaseBestType

	cancelFunc context.CancelFunc
}

func NewPersonE(ctx context.Context, cancelFunc context.CancelFunc) *PersonEType {
	return &PersonEType{
		BaseBestType: *go_best_type.NewBaseBestType(ctx, 0),
		cancelFunc:   cancelFunc,
	}
}

func (p *PersonEType) ProcessAsk(ask *go_best_type.AskType, bts map[string]go_best_type.IBestType) {
	switch ask.Action {
	case ActionType_Finished:
		fmt.Printf("【CEO】产品开发完成，恭喜各位，可以休息了！！！\n")
		p.cancelFunc()
		return
	}
}

func (p *PersonEType) Start(personA *PersonAType) {
	personA.Ask(&go_best_type.AskType{
		Action: ActionType_InitNeed,
	})
}

func (p *PersonEType) Exited() {
	fmt.Printf("CEO 下班了\n")
}
