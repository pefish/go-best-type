package main

import (
	"context"
	"time"

	go_best_type "github.com/pefish/go-best-type"
)

type PersonCType struct {
	go_best_type.BaseBestType
}

func NewPersonC(ctx context.Context, bts map[string]go_best_type.IBestType) *PersonCType {
	p := &PersonCType{}
	p.BaseBestType = *go_best_type.NewBaseBestType(ctx, p, bts, 0)
	return p
}

func (p *PersonCType) ProcessAsk(ask *go_best_type.AskType) {
	switch ask.Action {
	case ActionType_Develop:
		go func() {
			p.Logger().InfoF("收到开发任务 <%s>，开发中。。。", ask.Action)
			time.Sleep(5 * time.Second)
			p.Logger().InfoF("开发完成。向测试工程师提交测试")
			p.BtsCollect()["personD"].Ask(&go_best_type.AskType{
				Action: "test",
			})
		}()
	case ActionType_Bug:
		go func() {
			p.Logger().InfoF("收到 Bug <%s>，修复中。。。", ask.Action)
			time.Sleep(5 * time.Second)
			p.Logger().InfoF("修复完成。向测试工程师提交测试")
			p.BtsCollect()["personD"].Ask(&go_best_type.AskType{
				Action: "test",
			})
		}()
	}
}

func (p *PersonCType) OnExited() {
	p.Logger().InfoF("下班了")
}

func (p *PersonCType) Name() string {
	return "开发工程师"
}
