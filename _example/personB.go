package main

import (
	"context"
	"fmt"
	go_best_type "github.com/pefish/go-best-type"
	"time"
)

type PersonBType struct {
	go_best_type.BaseBestType
}

func NewPersonB(ctx context.Context) *PersonBType {
	return &PersonBType{
		BaseBestType: *go_best_type.NewBaseBestType(ctx, 0),
	}
}

func (p *PersonBType) ProcessAsk(ask *go_best_type.AskType, bts map[string]go_best_type.IBestType) {
	switch ask.Action {
	case ActionType_DesignTask:
		go func() {
			fmt.Printf("【UI 设计师】收到设计任务 <%s>，设计中。。。\n", ask.Action)
			time.Sleep(5 * time.Second)
			fmt.Printf("【UI 设计师】设计完成。向开发工程师发送开发任务\n")
			bts["personC"].Ask(&go_best_type.AskType{
				Action: "develop",
			})
		}()
	case ActionType_DesignChange:
		go func() {
			fmt.Printf("【UI 设计师】收到设计变更任务 <%s>，设计中。。。\n", ask.Action)
			time.Sleep(5 * time.Second)
			fmt.Printf("【UI 设计师】设计完成。向开发工程师发送开发任务\n")
			bts["personC"].Ask(&go_best_type.AskType{
				Action: "develop",
			})
		}()
	}
}

func (p *PersonBType) OnExited() {
	fmt.Printf("【UI 设计师】下班了\n")
}
