package main

import (
	"time"

	go_best_type "github.com/pefish/go-best-type"
)

type PersonBType struct {
	go_best_type.BaseBestType
}

func NewPersonB() *PersonBType {
	p := &PersonBType{}
	p.BaseBestType = *go_best_type.NewBaseBestType(p)
	return p
}

func (p *PersonBType) Start(exitChan <-chan go_best_type.ExitType, ask *go_best_type.AskType) {
}

func (p *PersonBType) ProcessOtherAsk(exitChan <-chan go_best_type.ExitType, ask *go_best_type.AskType) {
	switch ask.Action {
	case ActionType_DesignTask:
		go func() {
			p.Logger().InfoF("收到设计任务 <%s>，设计中。。。", ask.Action)
			time.Sleep(2 * time.Second)
			p.Logger().InfoF("设计完成。向开发工程师发送开发任务")
			p.BestTypeManager().Get("开发工程师").Ask(&go_best_type.AskType{
				Action: "develop",
			})
		}()
	case ActionType_DesignChange:
		go func() {
			p.Logger().InfoF("收到设计变更任务 <%s>，设计中。。。", ask.Action)
			time.Sleep(2 * time.Second)
			p.Logger().InfoF("设计完成。向开发工程师发送开发任务")
			p.BestTypeManager().Get("开发工程师").Ask(&go_best_type.AskType{
				Action: "develop",
			})
		}()
	}

	select {
	case <-exitChan:
		p.Logger().InfoF("<%s> 做完了", ask.Action)
	}
}

func (p *PersonBType) Name() string {
	return "UI 设计师"
}
