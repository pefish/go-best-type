package main

import (
	"context"
	"time"

	go_best_type "github.com/pefish/go-best-type"
)

type PersonDType struct {
	go_best_type.BaseBestType
}

func NewPersonD(ctx context.Context) *PersonDType {
	p := &PersonDType{}
	p.BaseBestType = *go_best_type.NewBaseBestType(ctx, p, 0)
	return p
}

func (p *PersonDType) ProcessAsk(ask *go_best_type.AskType, bts map[string]go_best_type.IBestType) {
	switch ask.Action {
	case ActionType_Test:
		go func() {
			p.Logger().InfoF("收到测试任务，测试中。。。\n")
			time.Sleep(5 * time.Second)
			p.Logger().InfoF("测试完成。提交产品验收\n")
			bts["personA"].Ask(&go_best_type.AskType{
				Action: "check notify",
			})
		}()
	}
}

func (p *PersonDType) OnExited() {
	p.Logger().InfoF("下班了\n")
}

func (p *PersonDType) Name() string {
	return "测试工程师"
}
