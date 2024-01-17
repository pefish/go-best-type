package main

import (
	"context"
	"fmt"
	"time"
)

type PersonCType struct {
	BaseBestType
}

func NewPersonC(ctx context.Context) *PersonCType {
	return &PersonCType{
		BaseBestType: BaseBestType{
			ctx:     ctx,
			askChan: make(chan *AskType),
		},
	}
}

func (p *PersonCType) ProcessAsk(ask *AskType, bts map[string]IBestType) {
	switch ask.Action {
	case ActionType_Develop:
		go func() {
			fmt.Printf("【开发工程师】收到开发任务 <%s>，开发中。。。\n", ask.Action)
			time.Sleep(5 * time.Second)
			fmt.Printf("【开发工程师】开发完成。向测试工程师提交测试\n")
			bts["personD"].Ask(&AskType{
				Action: "test",
			})
		}()
	case ActionType_Bug:
		go func() {
			fmt.Printf("【开发工程师】收到 Bug <%s>，修复中。。。\n", ask.Action)
			time.Sleep(5 * time.Second)
			fmt.Printf("【开发工程师】修复完成。向测试工程师提交测试\n")
			bts["personD"].Ask(&AskType{
				Action: "test",
			})
		}()
	}
}

func (p *PersonCType) Exited() {
	fmt.Printf("【开发工程师】下班了\n")
}
