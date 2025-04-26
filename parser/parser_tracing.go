package parser

import (
	"fmt"
	"strings"
)

var traceLevel int = 0

const traceIdentPlaceholder string = "\t"
const enabled = false

func identLevel() string {
	return strings.Repeat(traceIdentPlaceholder, traceLevel-1)
}

func tracePrint(fs string) {
	fmt.Printf("%s%s\n", identLevel(), fs)
}

func incIdent() { traceLevel = traceLevel + 1 }
func decIdent() { traceLevel = traceLevel - 1 }

func trace(msg string) string {
	if enabled {
		incIdent()
		tracePrint("BEGIN " + msg)
		return msg
	}
	return ""
}

func untrace(msg string) {
	if enabled {
		tracePrint("END " + msg)
		decIdent()
	}

}
