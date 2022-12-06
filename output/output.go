package output

import (
	"fmt"
)

func IteratePrint(printThisList []string) {
	fmt.Println("new list print:")
	for _, printLine := range printThisList {
		fmt.Println(printLine)
	}
}
