package main

import (
	"context"
	"fmt"
	"time"
)

type PersonDType struct {
	BaseBestType
}

func NewPersonD(ctx context.Context) *PersonDType {
	return &PersonDType{
		BaseBestType: BaseBestType{
			ctx:     ctx,
			askChan: make(chan *AskType),
		},
	}
}

func (p *PersonDType) ProcessAsk(ask *AskType, bts map[string]IBestType) {
	switch ask.Action {
	case ActionType_Test:
		go func() {
			fmt.Printf("【测试工程师】收到测试任务，测试中。。。\n")
			time.Sleep(5 * time.Second)
			fmt.Printf("【测试工程师】测试完成。提交产品验收\n")
			bts["personA"].Ask(&AskType{
				Action: "check notify",
			})
		}()
	}
}

func (p *PersonDType) Exited() {
	fmt.Printf("【测试工程师】下班了\n")
}
