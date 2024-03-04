package main

import (
	"time"

	go_best_type "github.com/pefish/go-best-type"
)

type PersonCType struct {
	go_best_type.BaseBestType
}

func NewPersonC(name string) *PersonCType {
	p := &PersonCType{}
	p.BaseBestType = *go_best_type.NewBaseBestType(p, name)
	return p
}

func (p *PersonCType) Start(exitChan <-chan go_best_type.ExitType, ask *go_best_type.AskType) error {
	return nil
}

func (p *PersonCType) ProcessOtherAsk(exitChan <-chan go_best_type.ExitType, ask *go_best_type.AskType) error {
	switch ask.Action {
	case ActionType_Develop:
		go func() {
			p.Logger().InfoF("收到开发任务 <%s>，开发中。。。", ask.Action)
			time.Sleep(2 * time.Second)
			p.Logger().InfoF("开发完成。向测试工程师提交测试")
			p.BestTypeManager().Get("测试工程师").Ask(&go_best_type.AskType{
				Action: "test",
			})
		}()
	case ActionType_Bug:
		go func() {
			p.Logger().InfoF("收到 Bug <%s>，修复中。。。", ask.Action)
			time.Sleep(2 * time.Second)
			p.Logger().InfoF("修复完成。向测试工程师提交测试")
			p.BestTypeManager().Get("测试工程师").Ask(&go_best_type.AskType{
				Action: "test",
			})
		}()
	}

	select {
	case <-exitChan:
		p.Logger().InfoF("<%s> 做完了", ask.Action)
	}

	return nil
}
