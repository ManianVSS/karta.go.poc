package xlang

type StepResult struct {
	name       string
	successful bool
	err        error
	result     any
}
