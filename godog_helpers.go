package kubernetes_ctx

import "github.com/cucumber/godog"

// ScenarioContext wraps godog.ScenarioContext in order to easily
// mock it for internal tests.
type ScenarioContext interface {
	BeforeScenario(fn func(sc *godog.Scenario))
	AfterScenario(fn func(sc *godog.Scenario, err error))
	BeforeStep(fn func(st *godog.Step))
	AfterStep(fn func(st *godog.Step, err error))
	Step(expr, stepFunc interface{})
}
