package xlang

func Main() {
	mystep := GenericStep{}
	mystep.init(nil, "sample", map[string]string{"a": "10", "b": "true", "c": "30.0"}, "Some random Text\nLine 2.", nil)

	echo := Echo{}
	echo.init(nil, "echo", map[string]string{"message": "Hello world!!", "b": "true", "c": "30.0"}, "Some random Text\nLine 2.", &mystep)
	echo.execute(Scope{})
}
