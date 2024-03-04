package main

import (
	"time"

	go_best_type "github.com/pefish/go-best-type"
)

type PersonAType struct {
	go_best_type.BaseBestType
	ceoAnswer chan interface{}
}

func NewPersonA(name string) *PersonAType {
	p := &PersonAType{}
	p.BaseBestType = *go_best_type.NewBaseBestType(p, name)
	return p
}

func (p *PersonAType) Start(exitChan <-chan go_best_type.ExitType, ask *go_best_type.AskType) error {
	return nil
}

func (p *PersonAType) ProcessOtherAsk(exitChan <-chan go_best_type.ExitType, ask *go_best_type.AskType) error {
	switch ask.Action {
	case ActionType_InitNeed:
		// 时间长的工作不能影响耳朵收听，新开协程
		go func() {
			p.Logger().InfoF("收到新需求 <%s>，处理需求中。。。", ask.Action)
			time.Sleep(2 * time.Second)
			p.Logger().InfoF("需求处理完成。画原型图中。。。")
			time.Sleep(2 * time.Second)
			p.Logger().InfoF("原型图完成。向 UI 设计师发送设计任务")
			p.BestTypeManager().Get("UI 设计师").Ask(&go_best_type.AskType{
				Action: "design task",
			})
		}()
		p.ceoAnswer = ask.AnswerChan
	case ActionType_ChangeNeed:
		go func() {
			p.Logger().InfoF("收到需求变更 <%s>，处理需求中。。。", ask.Action)
			time.Sleep(2 * time.Second)
			p.Logger().InfoF("需求处理完成。画原型图中。。。")
			time.Sleep(2 * time.Second)
			p.Logger().InfoF("原型图完成。向 UI 设计师发送设计任务")
			p.BestTypeManager().Get("personB").Ask(&go_best_type.AskType{
				Action: "design task",
			})
		}()
	case ActionType_CheckNotify:
		go func() {
			p.Logger().InfoF("收到产品验收请求，验收产品中。。。")
			time.Sleep(2 * time.Second)
			p.Logger().InfoF("产品验收完成，合格")
			p.ceoAnswer <- "finished"
		}()
	}

	select {
	case <-exitChan:
		p.Logger().InfoF("<%s> 做完了", ask.Action)
	}

	return nil
}
