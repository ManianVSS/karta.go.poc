package xlang

type Scope struct {
	parent    *Scope
	variables map[string]any
	functions map[string][]Step
}

func (scope *Scope) get_variable(name string) (any, bool) {

	if value, ok := scope.variables[name]; ok {
		return value, ok
	} else if scope.parent != nil {
		return scope.parent.get_variable(name)
	} else {
		return nil, false
	}
}

func (scope *Scope) get_function(name string) ([]Step, bool) {

	if value, ok := scope.functions[name]; ok {
		return value, ok
	} else if scope.parent != nil {
		return scope.parent.get_function(name)
	} else {
		return nil, false
	}
}
