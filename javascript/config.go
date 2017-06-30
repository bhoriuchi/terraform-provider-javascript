package javascript

import "github.com/robertkrimen/otto"

const (
	JS_BOOL   = "bool"
	JS_FLOAT  = "float"
	JS_INT    = "int"
	JS_STRING = "string"
)

var JS_TYPES = []string{JS_BOOL, JS_FLOAT, JS_INT, JS_STRING}

type Config struct {
	Script  string
	Context string
}

type Value struct {
	Name string
	Type string
}

func (c *Config) Vm() (*otto.Otto, error) {
	vm := otto.New()
	return vm, nil
}

func containsString(list []string, value string) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false
}
