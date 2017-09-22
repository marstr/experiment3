package print

import "fmt"

type Printer interface {
	Print(message string, iterations int)
}

type StandardOutput struct{}

func (so StandardOutput) Print(message string, iterations int) {
	for i := 0; i < iterations; i++ {
		fmt.Println(message)
	}
}
