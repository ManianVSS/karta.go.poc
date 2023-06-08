package xlang

type Scope struct {
	parent    *Scope
	variables map[string]any
	functions map[string]*FunctionDefinition
}

func (scope *Scope) get_variable(name string) (any, bool) {

	if value, ok := (*scope).variables[name]; ok {
		return value, ok
	} else if value, ok := (*scope).parent.get_variable(name); ok {
		return value, ok
	} else {
		return nil, false
	}
}

func (scope *Scope) get_function(name string) (*FunctionDefinition, bool) {

	if value, ok := (*scope).functions[name]; ok {
		return value, ok
	} else if value, ok := (*scope).parent.get_function(name); ok {
		return value, ok
	} else {
		return nil, false
	}
}
