package main

import (
	"context"
	"fmt"
	"time"
)

type PersonAType struct {
	BaseBestType
}

func NewPersonA(ctx context.Context) *PersonAType {
	return &PersonAType{
		BaseBestType: BaseBestType{
			ctx:     ctx,
			askChan: make(chan *AskType),
		},
	}
}

func (p *PersonAType) ProcessAsk(ask *AskType, bts map[string]IBestType) {
	switch ask.Action {
	case ActionType_InitNeed:
		// 时间长的工作不能影响耳朵收听，新开协程
		go func() {
			fmt.Printf("【产品经理】收到新需求 <%s>，处理需求中。。。\n", ask.Action)
			time.Sleep(5 * time.Second)
			fmt.Printf("【产品经理】需求处理完成。画原型图中。。。\n")
			time.Sleep(5 * time.Second)
			fmt.Printf("【产品经理】原型图完成。向 UI 设计师发送设计任务\n")
			bts["personB"].Ask(&AskType{
				Action: "design task",
			})
		}()
	case ActionType_ChangeNeed:
		go func() {
			fmt.Printf("【产品经理】收到需求变更 <%s>，处理需求中。。。\n", ask.Action)
			time.Sleep(5 * time.Second)
			fmt.Printf("【产品经理】需求处理完成。画原型图中。。。\n")
			time.Sleep(5 * time.Second)
			fmt.Printf("【产品经理】原型图完成。向 UI 设计师发送设计任务\n")
			bts["personB"].Ask(&AskType{
				Action: "design task",
			})
		}()
	case ActionType_CheckNotify:
		go func() {
			fmt.Printf("【产品经理】收到产品验收请求，验收产品中。。。\n")
			time.Sleep(5 * time.Second)
			fmt.Printf("【产品经理】产品验收完成，合格\n")
			bts["personE"].Ask(&AskType{
				Action: "finished",
			})
		}()
	}
}

func (p *PersonAType) Exited() {
	fmt.Printf("【产品经理】下班了\n")
}