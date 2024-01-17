package main

import (
	"context"
	"fmt"
	go_best_type "github.com/pefish/go-best-type"
	"time"
)

type PersonDType struct {
	go_best_type.BaseBestType
}

func NewPersonD(ctx context.Context) *PersonDType {
	return &PersonDType{
		BaseBestType: *go_best_type.NewBaseBestType(ctx),
	}
}

func (p *PersonDType) ProcessAsk(ask *go_best_type.AskType, bts map[string]go_best_type.IBestType) {
	switch ask.Action {
	case ActionType_Test:
		go func() {
			fmt.Printf("【测试工程师】收到测试任务，测试中。。。\n")
			time.Sleep(5 * time.Second)
			fmt.Printf("【测试工程师】测试完成。提交产品验收\n")
			bts["personA"].Ask(&go_best_type.AskType{
				Action: "check notify",
			})
		}()
	}
}

func (p *PersonDType) Exited() {
	fmt.Printf("【测试工程师】下班了\n")
}
