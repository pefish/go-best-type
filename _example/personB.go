package main

import (
	"time"

	go_best_type "github.com/pefish/go-best-type"
)

type PersonBType struct {
	go_best_type.BaseBestType
}

func NewPersonB(bestTypeManager *go_best_type.BestTypeManager) *PersonBType {
	p := &PersonBType{}
	p.BaseBestType = *go_best_type.NewBaseBestType(p, bestTypeManager, 0)
	return p
}

func (p *PersonBType) Start(ask *go_best_type.AskType) {}

func (p *PersonBType) Stop(ask *go_best_type.AskType) {
	p.Logger().InfoF("下班了")
}

func (p *PersonBType) Terminal(ask *go_best_type.AskType) {}

func (p *PersonBType) ProcessOtherAsk(ask *go_best_type.AskType) {
	switch ask.Action {
	case ActionType_DesignTask:
		go func() {
			p.Logger().InfoF("收到设计任务 <%s>，设计中。。。", ask.Action)
			time.Sleep(5 * time.Second)
			p.Logger().InfoF("设计完成。向开发工程师发送开发任务")
			p.BestTypeManager().Get("personC").Ask(&go_best_type.AskType{
				Action: "develop",
			})
		}()
	case ActionType_DesignChange:
		go func() {
			p.Logger().InfoF("收到设计变更任务 <%s>，设计中。。。", ask.Action)
			time.Sleep(5 * time.Second)
			p.Logger().InfoF("设计完成。向开发工程师发送开发任务")
			p.BestTypeManager().Get("personC").Ask(&go_best_type.AskType{
				Action: "develop",
			})
		}()
	}
}

func (p *PersonBType) Name() string {
	return "UI 设计师"
}
