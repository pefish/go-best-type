package main

import (
	"context"
	"fmt"
	go_best_type "github.com/pefish/go-best-type"
	"time"
)

type PersonCType struct {
	go_best_type.BaseBestType
}

func NewPersonC(ctx context.Context) *PersonCType {
	return &PersonCType{
		BaseBestType: *go_best_type.NewBaseBestType(ctx),
	}
}

func (p *PersonCType) ProcessAsk(ask *go_best_type.AskType, bts map[string]go_best_type.IBestType) {
	switch ask.Action {
	case ActionType_Develop:
		go func() {
			fmt.Printf("【开发工程师】收到开发任务 <%s>，开发中。。。\n", ask.Action)
			time.Sleep(5 * time.Second)
			fmt.Printf("【开发工程师】开发完成。向测试工程师提交测试\n")
			bts["personD"].Ask(&go_best_type.AskType{
				Action: "test",
			})
		}()
	case ActionType_Bug:
		go func() {
			fmt.Printf("【开发工程师】收到 Bug <%s>，修复中。。。\n", ask.Action)
			time.Sleep(5 * time.Second)
			fmt.Printf("【开发工程师】修复完成。向测试工程师提交测试\n")
			bts["personD"].Ask(&go_best_type.AskType{
				Action: "test",
			})
		}()
	}
}

func (p *PersonCType) Exited() {
	fmt.Printf("【开发工程师】下班了\n")
}
