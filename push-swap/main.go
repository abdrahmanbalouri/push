package main

import (
	"fmt"
	"log"
	"os"
     "g/solve"
	"strconv"
	"strings"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("achkatkahwar")
		return
	}
	input1 := os.Args[1]
	stackA := stacka(input1)
	fmt.Println(stackA)
	switch len(stackA) {
	case 2:
		solver.SortTwo(stackA)
	case 3:
		solver.SortThree(stackA)
	default:
		solver.SortStack(&stackA)
	}

	fmt.Println(stackA)

}

func stacka(s string) (stackA []int) {
	Map := make(map[int]bool)
	for _, num := range strings.Fields(s) {
		n, err := strconv.Atoi(num)
		if err != nil {
			log.Fatal("malkom hagara")
		}
		if exist := Map[n]; exist {
			log.Fatal("malk mrid haka")
		}
		Map[n] = true
		stackA = append(stackA, n)
	}
	return stackA
}
