package xlang

import "fmt"

func init() {
	stepDefMap["for"] = createForStatementStep
	stepDefMap["init"] = createLoopInitStatementStep
	stepDefMap["update"] = createLoopUpdateStatementStep
	stepDefMap["do"] = createLoopDoStatementStep
}

type ForStatement struct {
	BaseStep
	forInitStatement   *LoopInitStatement
	conditionStep      Step
	forUpdateStatement *LoopUpdateStatement
	forDoStatement     *LoopDoStatement
}

type LoopInitStatement struct {
	BaseStep
}

type LoopUpdateStatement struct {
	BaseStep
}

type LoopDoStatement struct {
	BaseStep
}

func (forStatement *ForStatement) InitalizeAndCheck() error {

	if _, err := checkAtleastNStep(forStatement.BaseStep, 4); err != nil {
		return err
	}

	initBlock := forStatement.steps[0]

	if initBlock.Name() != "init" {
		return fmt.Errorf("%s's first step needs to be a init block", forStatement.name)
	}

	forStatement.forInitStatement = initBlock.(*LoopInitStatement)

	forStatement.conditionStep = forStatement.steps[1]

	updateBlock := forStatement.steps[2]

	if updateBlock.Name() != "update" {
		return fmt.Errorf("%s's first step needs to be a init block", forStatement.name)
	}

	forStatement.forUpdateStatement = updateBlock.(*LoopUpdateStatement)

	doBlock := forStatement.steps[3]

	if doBlock.Name() != "do" {
		return fmt.Errorf("%s's first step needs to be a init block", forStatement.name)
	}

	forStatement.forDoStatement = doBlock.(*LoopDoStatement)

	return nil
}

func (forStatement *ForStatement) Execute(scope *Scope) (any, error) {

	if err := forStatement.InitalizeAndCheck(); err != nil {
		return nil, err
	}

	forScope := Scope{}
	forScope.variables = map[string]any{}
	forScope.functions = map[string][]Step{}
	forScope.parent = scope

	var atleastOneIteration bool = false
	if initResult, initErr := forStatement.forInitStatement.Execute(&forScope); initErr == nil {
		for iterationIndex := 0; true; iterationIndex++ {
			forScope.variables["__iterationIndex__"] = iterationIndex
			if conditionResult, conditionErr := forStatement.conditionStep.Execute(&forScope); conditionErr == nil {
				if ToBool(conditionResult) {
					atleastOneIteration = true
					if iterationResult, iterationErr := RunSteps(&forScope, forStatement.forDoStatement, forStatement.forUpdateStatement); iterationErr != nil {
						return iterationResult, iterationErr
					}
				} else {
					return atleastOneIteration, nil
				}
			} else {
				return atleastOneIteration, conditionErr
			}
		}
	} else {
		return initResult, initErr
	}

	return atleastOneIteration, nil
}

func createForStatementStep(name string, parameters map[string]string, body string) (Step, error) {
	forStatement := &ForStatement{}
	forStatement.name = name
	forStatement.parameters = parameters
	forStatement.body = body
	return forStatement, nil
}

func createLoopInitStatementStep(name string, parameters map[string]string, body string) (Step, error) {
	loopInitStatement := &LoopInitStatement{}
	loopInitStatement.name = name
	loopInitStatement.parameters = parameters
	loopInitStatement.body = body
	return loopInitStatement, nil
}

func createLoopUpdateStatementStep(name string, parameters map[string]string, body string) (Step, error) {
	loopUpateStatement := &LoopUpdateStatement{}
	loopUpateStatement.name = name
	loopUpateStatement.parameters = parameters
	loopUpateStatement.body = body
	return loopUpateStatement, nil
}

func createLoopDoStatementStep(name string, parameters map[string]string, body string) (Step, error) {
	loopDoStatement := &LoopDoStatement{}
	loopDoStatement.name = name
	loopDoStatement.parameters = parameters
	loopDoStatement.body = body
	return loopDoStatement, nil
}
