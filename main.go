package main

import (
	"fmt"
	"strconv"

	"github.com/AmogusAzul/ColorTextManagerTest/interfaces"
)

func main() {

	var scheme_manager = interfaces.MakeSchemeManager()

	scheme_manager.Factory.CreateText("Basic sum").Display()

	var a int
	var b int
	var erra error = fmt.Errorf("")
	var errb error = fmt.Errorf("")
	for erra != nil || errb != nil {
		a, erra = strconv.Atoi(scheme_manager.Factory.CreateInput("First Number: ").Read())
		b, errb = strconv.Atoi(scheme_manager.Factory.CreateInput("Second Number: ").Read())
	}

	scheme_manager.Factory.CreateText(fmt.Sprintf("Result: %d", a+b)).Display()
}
