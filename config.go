package main

import (
	"io"
	"os"
	"sync"
)

// global state for output; useful with testing
var defaultOut io.Writer = os.Stdout
var out io.Writer = defaultOut
var mu = &sync.Mutex{}

// with inspriation from http://golang.org/src/pkg/log/log.go?s=7258:7285#L229
func setOutput(w io.Writer) {
	mu.Lock()
	defer mu.Unlock()
	out = w
}

func revertOutput() {
	setOutput(defaultOut)
}

// global state for custom imports file name; useful with testing
const defaultCustomName string = "_gen.go"

var customName string = defaultCustomName

func setCustomName(s string) {
	mu.Lock()
	defer mu.Unlock()
	customName = s
}

func revertCustomName() {
	setCustomName(defaultCustomName)
}
