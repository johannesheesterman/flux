package main

import (  
    "fmt"
	"syscall/js"
)

func runFlux() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 1 { 
			return "Invalid no of arguments passed"
		}

		fmt.Println("Run Flux!")
		return nil
	})
}

func main() {  
    js.Global().Set("runFlux", runFlux())
	<-make(chan bool) // Keep program running
}