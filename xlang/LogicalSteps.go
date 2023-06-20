package xlang

func init() {
	stepDefMap["and"] = createAndStatementStep
	stepDefMap["or"] = createOrStatementStep
	stepDefMap["not"] = createNotStatementStep
}

type AndStatement struct {
	BaseStep
}

type OrStatement struct {
	BaseStep
}

type NotStatement struct {
	BaseStep
}

func (andStatement *AndStatement) InitalizeAndCheck() error {
	_, error := checkAtleastNStep(andStatement.BaseStep, 2)
	return error
}

func (andStatement *AndStatement) Execute(scope *Scope) (any, error) {

	if err := andStatement.InitalizeAndCheck(); err != nil {
		return nil, err
	}

	for _, stepTobeAnded := range andStatement.steps {
		if result, err := stepTobeAnded.Execute(scope); err == nil {
			if !ToBool(result) {
				return false, nil
			}
		} else {
			return false, err
		}
	}
	return true, nil
}

func createAndStatementStep(name string, parameters map[string]string, body string) (Step, error) {
	andStatement := &AndStatement{}
	andStatement.name = name
	andStatement.parameters = parameters
	andStatement.body = body
	return andStatement, nil
}

func (orStatement *OrStatement) InitalizeAndCheck() error {
	_, error := checkAtleastNStep(orStatement.BaseStep, 2)
	return error
}

func (orStatement *OrStatement) Execute(scope *Scope) (any, error) {

	if err := orStatement.InitalizeAndCheck(); err != nil {
		return nil, err
	}

	for _, stepTobeOred := range orStatement.steps {
		if result, err := stepTobeOred.Execute(scope); err == nil {
			if ToBool(result) {
				return true, nil
			}
		} else {
			return false, err
		}
	}
	return false, nil
}

func createOrStatementStep(name string, parameters map[string]string, body string) (Step, error) {
	orStatement := &OrStatement{}
	orStatement.name = name
	orStatement.parameters = parameters
	orStatement.body = body
	return orStatement, nil
}

func (notStatement *NotStatement) InitalizeAndCheck() error {
	_, error := checkOnlyNStep(notStatement.BaseStep, 1)
	return error
}

func (notStatement *NotStatement) Execute(scope *Scope) (any, error) {

	if err := notStatement.InitalizeAndCheck(); err != nil {
		return nil, err
	}

	if result, err := notStatement.steps[0].Execute(scope); err == nil {
		return !ToBool(result), nil
	} else {
		return false, err
	}
}

func createNotStatementStep(name string, parameters map[string]string, body string) (Step, error) {
	notStatement := &NotStatement{}
	notStatement.name = name
	notStatement.parameters = parameters
	notStatement.body = body
	return notStatement, nil
}
