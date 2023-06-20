package xlang

import "fmt"

func init() {
	stepDefMap["if"] = createIfStatementStep
	stepDefMap["then"] = createThenStatementStep
	stepDefMap["elseif"] = createElseIfStatementStep
	stepDefMap["else"] = createElseStatementStep
}

type IfStatement struct {
	BaseStep
	conditionStep Step
	thenBlock     *ThenStatement
	elseIfBlocks  []*ElseIfStatement
	elseBlock     *ElseStatement
}

type ThenStatement struct {
	BaseStep
}

type ElseIfStatement struct {
	BaseStep
	conditionStep Step
	thenBlock     Step
}

type ElseStatement struct {
	BaseStep
}

func (ifStatement *IfStatement) InitalizeAndCheck() error {

	stepCount, err := checkAtleastNStep(ifStatement.BaseStep, 2)

	if err != nil {
		return err
	}

	ifStatement.conditionStep = ifStatement.steps[0]

	thenBlock := ifStatement.steps[1]

	if thenBlock.Name() != "then" {
		return fmt.Errorf("%s's second step needs to be a then block", ifStatement.name)
	}

	ifStatement.thenBlock = thenBlock.(*ThenStatement)

	if stepCount > 2 {
		lastIndex := stepCount - 1

		if stepCount > 3 {
			for i := 2; i < lastIndex; i++ {
				elseIfBlock := ifStatement.steps[i]
				if elseIfBlock.Name() != "elseif" {
					return fmt.Errorf("%s's step numbered %d needs to be a elseif block", ifStatement.name, i)
				}
				ifStatement.elseIfBlocks = append(ifStatement.elseIfBlocks, ifStatement.steps[i].(*ElseIfStatement))
			}
			// ifStatement.elseIfBlocks = ifStatement.nestedSteps[2:lastIndex]
		}

		elseBlock := ifStatement.steps[lastIndex]
		if elseBlock.Name() != "else" {
			return fmt.Errorf("%s's last step needs to be an else block", ifStatement.name)
		}
		ifStatement.elseBlock = elseBlock.(*ElseStatement)

	}
	return nil
}

func (ifstatement *IfStatement) Execute(scope *Scope) (any, error) {

	if err := ifstatement.InitalizeAndCheck(); err != nil {
		return nil, err
	}

	ifScope := Scope{}
	ifScope.variables = map[string]any{}
	ifScope.functions = map[string][]Step{}
	ifScope.parent = scope

	if result, err := ifstatement.conditionStep.Execute(&ifScope); err == nil {
		if ToBool(result) {
			return ifstatement.thenBlock.Execute(scope)
		} else {
			if ifstatement.elseIfBlocks != nil {
				for _, elseIfStatement := range ifstatement.elseIfBlocks {
					if result, err := elseIfStatement.Execute(&ifScope); err == nil {
						if ToBool(result) {
							return true, nil
						}
					} else {
						return false, err
					}
				}
			}

			if ifstatement.elseBlock != nil {
				return ifstatement.elseBlock.Execute(scope)
			}
		}
	} else {
		return false, err
	}

	return false, nil
}

func createIfStatementStep(name string, parameters map[string]string, body string) (Step, error) {
	ifStatement := &IfStatement{}
	ifStatement.name = name
	ifStatement.parameters = parameters
	ifStatement.body = body
	return ifStatement, nil
}

func createThenStatementStep(name string, parameters map[string]string, body string) (Step, error) {
	thenStatement := &ThenStatement{}
	thenStatement.name = name
	thenStatement.parameters = parameters
	thenStatement.body = body
	return thenStatement, nil
}

func (elseIfStatement *ElseIfStatement) InitalizeAndCheck() error {

	if _, err := checkOnlyNStep(elseIfStatement.BaseStep, 2); err != nil {
		return err
	}

	elseIfStatement.conditionStep = elseIfStatement.steps[0]

	thenBlock := elseIfStatement.steps[1]

	if thenBlock.Name() != "then" {
		return fmt.Errorf("%s's second step needs to be a then block", elseIfStatement.name)
	}

	elseIfStatement.thenBlock = thenBlock

	return nil
}

func (elseIfstatement *ElseIfStatement) Execute(scope *Scope) (any, error) {

	if err := elseIfstatement.InitalizeAndCheck(); err != nil {
		return nil, err
	}

	if result, err := elseIfstatement.conditionStep.Execute(scope); err == nil {
		if ToBool(result) {
			return elseIfstatement.thenBlock.Execute(scope)
		}
	} else {
		return false, err
	}

	return false, nil
}

func createElseIfStatementStep(name string, parameters map[string]string, body string) (Step, error) {
	elseIfStatement := &ElseIfStatement{}
	elseIfStatement.name = name
	elseIfStatement.parameters = parameters
	elseIfStatement.body = body
	return elseIfStatement, nil
}

func createElseStatementStep(name string, parameters map[string]string, body string) (Step, error) {
	elseStatement := &ElseStatement{}
	elseStatement.name = name
	elseStatement.parameters = parameters
	elseStatement.body = body
	return elseStatement, nil
}
