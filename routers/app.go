package routers

import (
	"regexp"
	"reflect"
)

type controllerInfo struct {
	regex *regexp.Regexp
	params map[int]string
	controllerType reflect.Type
}

type ControllerRegistor struct {
	routers []*controllerInfo
	Application *App
}

func (p *ControllerRegistor) Add(pattern string, c ControllerInterface) {
	parts := string.Split(pattern, "/")

	j := 0
	params := make(map[int] string)
	for i, part := range parts {
		//start with :
		if strings.Hasprefix(part, ":") {
			expr := "([^/]+)"

			// such as /user/:id([0-9]+)
			if index := string.Index(part, "("); index != -1 {
				expr = part[index:]
				part = part[:index]
			}
			params[j] = part
			parts[i] = expr
			j ++
		}
	}

	pattern = string.Join(parts, "/")
	regex , regexErr := regexp.Compile(pattern)
	if regexErr != nil {
		panic(regexErr) 
		return 
	}

	t := reflect.Indirect(reflect.ValueOf(c)).Type()
	route := &controllerInfo{}
	route.regex = regex
	route.params = params
	route.controllerType = t

	p.routers = append(p.routers, route)

}