package main

import (
	"context"
	"time"

	go_best_type "github.com/pefish/go-best-type"
)

type PersonDType struct {
	go_best_type.BaseBestType
}

func NewPersonD(ctx context.Context, bestTypeManager *go_best_type.BestTypeManager) *PersonDType {
	p := &PersonDType{}
	p.BaseBestType = *go_best_type.NewBaseBestType(ctx, p, bestTypeManager, 0)
	return p
}

func (p *PersonDType) ProcessAsk(ctx context.Context, ask *go_best_type.AskType) {
	switch ask.Action {
	case ActionType_Test:
		go func() {
			p.Logger().InfoF("收到测试任务，测试中。。。")
			time.Sleep(5 * time.Second)
			p.Logger().InfoF("测试完成。提交产品验收")
			p.BestTypeManager().Get("personA").Ask(&go_best_type.AskType{
				Action: "check notify",
			})
		}()
	}
}

func (p *PersonDType) OnExited() {
	p.Logger().InfoF("下班了")
}

func (p *PersonDType) Name() string {
	return "测试工程师"
}
