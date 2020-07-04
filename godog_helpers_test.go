package kubernetes_ctx_test

import (
	"github.com/cucumber/godog"
)

// MockScenarioContext returns a mocked ScenarioContext.
func MockScenarioContext() *scenarioContextMock {
	return &scenarioContextMock{}
}

type (
	scenarioContextMock struct {
		beforeScenarioList []func(sc *godog.Scenario)
		afterScenarioList  []func(sc *godog.Scenario, err error)
		beforeStepList     []func(sc *godog.Step)
		afterStepList      []func(sc *godog.Step, err error)
		stepList           []stepMock
	}
	stepMock struct {
		expr interface{}
		fnc  interface{}
	}
)

func (s *scenarioContextMock) BeforeScenario(fn func(sc *godog.Scenario)) {
	s.beforeScenarioList = append(s.beforeScenarioList, fn)
}

func (s *scenarioContextMock) AfterScenario(fn func(sc *godog.Scenario, err error)) {
	s.afterScenarioList = append(s.afterScenarioList, fn)
}

func (s *scenarioContextMock) BeforeStep(fn func(st *godog.Step)) {
	s.beforeStepList = append(s.beforeStepList, fn)
}

func (s *scenarioContextMock) AfterStep(fn func(st *godog.Step, err error)) {
	s.afterStepList = append(s.afterStepList, fn)
}

func (s *scenarioContextMock) Step(expr, stepFunc interface{}) {
	s.stepList = append(s.stepList, stepMock{expr: expr, fnc: stepFunc})
}

func (s *scenarioContextMock) RunScenario() {
	for _, fn := range s.beforeScenarioList {
		fn(nil)
	}
}