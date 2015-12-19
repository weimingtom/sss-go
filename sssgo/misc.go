package sssgo

import (
	"os"
	"fmt"
)

func MiscTrace(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, format, a...)
}

func MiscAppInit() {

}

func MiscAppExit() {
	
}