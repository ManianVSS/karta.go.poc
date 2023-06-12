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

func (andStatement *AndStatement) Execute(scope *Scope, basedir string) (any, error) {

	if err := andStatement.InitalizeAndCheck(); err != nil {
		return nil, err
	}

	for _, stepTobeAnded := range andStatement.nestedSteps {
		if result, err := stepTobeAnded.Execute(scope, basedir); err == nil {
			if !ToBool(result) {
				return false, nil
			}
		} else {
			return false, err
		}
	}
	return true, nil
}

func createAndStatementStep(parent Step, tag string, attributes map[string]string, text string) (Step, error) {
	andStatement := &AndStatement{}
	andStatement.tag = tag
	andStatement.attributes = attributes
	andStatement.text = text
	return andStatement, nil
}

func (orStatement *OrStatement) InitalizeAndCheck() error {
	_, error := checkAtleastNStep(orStatement.BaseStep, 2)
	return error
}

func (orStatement *OrStatement) Execute(scope *Scope, basedir string) (any, error) {

	if err := orStatement.InitalizeAndCheck(); err != nil {
		return nil, err
	}

	for _, stepTobeOred := range orStatement.nestedSteps {
		if result, err := stepTobeOred.Execute(scope, basedir); err == nil {
			if ToBool(result) {
				return true, nil
			}
		} else {
			return false, err
		}
	}
	return false, nil
}

func createOrStatementStep(parent Step, tag string, attributes map[string]string, text string) (Step, error) {
	orStatement := &OrStatement{}
	orStatement.tag = tag
	orStatement.attributes = attributes
	orStatement.text = text
	return orStatement, nil
}

func (notStatement *NotStatement) InitalizeAndCheck() error {
	_, error := checkOnlyNStep(notStatement.BaseStep, 1)
	return error
}

func (notStatement *NotStatement) Execute(scope *Scope, basedir string) (any, error) {

	if err := notStatement.InitalizeAndCheck(); err != nil {
		return nil, err
	}

	if result, err := notStatement.nestedSteps[0].Execute(scope, basedir); err == nil {
		return !ToBool(result), nil
	} else {
		return false, err
	}
}

func createNotStatementStep(parent Step, tag string, attributes map[string]string, text string) (Step, error) {
	notStatement := &NotStatement{}
	notStatement.tag = tag
	notStatement.attributes = attributes
	notStatement.text = text
	return notStatement, nil
}
