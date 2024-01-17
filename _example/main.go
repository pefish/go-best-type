package main

import (
	"context"
	go_best_type "github.com/pefish/go-best-type"
	"sync"
)

const (
	ActionType_InitNeed     go_best_type.ActionType = "init need"
	ActionType_ChangeNeed   go_best_type.ActionType = "change need"
	ActionType_CheckNotify  go_best_type.ActionType = "check notify"
	ActionType_DesignTask   go_best_type.ActionType = "design task"
	ActionType_DesignChange go_best_type.ActionType = "design change"
	ActionType_Develop      go_best_type.ActionType = "develop"
	ActionType_Bug          go_best_type.ActionType = "bug"
	ActionType_Test         go_best_type.ActionType = "test"
	ActionType_Finished     go_best_type.ActionType = "finished"
)

func main() {
	var wg sync.WaitGroup

	// 每个人非常积极负责任的传达消息，不存在需要催的情况，这样的效率才是最高的
	// 假设一个产品开发的场景（A：产品经理，B：UI设计师，C：开发工程师，D：测试工程师，E：CEO）
	// A：接收处理需求、接收产品验收通知、接收需求的变更
	// B：接收设计任务、接收风格的变更
	// C：接收开发任务、接收 Bug
	// D：接收测试任务、接收 Bug 验收

	adminCtx, cancel := context.WithCancel(context.Background())

	personA := NewPersonA(adminCtx)
	personB := NewPersonB(adminCtx)
	personC := NewPersonC(adminCtx)
	personD := NewPersonD(adminCtx)
	personE := NewPersonE(adminCtx, cancel)

	wg.Add(5)
	go func() {
		defer wg.Done()
		personA.Listen(personA, map[string]go_best_type.IBestType{
			"personB": personB,
			"personE": personE,
		})
	}()
	go func() {
		defer wg.Done()
		personB.Listen(personB, map[string]go_best_type.IBestType{
			"personC": personC,
		})
	}()
	go func() {
		defer wg.Done()
		personC.Listen(personC, map[string]go_best_type.IBestType{
			"personD": personD,
		})
	}()
	go func() {
		defer wg.Done()
		personD.Listen(personD, map[string]go_best_type.IBestType{
			"personA": personA,
		})
	}()

	// CEO 提出需求
	personE.Start(personA)
	personE.Listen(personE, nil)
	wg.Done()

	wg.Wait()
}
