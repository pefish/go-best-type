package main

import (
	"context"
	"time"

	go_best_type "github.com/pefish/go-best-type"
)

type PersonBType struct {
	go_best_type.BaseBestType
}

func NewPersonB(ctx context.Context) *PersonBType {
	p := &PersonBType{}
	p.BaseBestType = *go_best_type.NewBaseBestType(ctx, p, 0)
	return p
}

func (p *PersonBType) ProcessAsk(ask *go_best_type.AskType, bts map[string]go_best_type.IBestType) {
	switch ask.Action {
	case ActionType_DesignTask:
		go func() {
			p.Logger().InfoF("收到设计任务 <%s>，设计中。。。\n", ask.Action)
			time.Sleep(5 * time.Second)
			p.Logger().InfoF("设计完成。向开发工程师发送开发任务\n")
			bts["personC"].Ask(&go_best_type.AskType{
				Action: "develop",
			})
		}()
	case ActionType_DesignChange:
		go func() {
			p.Logger().InfoF("收到设计变更任务 <%s>，设计中。。。\n", ask.Action)
			time.Sleep(5 * time.Second)
			p.Logger().InfoF("设计完成。向开发工程师发送开发任务\n")
			bts["personC"].Ask(&go_best_type.AskType{
				Action: "develop",
			})
		}()
	}
}

func (p *PersonBType) OnExited() {
	p.Logger().InfoF("下班了\n")
}

func (p *PersonBType) Name() string {
	return "UI 设计师"
}
