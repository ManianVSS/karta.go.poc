package xlang

type Scope struct {
	parent    *Scope
	variables map[string]any
	functions map[string][]Step
}

func (scope *Scope) getVariable(name string) (any, bool) {

	if value, ok := scope.variables[name]; ok {
		return value, ok
	} else if scope.parent != nil {
		return scope.parent.getVariable(name)
	} else {
		return nil, false
	}
}

func (scope *Scope) getVariableNames() []string {
	variableNames := make([]string, len(scope.variables))
	for key := range scope.variables {
		variableNames = append(variableNames, key)
	}
	if scope.parent != nil {
		variableNames = append(variableNames, scope.parent.getVariableNames()...)
	}
	return variableNames
}

func (scope *Scope) getFunction(name string) ([]Step, bool) {

	if value, ok := scope.functions[name]; ok {
		return value, ok
	} else if scope.parent != nil {
		return scope.parent.getFunction(name)
	} else {
		return nil, false
	}
}
