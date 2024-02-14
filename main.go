package main

import (
	"Widget_Toolkit/interfaces"
	"fmt"
)

func main() {

	var scheme string
	var scheme_manager *interfaces.SchemeManager
	var err error = fmt.Errorf("")

	for err != nil {
		interfaces.CreateET(err).Display()
		fmt.Print("Choose which color you want the text to be displayed: \033[96mCyan \033[0mor \033[93mYellow")
		fmt.Scanln(&scheme)
		scheme_manager, err = interfaces.MakeSchemeManager(scheme)

	}

	fmt.Print(scheme_manager.Scheme)
}
