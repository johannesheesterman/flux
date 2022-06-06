package parser

import (
	"fmt"
	"os/exec"
)

func Process(args []interface{}) interface{} {
	argsStr := StrArr(args)
	cmd := exec.Command(argsStr[0], argsStr[1:]...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return ""
	}
	return string(out)
}

func Print(args []interface{}) interface{} {
	fmt.Println(args...)
	return nil
}

func StrArr(args []interface{}) []string {
	var strArr []string
	for _, arg := range args {
		strArr = append(strArr, arg.(string))
	}
	return strArr
}
