package helper

import (
	"encoding/json"
	"fmt"
)

func Dump(contents ...interface{}) {
	fmt.Print("************")
	for _, content := range contents {
		if str, ok := content.(string); ok {
			fmt.Print(str)
		} else {
			c, _ := json.Marshal(content)
			fmt.Print(string(c))
			fmt.Print("||")
		}
	}
	fmt.Println("***********")
}
