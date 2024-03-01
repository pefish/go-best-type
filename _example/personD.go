package main

import (
	"time"

	go_best_type "github.com/pefish/go-best-type"
)

type PersonDType struct {
	go_best_type.BaseBestType
}

func NewPersonD(name string) *PersonDType {
	p := &PersonDType{}
	p.BaseBestType = *go_best_type.NewBaseBestType(p, name)
	return p
}

func (p *PersonDType) Start(exitChan <-chan go_best_type.ExitType, ask *go_best_type.AskType) {
}

func (p *PersonDType) ProcessOtherAsk(exitChan <-chan go_best_type.ExitType, ask *go_best_type.AskType) {
	switch ask.Action {
	case ActionType_Test:
		go func() {
			p.Logger().InfoF("收到测试任务，测试中。。。")
			time.Sleep(2 * time.Second)
			p.Logger().InfoF("测试完成。提交产品验收")
			p.BestTypeManager().Get("产品经理").Ask(&go_best_type.AskType{
				Action: "check notify",
			})
		}()
	}

	select {
	case <-exitChan:
		p.Logger().InfoF("<%s> 做完了", ask.Action)
	}
}
