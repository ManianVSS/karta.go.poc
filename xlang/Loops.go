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

	initBlock := forStatement.nestedSteps[0]

	if initBlock.Name() != "init" {
		return fmt.Errorf("%s's first step needs to be a init block", forStatement.tag)
	}

	forStatement.forInitStatement = initBlock.(*LoopInitStatement)

	forStatement.conditionStep = forStatement.nestedSteps[1]

	updateBlock := forStatement.nestedSteps[2]

	if updateBlock.Name() != "update" {
		return fmt.Errorf("%s's first step needs to be a init block", forStatement.tag)
	}

	forStatement.forUpdateStatement = updateBlock.(*LoopUpdateStatement)

	doBlock := forStatement.nestedSteps[3]

	if doBlock.Name() != "do" {
		return fmt.Errorf("%s's first step needs to be a init block", forStatement.tag)
	}

	forStatement.forDoStatement = doBlock.(*LoopDoStatement)

	return nil
}

func (forStatement *ForStatement) Execute(scope *Scope, basedir string) (any, error) {

	if err := forStatement.InitalizeAndCheck(); err != nil {
		return nil, err
	}

	forScope := Scope{}
	forScope.variables = map[string]any{}
	forScope.functions = map[string][]Step{}
	forScope.parent = scope

	var atleastOneIteration bool = false
	if initResult, initErr := forStatement.forInitStatement.Execute(&forScope, basedir); initErr == nil {
		for iterationIndex := 0; true; iterationIndex++ {
			forScope.variables["__iterationIndex__"] = iterationIndex
			if conditionResult, conditionErr := forStatement.conditionStep.Execute(&forScope, basedir); conditionErr == nil {
				if ToBool(conditionResult) {
					atleastOneIteration = true
					if iterationResult, iterationErr := RunSteps(&forScope, basedir, forStatement.forDoStatement, forStatement.forUpdateStatement); iterationErr != nil {
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

func createForStatementStep(parent Step, tag string, attributes map[string]string, text string) (Step, error) {
	forStatement := &ForStatement{}
	forStatement.tag = tag
	forStatement.attributes = attributes
	forStatement.text = text
	return forStatement, nil
}

func createLoopInitStatementStep(parent Step, tag string, attributes map[string]string, text string) (Step, error) {
	loopInitStatement := &LoopInitStatement{}
	loopInitStatement.tag = tag
	loopInitStatement.attributes = attributes
	loopInitStatement.text = text
	return loopInitStatement, nil
}

func createLoopUpdateStatementStep(parent Step, tag string, attributes map[string]string, text string) (Step, error) {
	loopUpateStatement := &LoopUpdateStatement{}
	loopUpateStatement.tag = tag
	loopUpateStatement.attributes = attributes
	loopUpateStatement.text = text
	return loopUpateStatement, nil
}

func createLoopDoStatementStep(parent Step, tag string, attributes map[string]string, text string) (Step, error) {
	loopDoStatement := &LoopDoStatement{}
	loopDoStatement.tag = tag
	loopDoStatement.attributes = attributes
	loopDoStatement.text = text
	return loopDoStatement, nil
}
