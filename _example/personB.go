package main

import (
	"context"
	"fmt"
	"time"
)

type PersonBType struct {
	BaseBestType
}

func NewPersonB(ctx context.Context) *PersonBType {
	return &PersonBType{
		BaseBestType: BaseBestType{
			ctx:     ctx,
			askChan: make(chan *AskType),
		},
	}
}

func (p *PersonBType) ProcessAsk(ask *AskType, bts map[string]IBestType) {
	switch ask.Action {
	case ActionType_DesignTask:
		go func() {
			fmt.Printf("【UI 设计师】收到设计任务 <%s>，设计中。。。\n", ask.Action)
			time.Sleep(5 * time.Second)
			fmt.Printf("【UI 设计师】设计完成。向开发工程师发送开发任务\n")
			bts["personC"].Ask(&AskType{
				Action: "develop",
			})
		}()
	case ActionType_DesignChange:
		go func() {
			fmt.Printf("【UI 设计师】收到设计变更任务 <%s>，设计中。。。\n", ask.Action)
			time.Sleep(5 * time.Second)
			fmt.Printf("【UI 设计师】设计完成。向开发工程师发送开发任务\n")
			bts["personC"].Ask(&AskType{
				Action: "develop",
			})
		}()
	}
}

func (p *PersonBType) Exited() {
	fmt.Printf("【UI 设计师】下班了\n")
}
