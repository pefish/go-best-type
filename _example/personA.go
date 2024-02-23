package main

import (
	"context"
	"time"

	go_best_type "github.com/pefish/go-best-type"
)

type PersonAType struct {
	go_best_type.BaseBestType
}

func NewPersonA(ctx context.Context) *PersonAType {
	p := &PersonAType{}
	p.BaseBestType = *go_best_type.NewBaseBestType(ctx, p, 0)
	return p
}

func (p *PersonAType) ProcessAsk(ask *go_best_type.AskType, bts map[string]go_best_type.IBestType) {
	switch ask.Action {
	case ActionType_InitNeed:
		// 时间长的工作不能影响耳朵收听，新开协程
		go func() {
			p.Logger().InfoF("收到新需求 <%s>，处理需求中。。。\n", ask.Action)
			time.Sleep(5 * time.Second)
			p.Logger().InfoF("需求处理完成。画原型图中。。。\n")
			time.Sleep(5 * time.Second)
			p.Logger().InfoF("原型图完成。向 UI 设计师发送设计任务\n")
			bts["personB"].Ask(&go_best_type.AskType{
				Action: "design task",
			})
		}()
	case ActionType_ChangeNeed:
		go func() {
			p.Logger().InfoF("收到需求变更 <%s>，处理需求中。。。\n", ask.Action)
			time.Sleep(5 * time.Second)
			p.Logger().InfoF("需求处理完成。画原型图中。。。\n")
			time.Sleep(5 * time.Second)
			p.Logger().InfoF("原型图完成。向 UI 设计师发送设计任务\n")
			bts["personB"].Ask(&go_best_type.AskType{
				Action: "design task",
			})
		}()
	case ActionType_CheckNotify:
		go func() {
			p.Logger().InfoF("收到产品验收请求，验收产品中。。。\n")
			time.Sleep(5 * time.Second)
			p.Logger().InfoF("产品验收完成，合格\n")
			bts["personE"].Ask(&go_best_type.AskType{
				Action: "finished",
			})
		}()
	}
}

func (p *PersonAType) OnExited() {
	p.Logger().InfoF("下班了\n")
}

func (p *PersonAType) Name() string {
	return "产品经理"
}
