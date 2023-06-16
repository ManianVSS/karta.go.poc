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

	ifStatement.conditionStep = ifStatement.nestedSteps[0]

	thenBlock := ifStatement.nestedSteps[1]

	if thenBlock.Name() != "then" {
		return fmt.Errorf("%s's second step needs to be a then block", ifStatement.tag)
	}

	ifStatement.thenBlock = thenBlock.(*ThenStatement)

	if stepCount > 2 {
		lastIndex := stepCount - 1

		if stepCount > 3 {
			for i := 2; i < lastIndex; i++ {
				elseIfBlock := ifStatement.nestedSteps[i]
				if elseIfBlock.Name() != "elseif" {
					return fmt.Errorf("%s's step numbered %d needs to be a elseif block", ifStatement.tag, i)
				}
				ifStatement.elseIfBlocks = append(ifStatement.elseIfBlocks, ifStatement.nestedSteps[i].(*ElseIfStatement))
			}
			// ifStatement.elseIfBlocks = ifStatement.nestedSteps[2:lastIndex]
		}

		elseBlock := ifStatement.nestedSteps[lastIndex]
		if elseBlock.Name() != "else" {
			return fmt.Errorf("%s's last step needs to be an else block", ifStatement.tag)
		}
		ifStatement.elseBlock = elseBlock.(*ElseStatement)

	}
	return nil
}

func (ifstatement *IfStatement) Execute(scope *Scope, basedir string) (any, error) {

	if err := ifstatement.InitalizeAndCheck(); err != nil {
		return nil, err
	}

	ifScope := Scope{}
	ifScope.variables = map[string]any{}
	ifScope.functions = map[string][]Step{}
	ifScope.parent = scope

	if result, err := ifstatement.conditionStep.Execute(&ifScope, basedir); err == nil {
		if ToBool(result) {
			return ifstatement.thenBlock.Execute(scope, basedir)
		} else {
			if ifstatement.elseIfBlocks != nil {
				for _, elseIfStatement := range ifstatement.elseIfBlocks {
					if result, err := elseIfStatement.Execute(&ifScope, basedir); err == nil {
						if ToBool(result) {
							return true, nil
						}
					} else {
						return false, err
					}
				}
			}

			if ifstatement.elseBlock != nil {
				return ifstatement.elseBlock.Execute(scope, basedir)
			}
		}
	} else {
		return false, err
	}

	return false, nil
}

func createIfStatementStep(parent Step, tag string, attributes map[string]string, text string) (Step, error) {
	ifStatement := &IfStatement{}
	ifStatement.tag = tag
	ifStatement.attributes = attributes
	ifStatement.text = text
	return ifStatement, nil
}

func createThenStatementStep(parent Step, tag string, attributes map[string]string, text string) (Step, error) {
	thenStatement := &ThenStatement{}
	thenStatement.tag = tag
	thenStatement.attributes = attributes
	thenStatement.text = text
	return thenStatement, nil
}

func (elseIfStatement *ElseIfStatement) InitalizeAndCheck() error {

	if _, err := checkOnlyNStep(elseIfStatement.BaseStep, 2); err != nil {
		return err
	}

	elseIfStatement.conditionStep = elseIfStatement.nestedSteps[0]

	thenBlock := elseIfStatement.nestedSteps[1]

	if thenBlock.Name() != "then" {
		return fmt.Errorf("%s's second step needs to be a then block", elseIfStatement.tag)
	}

	elseIfStatement.thenBlock = thenBlock

	return nil
}

func (elseIfstatement *ElseIfStatement) Execute(scope *Scope, basedir string) (any, error) {

	if err := elseIfstatement.InitalizeAndCheck(); err != nil {
		return nil, err
	}

	if result, err := elseIfstatement.conditionStep.Execute(scope, basedir); err == nil {
		if ToBool(result) {
			return elseIfstatement.thenBlock.Execute(scope, basedir)
		}
	} else {
		return false, err
	}

	return false, nil
}

func createElseIfStatementStep(parent Step, tag string, attributes map[string]string, text string) (Step, error) {
	elseIfStatement := &ElseIfStatement{}
	elseIfStatement.tag = tag
	elseIfStatement.attributes = attributes
	elseIfStatement.text = text
	return elseIfStatement, nil
}

func createElseStatementStep(parent Step, tag string, attributes map[string]string, text string) (Step, error) {
	elseStatement := &ElseStatement{}
	elseStatement.tag = tag
	elseStatement.attributes = attributes
	elseStatement.text = text
	return elseStatement, nil
}
