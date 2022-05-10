package main

import (
	"os/exec"
)

func Process(args []interface{}) interface{} {
	// map args to string[]
	var argsStr []string
	for _, arg := range args {
		argsStr = append(argsStr, arg.(string))
	}

	cmd := exec.Command(argsStr[0], argsStr[1:]...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return ""
	}
	return string(out)
}
