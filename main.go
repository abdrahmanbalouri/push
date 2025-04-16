package main

import (
	"fmt"
	"log"
	"os"
	solver "push-swap/solve"
	"strconv"
	"strings"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("achkatkahwar")
		return
	}
	stackA := ReadArgs()
	solver.Solve(stackA)

}

func ReadArgs() (stackA []int) {
	stackAMap := make(map[int]bool)
	for _, num := range strings.Fields(os.Args[1]) {
		n, err := strconv.Atoi(num)
		if err != nil {
			log.Fatal("malkom hagara")
		}
		if exist := stackAMap[n]; exist {
			log.Fatal("malk mrid haka")
		}
		stackAMap[n] = true
		stackA = append(stackA, n)
	}
	return stackA
}
